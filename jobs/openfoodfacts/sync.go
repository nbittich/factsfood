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
	"github.com/mitchellh/mapstructure"
	"github.com/nbittich/factsfood/config"
	"github.com/nbittich/factsfood/jobs"
	"github.com/nbittich/factsfood/services/db"
	"github.com/nbittich/factsfood/services/utils"
	"github.com/nbittich/factsfood/types"
	jobTypes "github.com/nbittich/factsfood/types/job"
	offType "github.com/nbittich/factsfood/types/openfoodfacts"
	"golang.org/x/net/context"
)

const (
	InitialSyncJobKey           = "OFF_INITIAL_SYNC_JOB"
	SyncJobKey                  = "OFF_SYNC_JOB"
	OpenFoodFactsCollection     = "openfoodfacts"
	endpointParamKey            = "endpoint"
	csvSeparatorKey             = "separator"
	gzipKey                     = "gzip"
	parallelismKey              = "parallelism"
	LineBufferSize              = 131_072 // 128kb buffer size
	MaxFailedCsvLineInJobResult = 256     // Max failed csv lines to persist in job result
	InitialCapLogs              = 10      // Logs slice initial capacity
	SleepBetweenBatchesMs       = 100     // Sleep 100ms to allow mongodb between each batch to rest a lil bit.
	//                                    // If you change this, make sure you  also change the job config in db.
	//                                    // See BatchSize100Ms below
)

type Sync struct{}

type FailedCSVLine struct {
	Line  string `json:"line"`
	Error string `json:"error"`
}

type workerParam struct {
	ctx            context.Context
	batchSize100Ms uint
	separator      rune
	data           []byte
	wg             *sync.WaitGroup
	errCh          chan error
	warningCh      chan FailedCSVLine
}

type jobParam struct {
	Endpoint       string `mapstructure:"endpoint" validate:"required,url"`
	SeparatorStr   string `mapstructure:"separator" validate:"required,len=1"`
	separator      rune
	Gzipped        bool   `mapstructure:"gzip"`
	Parallelism    *int64 `mapstructure:"parallelism"`
	BatchSize100Ms *uint  `mapstructure:"batchSize100Ms"`
}

func (is Sync) Process(job *jobTypes.Job) (*jobTypes.JobResult, error) {
	failedCsvLine := make([]*FailedCSVLine, 0, MaxFailedCsvLineInJobResult)
	jr := jobTypes.JobResult{
		Key:       job.Key,
		Status:    types.ERROR,
		CreatedAt: time.Now(),
		Logs:      make([]jobTypes.Log, 0, InitialCapLogs),
	}
	jp, err := validateJobAndGetParam(job, &jr)
	if err != nil {
		return jobs.StatusError(&jr, err)
	}

	tempPath := path.Join(config.TempDir, fmt.Sprintf("%s.csv", uuid.New().String()))
	jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("downloading CSV from %s and saved to %s", jp.Endpoint, tempPath)))

	fileLen, err := jobs.DownloadFile(jp.Endpoint, tempPath, jp.Gzipped)
	if err != nil {
		return jobs.StatusError(&jr, err)
	}

	jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("CSV file %s downloaded, len: %d", tempPath, fileLen)))

	// CSV shenaningans
	jr.Logs = append(jr.Logs, jobs.NewLog("mmap csv file"))
	f, err := os.Open(tempPath)
	if err != nil {
		return jobs.StatusError(&jr, err)
	}
	defer f.Close()

	data, err := mmap.Map(f, mmap.RDONLY, 0)
	defer data.Unmap()
	if err != nil {
		return jobs.StatusError(&jr, err)
	}

	offset := headerLine(data) // skip header line
	maxChunkSize := (fileLen - offset) / *jp.Parallelism
	maxChunkSize -= maxChunkSize % int64(os.Getpagesize())
	var wg sync.WaitGroup
	errorChan := make(chan error, *jp.Parallelism)
	warningChan := make(chan FailedCSVLine, *jp.Parallelism)
	ctx, cancel := context.WithCancel(context.Background())

	jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("Extracting CSV using '%c' separator", jp.separator)))

	log.Println("spawning goroutine(s)")

	for offset < fileLen {
		end := offset + maxChunkSize
		if end > fileLen {
			end = fileLen
		}

		currentLastNewLine := offset + int64(lastNewline(data[offset:end]))
		partition := data[offset:currentLastNewLine]
		wg.Add(1)

		go worker(workerParam{
			ctx:            ctx,
			separator:      jp.separator,
			batchSize100Ms: *jp.BatchSize100Ms,
			data:           partition,
			wg:             &wg,
			errCh:          errorChan,
			warningCh:      warningChan,
		})
		offset = currentLastNewLine
	}

	go func() {
		wg.Wait()
		close(errorChan)
		close(warningChan)
	}()

	stopWritingWarnings := false
	for errorChan != nil || warningChan != nil {
		select {
		case e, ok := <-errorChan:
			if ok {
				if e != nil {
					cancel()
					return jobs.StatusError(&jr, e)
				}
			} else {
				errorChan = nil
			}
		case w, ok := <-warningChan:
			if ok {
				if cap(failedCsvLine) == len(failedCsvLine) && !stopWritingWarnings {
					log.Println("capacity of warnings reached, fallback to logging instead")
					stopWritingWarnings = true
				}
				if stopWritingWarnings {
					log.Printf("error in unmarshalling csv: %s \t\nline: %s\n", w.Error, w.Line)
				} else {
					failedCsvLine = append(failedCsvLine, &w)
				}
			} else {
				warningChan = nil
			}
		}
	}

	jr.Status = types.SUCCESS
	if len(failedCsvLine) > 0 {
		metadata := make(map[string]interface{}, 1)
		metadata["failedCsvLines"] = failedCsvLine
		jr.Metadata = metadata
	}

	// delete csv
	go func(p string) {
		time.Sleep(time.Minute * 5) // wait for unmap, f.close
		if err := os.Remove(p); err != nil {
			log.Println("could not delete", p, err)
		}
	}(tempPath)

	jr.UpdatedAt = time.Now()
	jr.Logs = append(jr.Logs, jobs.NewLog("sync finished"))

	return &jr, nil
}

