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
	"github.com/gocarina/gocsv"
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

	jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("Extracting CSV using '%c' separator", separator)))

	for offset < fileLen {
		end := offset + maxChunkSize
		if end > fileLen {
			end = fileLen
		}

		currentLastNewLine := offset + int64(lastNewline(data[offset:end]))
		partition := data[offset:currentLastNewLine]
		wg.Add(1)

		go worker(workerParam{
			ctx:       ctx,
			separator: separator,
			data:      partition,
			wg:        &wg,
			errCh:     ch,
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
	ctx       context.Context
	separator rune
	data      []byte
	wg        *sync.WaitGroup
	errCh     chan error
}

func worker(wp workerParam) {
	defer wp.wg.Done()
	batchSize := 100
	buf := make([]byte, 0, 131_072) // 128kb buffer
	batch := make([]types.Identifiable, 0, batchSize)
	col := db.GetCollection("openfoodfacts")
	out := make([]*offType.OpenFoodFactCSVEntry, 0, 1) // only to please gocsv

	for _, v := range wp.data {
		select {
		case <-wp.ctx.Done():
			log.Println("CSV Goroutine cancelled")
			return
		default:
			if len(batch) == batchSize {
				// flush
				if err := db.InsertMany(wp.ctx, batch, col); err != nil {
					wp.errCh <- err
					return
				}
				batch = batch[:0]
			}
			if v == '\n' {
				// process buf
				csvReader := csv.NewReader(bytes.NewReader(buf))
				csvReader.Comma = wp.separator
				err := gocsv.UnmarshalCSVWithoutHeaders(csvReader, &out)
				if err != nil {
					log.Println("warning! error in row:", string(buf))
				} else {
					if len(out) != 1 {
						wp.errCh <- fmt.Errorf("len of entry != 1: %d. %v", len(out), out)
						break
					}
					batch = append(batch, out[0])

				}
				// clear buffers
				buf = buf[:0]
				out = out[:0]
			} else {
				buf = append(buf, v)
			}

		}
	}
	if len(batch) > 0 {
		// final flush
		if err := db.InsertMany(wp.ctx, batch, col); err != nil {
			wp.errCh <- err
			return
		}
	}
}
