package openfoodfacts

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/nbittich/factsfood/jobs"
	"github.com/nbittich/factsfood/services/db"
	"github.com/nbittich/factsfood/types"
	jobTypes "github.com/nbittich/factsfood/types/job"
	"github.com/nbittich/factsfood/types/openfoodfacts"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

const (
	SyncImgJobKey = "OFF_SYNC_IMG_JOB"
)

type SyncImg struct{}

type syncImgJobParam struct {
	Parallelism    *int64 `mapstructure:"parallelism"`
	BatchSize100Ms *uint  `mapstructure:"batchSize100Ms"`
}

type syncImgWorkerParam struct {
	ctx            context.Context
	batchSize100Ms uint
	offset         int64
	chunkSize      int64
	page           db.PageOptions
	wg             *sync.WaitGroup
	errCh          chan error
}

var imgCollection = db.GetCollection(OpenFoodFactsImgCollection)

func (sij SyncImg) Process(job *jobTypes.Job) (*jobTypes.JobResult, error) {
	jr := jobTypes.JobResult{
		Key:       job.Key,
		Status:    types.ERROR,
		CreatedAt: time.Now(),
		Logs:      make([]jobTypes.Log, 0, InitialCapLogs),
	}
	jp, err := jobs.ValidateJobAndGetParam(job, &jr,
		func(jp *syncImgJobParam) (*syncImgJobParam, error) {
			if jp.Parallelism == nil {
				jr.Logs = append(jr.Logs, jobs.NewLog("warning! missing parallelism param. fallback to thread counts"))
				*jp.Parallelism = int64(runtime.NumCPU())
			}
			if jp.BatchSize100Ms == nil {
				jr.Logs = append(jr.Logs, jobs.NewLog("warning! missing batchSize100Ms param. fallback to 100"))
				*jp.BatchSize100Ms = 100
			}
			return jp, nil
		}, SyncImgJobKey)
	if err != nil {
		return jobs.StatusError(&jr, err)
	}
	jr.Logs = append(jr.Logs, jobs.NewLog("Starting syncing images"))

	ctx, cancel := context.WithCancel(context.Background())
	offCount, err := db.CountAll(ctx, offCollection)
	jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("OFF count:%d", offCount)))

	if err != nil {
		cancel()
		return jobs.StatusError(&jr, err)
	}

	var wg sync.WaitGroup
	errorChan := make(chan error, *jp.Parallelism)
	chunkSize := offCount / *jp.Parallelism
	offset := int64(0)
	for offset <= offCount {
		wg.Add(1)
		jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("spawning goroutine for offset:%d, chunkSize:%d", offset, chunkSize)))

		go syncImgWorker(syncImgWorkerParam{
			ctx:            ctx,
			batchSize100Ms: *jp.BatchSize100Ms,
			offset:         offset,
			chunkSize:      chunkSize,
			wg:             &wg,
			errCh:          errorChan,
		})
		offset += chunkSize
	}

	go func() {
		wg.Wait()
		close(errorChan)
	}()

	for err := range errorChan {
		cancel()
		return jobs.StatusError(&jr, err)
	}

	jr.UpdatedAt = time.Now()
	jr.Logs = append(jr.Logs, jobs.NewLog("sync images finished"))

	return &jr, nil
}

func syncImgWorker(wp syncImgWorkerParam) {
	defer wp.wg.Done()
	currBatch := wp.offset
	pageSize := int64(wp.batchSize100Ms)
	maxbatch := wp.offset + wp.chunkSize
	for currBatch < maxbatch {
		select {
		case <-wp.ctx.Done():
			log.Println("Goroutine cancelled")
			return
		default:
			pipeline := mongo.Pipeline{
				{
					{Key: "$lookup", Value: bson.D{
						{Key: "from", Value: OpenFoodFactsImgCollection},
						{Key: "localField", Value: "_id"},
						{Key: "foreignField", Value: "openfoodfacts_id"},
						{Key: "as", Value: "openfoodfact_img"},
					}},
				},
				{
					{Key: "$unwind", Value: bson.D{
						{Key: "path", Value: "$openfoodfact_img"},
						{Key: "preserveNullAndEmptyArrays", Value: true},
					}},
				},
				{{Key: "$skip", Value: currBatch}},
				{{Key: "$limit", Value: pageSize}},
			}

			cursor, err := offCollection.Aggregate(wp.ctx, pipeline)
			if err != nil {
				wp.errCh <- err
				return
			}

			results, err := db.CursorToSlice[openfoodfacts.FactsFood](wp.ctx, cursor, int(pageSize))
			if err != nil {
				wp.errCh <- err
				return
			}

			for _, r := range results {
				fmt.Printf("todo: %v\n", r)
			}
			currBatch += pageSize
			time.Sleep(time.Millisecond * sleepBetweenBatchesMs)
		}
	}
}
