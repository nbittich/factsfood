package openfoodfacts

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"sync"
	"time"

	"github.com/edsrzf/mmap-go"
	"github.com/google/uuid"
	"github.com/nbittich/factsfood/config"
	"github.com/nbittich/factsfood/jobs"
	"github.com/nbittich/factsfood/services/db"
	"github.com/nbittich/factsfood/types"
	jobTypes "github.com/nbittich/factsfood/types/job"
	offType "github.com/nbittich/factsfood/types/openfoodfacts"
	"golang.org/x/net/context"
)

const (
	InitialSyncJobKey = "OFF_INITIAL_SYNC_JOB"
	endpointParamKey  = "endpoint"
	csvSeparatorKey   = "separator"
	gzipKey           = "gzip"
	parallelismKey    = "parallelism"
)

type InitialSync struct{}

func (is InitialSync) Process(job *jobTypes.Job) (*jobTypes.JobResult, error) {
	jr := jobTypes.JobResult{
		Key:       job.Key,
		Status:    types.ERROR,
		CreatedAt: time.Now(),
		Logs:      make([]jobTypes.Log, 0, 10),
	}
	if job.Disabled {
		return jobs.StatusError(&jr, jobTypes.DISABLED)
	}

	if job.Key != InitialSyncJobKey {
		return jobs.StatusError(&jr, jobTypes.BADKEY)
	}

	endpoint, ok := job.Params[endpointParamKey].(string)

	if !ok {
		jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("missing or invalid endpoint param %s: %v", endpointParamKey, job.Params)))
		return jobs.StatusError(&jr, jobTypes.INVALIDPARAM)
	}

	separatorStr, ok := job.Params[csvSeparatorKey].(string)
	if !ok || len(separatorStr) != 1 {
		jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("missing or invalid separator param %s: %v", csvSeparatorKey, job.Params)))
		return jobs.StatusError(&jr, jobTypes.INVALIDPARAM)
	}
	separator := []rune(separatorStr)[0]

	gzipped, ok := job.Params[gzipKey].(bool)
	if !ok {
		jr.Logs = append(jr.Logs, jobs.NewLog("warning! missing or invalid gzip param. fallback to false"))
		gzipped = false
	}

	var parallelism int
	parallelismI32, ok := job.Params[parallelismKey].(int32)
	if !ok {
		jr.Logs = append(jr.Logs, jobs.NewLog("warning! missing or invalid parallelism param. fallback to thread counts"))
		parallelism = runtime.NumCPU()
	} else {
		parallelism = int(parallelismI32)
	}

	tempPath := path.Join(config.TempDir, fmt.Sprintf("%s.csv", uuid.New().String()))
	jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("downloading CSV from %s and saved to %s", endpoint, tempPath)))

	fileLen, err := jobs.DownloadFile(endpoint, tempPath, gzipped)
	if err != nil {
		return jobs.StatusError(&jr, err)
	}

	jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("CSV file %s downloaded, len: %d", tempPath, fileLen)))

	// CSV shenaningans
	// TODO extract to a separate func as it will be used also in delta-sync.go
	jr.Logs = append(jr.Logs, jobs.NewLog("mmap csv file"))
	f, err := os.Open(tempPath)
	if err != nil {
		return jobs.StatusError(&jr, err)
	}
	defer f.Close()
	data, err := mmap.Map(f, mmap.RDONLY, 0)
	if err != nil {
		return jobs.StatusError(&jr, err)
	}

	offset := headerLine(data) // skip header line
	maxChunkSize := (fileLen - offset) / int64(parallelism)
	maxChunkSize -= maxChunkSize - 4096 // 4k page
	var wg sync.WaitGroup
	ch := make(chan error, 1)
	ctx, cancel := context.WithCancel(context.Background())

	jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("Extracting CSV using '%b' separator", separator)))

	for offset < fileLen {
		end := offset + maxChunkSize
		if end > fileLen {
			end = fileLen
		}

		currentLastNewLine := offset + int64(lastNewline(data[offset:end]))
		wg.Add(1)

		go csvWorker(workerParam{
			ctx:                ctx,
			offset:             offset,
			separator:          separator,
			currentLastNewLine: currentLastNewLine,
			data:               data,
			wg:                 &wg,
			errCh:              ch,
		})
		offset = currentLastNewLine
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for e := range ch {
		if e != nil {
			cancel()
			return jobs.StatusError(&jr, e)
		}
	}
	jr.Status = types.SUCCESS
	jr.UpdatedAt = time.Now()

	return &jr, nil
}

func lastNewline(s []byte) int {
	i := len(s) - 1
	for i > 0 {
		if s[i] == '\n' {
			return i + 1
		}
		i -= 1
	}
	return len(s)
}

func headerLine(s []byte) int64 {
	i := 0
	for s[i] != '\n' {
		i++
	}
	return int64(i + 1)
}

type workerParam struct {
	ctx                context.Context
	offset             int64
	separator          rune
	currentLastNewLine int64
	data               []byte
	wg                 *sync.WaitGroup
	errCh              chan error
}

func mongoSinkWorker(ctx context.Context, wg *sync.WaitGroup, off offType.OpenFoodFactCSVEntry, errChan chan error) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		log.Println("Sink Channel closed")
		return
	default:
		col := db.GetCollection("openfoodfacts")
		if _, err := db.InsertOrUpdate(ctx, &off, col); err != nil {
			errChan <- err
			return
		}
	}
}

func csvWorker(wp workerParam) {
	defer wp.wg.Done()
	buf := make([]byte, 0, 131_072) // 128kb buffer
	select {
	case <-wp.ctx.Done():
		log.Println("CSV Goroutine cancelled")
		return
	default:
		for _, v := range wp.data[wp.offset:wp.currentLastNewLine] {
			if v == '\n' {
				// process buf
				csvReader := csv.NewReader(bytes.NewReader(buf))
				csvReader.Comma = wp.separator
				entry := offType.OpenFoodFactCSVEntry{}
				err := jobs.UnmarshalCSV(csvReader, &entry)
				if err != nil {
					wp.errCh <- err
					break
				}
				wp.wg.Add(1)
				go mongoSinkWorker(wp.ctx, wp.wg, entry, wp.errCh)
				// clear buf
				buf = buf[:0]
			} else {
				buf = append(buf, v)
			}
		}
	}
}
