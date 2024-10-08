package openfoodfacts

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"slices"
	"sync"
	"sync/atomic"
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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

const (
	InitialSyncJobKey           = "OFF_INITIAL_SYNC_JOB"
	SyncJobKey                  = "OFF_SYNC_JOB"
	lineBufferSize              = 131_072 // 128kb buffer size
	maxFailedCsvLineInJobResult = 256     // Max failed csv lines to persist in job result

)

type Sync struct{}

type failedCSVLine struct {
	Line  string `json:"line"`
	Error string `json:"error"`
}

type syncWorkerParam struct {
	ctx            context.Context
	batchSize100Ms uint
	separator      rune
	data           []byte
	wg             *sync.WaitGroup
	errCh          chan error
	warningCh      chan failedCSVLine
	counter        *atomic.Uint64
}

type syncJobParam struct {
	Endpoint       string `mapstructure:"endpoint" validate:"required,url"`
	SeparatorStr   string `mapstructure:"separator" validate:"required,len=1"`
	separator      rune
	Gzipped        bool   `mapstructure:"gzip"`
	Parallelism    *int64 `mapstructure:"parallelism"`
	BatchSize100Ms *uint  `mapstructure:"batchSize100Ms"`
}

// creates openfoodfacts collection and seq index
func init() {
	ctx, cancel := context.WithTimeout(context.Background(), config.MongoCtxTimeout)
	defer cancel()
	collections, err := db.DB.ListCollectionNames(ctx, &bson.D{}, options.ListCollections().SetNameOnly(true))
	if err != nil {
		panic(fmt.Sprint("could not list collections", err))
	}
	if slices.Contains(collections, OpenFoodFactsCollection) {
		log.Println(OpenFoodFactsCollection, "collection already exists")
		return
	}
	if err = db.DB.CreateCollection(ctx, OpenFoodFactsCollection, options.CreateCollection()); err != nil {
		panic(fmt.Sprint("could not create collection", err))
	}
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "seq", Value: 1}, // 1 for ascending, -1 for descending
		},
		Options: options.Index().SetUnique(true),
	}
	if _, err = offCollection.Indexes().CreateOne(ctx, indexModel); err != nil {
		panic(fmt.Sprint("could not create index", err))
	}
}

func (is Sync) Process(job *jobTypes.Job) (*jobTypes.JobResult, error) {
	failedCsvLine := make([]*failedCSVLine, 0, maxFailedCsvLineInJobResult)
	jr := jobTypes.JobResult{
		Key:       job.Key,
		Status:    types.ERROR,
		CreatedAt: time.Now(),
		Logs:      make([]jobTypes.Log, 0, InitialCapLogs),
	}
	jp, err := jobs.ValidateJobAndGetParam(job, &jr,
		func(jp *syncJobParam) (*syncJobParam, error) {
			jp.separator = rune(jp.SeparatorStr[0])

			if jp.Parallelism == nil {
				jr.Logs = append(jr.Logs, jobs.NewLog("warning! missing parallelism param. fallback to thread counts"))
				*jp.Parallelism = int64(runtime.NumCPU())
			}
			if jp.BatchSize100Ms == nil {
				jr.Logs = append(jr.Logs, jobs.NewLog("warning! missing batchSize100Ms param. fallback to 100"))
				*jp.BatchSize100Ms = 100
			}
			return jp, nil
		}, InitialSyncJobKey, SyncJobKey)
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

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	errorChan := make(chan error, *jp.Parallelism)
	warningChan := make(chan failedCSVLine, *jp.Parallelism)

	offset := headerLine(data) // skip header line
	maxChunkSize := (fileLen - offset) / *jp.Parallelism
	maxChunkSize -= maxChunkSize % int64(os.Getpagesize())

	jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("Extracting CSV using '%c' separator", jp.separator)))

	log.Println("spawning goroutine(s)")

	var counter atomic.Uint64
	for offset < fileLen {
		end := offset + maxChunkSize
		if end > fileLen {
			end = fileLen
		}

		currentLastNewLine := offset + int64(lastNewline(data[offset:end]))
		partition := data[offset:currentLastNewLine]
		wg.Add(1)

		go syncWorker(syncWorkerParam{
			ctx:            ctx,
			separator:      jp.separator,
			batchSize100Ms: *jp.BatchSize100Ms,
			data:           partition,
			wg:             &wg,
			errCh:          errorChan,
			warningCh:      warningChan,
			counter:        &counter,
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
				cancel()
				return jobs.StatusError(&jr, e)
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
		jr.Metadata = map[string]interface{}{"failedCsvLine": failedCsvLine}
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

func syncWorker(wp syncWorkerParam) {
	defer wp.wg.Done()
	batchSize := int(wp.batchSize100Ms)
	buf := make([]byte, 0, lineBufferSize)
	batch := make([]types.Identifiable, 0, batchSize)
	out := make([]*offType.OpenFoodFact, 0, 1) // only to please gocsv

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
					wp.warningCh <- failedCSVLine{
						Line:  string(buf),
						Error: err.Error(),
					}
				} else {
					if len(out) != 1 {
						wp.errCh <- fmt.Errorf("len of entry != 1: %d. %v", len(out), out)
						return
					}

					o := out[0]
					o.Seq = wp.counter.Add(1)
					batch = append(batch, o)

				}
				// clear buffers
				buf = buf[:0]
				out = out[:0]
				if len(batch) == batchSize {
					// flush
					if err := db.InsertOrUpdateMany(wp.ctx, batch, offCollection); err != nil {
						wp.errCh <- err
						return
					}
					batch = batch[:0]
					time.Sleep(time.Millisecond * sleepBetweenBatchesMs)
				}
			} else {
				buf = append(buf, v)
			}

		}
	}
	if len(batch) > 0 {
		// final flush
		if err := db.InsertOrUpdateMany(wp.ctx, batch, offCollection); err != nil {
			wp.errCh <- err
			return
		}
	}
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