func worker(wp workerParam) {
	defer wp.wg.Done()
	batchSize := int(wp.batchSize100Ms)
	buf := make([]byte, 0, LineBufferSize)
	batch := make([]types.Identifiable, 0, batchSize)
	col := db.GetCollection(OpenFoodFactsCollection)
	out := make([]*offType.OpenFoodFactCSVEntry, 0, 1) // only to please gocsv

	for _, v := range wp.data {
		select {
		case <-wp.ctx.Done():
			log.Println("CSV Goroutine cancelled")
			return
		default:
			if v == '\n' {
				// process buf
				csvReader := csv.NewReader(bytes.NewReader(buf))
				csvReader.Comma = wp.separator
				csvReader.LazyQuotes = true
				err := gocsv.UnmarshalCSVWithoutHeaders(csvReader, &out)
				if err != nil {
					wp.warningCh <- FailedCSVLine{
						Line:  string(buf),
						Error: err.Error(),
					}
				} else {
					if len(out) != 1 {
						wp.errCh <- fmt.Errorf("len of entry != 1: %d. %v", len(out), out)
						return
					}
					batch = append(batch, out[0])

				}
				// clear buffers
				buf = buf[:0]
				out = out[:0]
				if len(batch) == batchSize {
					// flush
					if err := db.InsertOrUpdateMany(wp.ctx, batch, col); err != nil {
						wp.errCh <- err
						return
					}
					batch = batch[:0]
					time.Sleep(time.Millisecond * SleepBetweenBatchesMs)
				}
			} else {
				buf = append(buf, v)
			}

		}
	}
	if len(batch) > 0 {
		// final flush
		if err := db.InsertOrUpdateMany(wp.ctx, batch, col); err != nil {
			wp.errCh <- err
			return
		}
	}
}

func validateJobAndGetParam(job *jobTypes.Job, jr *jobTypes.JobResult) (*jobParam, error) {
	if job.Disabled {
		return nil, jobTypes.DISABLED
	}

	if job.Key != InitialSyncJobKey && job.Key != SyncJobKey {
		return nil, jobTypes.BADKEY
	}
	var jp jobParam
	if err := mapstructure.Decode(job.Params, &jp); err != nil {
		jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("error: %s", err)))
		return nil, jobTypes.INVALIDPARAM
	}
	if err := utils.ValidateStruct(&jp); err != nil {
		jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("error: %s", err)))
		return nil, jobTypes.INVALIDPARAM
	}

	jp.separator = rune(jp.SeparatorStr[0])

	if jp.Parallelism == nil {
		jr.Logs = append(jr.Logs, jobs.NewLog("warning! missing parallelism param. fallback to thread counts"))
		*jp.Parallelism = int64(runtime.NumCPU())
	}
	if jp.BatchSize100Ms == nil {
		jr.Logs = append(jr.Logs, jobs.NewLog("warning! missing batchSize100Ms param. fallback to 100"))
		*jp.BatchSize100Ms = 100
	}
	return &jp, nil
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
