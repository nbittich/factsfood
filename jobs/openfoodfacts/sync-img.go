package openfoodfacts

import (
	"errors"
	"fmt"
	"log"
	"mime"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/nbittich/factsfood/config"
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
	resCh          chan imgSyncResult
}

type imgSyncResult struct {
	processed    int64
	synced       int64
	noSyncNeeded int64
}

var imgCollection = db.GetCollection(OpenFoodFactsImgCollection)

func (sij SyncImg) Process(job *jobTypes.Job) (*jobTypes.JobResult, error) {
	jr := jobTypes.JobResult{
		Key:       job.Key,
		Status:    types.ERROR,
		CreatedAt: time.Now(),
		Logs:      make([]jobTypes.Log, 0, InitialCapLogs),
	}
	sr := imgSyncResult{}
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
	resChan := make(chan imgSyncResult, *jp.Parallelism)
	chunkSize := offCount / *jp.Parallelism
	offset := int64(0)
	for offset < offCount {
		wg.Add(1)
		jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("spawning goroutine for offset:%d, chunkSize:%d", offset, chunkSize)))

		go syncImgWorker(syncImgWorkerParam{
			ctx:            ctx,
			batchSize100Ms: *jp.BatchSize100Ms,
			offset:         offset,
			chunkSize:      chunkSize,
			wg:             &wg,
			errCh:          errorChan,
			resCh:          resChan,
		})
		offset += chunkSize
	}

	go func() {
		wg.Wait()
		close(errorChan)
		close(resChan)
	}()
	for errorChan != nil || resChan != nil {
		select {
		case e, ok := <-errorChan:
			if ok {
				cancel()
				return jobs.StatusError(&jr, e)
			} else {
				errorChan = nil
			}
		case r, ok := <-resChan:
			if ok {
				sr.synced += r.synced
				sr.processed += r.processed
				sr.noSyncNeeded += r.noSyncNeeded
			} else {
				resChan = nil
			}
		}
	}

	jr.Metadata = map[string]interface{}{"result": sr}
	jr.Status = types.SUCCESS
	jr.UpdatedAt = time.Now()
	jr.Logs = append(jr.Logs, jobs.NewLog("sync images finished"))

	return &jr, nil
}

func syncImgWorker(wp syncImgWorkerParam) {
	defer wp.wg.Done()
	currBatch := wp.offset
	pageSize := int64(wp.batchSize100Ms)
	maxbatch := wp.offset + wp.chunkSize
	buf := make([]types.Identifiable, 0, wp.batchSize100Ms)
	res := imgSyncResult{}
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
				if (r.OpenFoodFactImg == nil) || r.OpenFoodFactImg.LastImageT != r.LastImageT {
					if r.LastImageT == 0 {
						continue // fixme maybe delete the openfoodfact_img entirely in this case
					}
					offi := r.OpenFoodFactImg
					if offi == nil {
						offi = new(openfoodfacts.OpenFoodFactImg)
					}
					p, e := downloadImg(r.ImageURL, offi.ImageURL)
					if e != nil {
						log.Println("warning while downloading image:", e)
					} else {
						offi.ImageURL = p
					}
					p, e = downloadImg(r.ImageSmallURL, offi.ImageSmallURL)
					if e != nil {
						log.Println("warning while downloading image:", e)
					} else {
						offi.ImageSmallURL = p
					}
					p, e = downloadImg(r.ImageNutritionURL, offi.ImageNutritionURL)
					if e != nil {
						log.Println("warning while downloading image:", e)
					} else {
						offi.ImageNutritionURL = p
					}
					p, e = downloadImg(r.ImageNutritionSmallURL, offi.ImageNutritionSmallURL)
					if e != nil {
						log.Println("warning while downloading image:", e)
					} else {
						offi.ImageNutritionSmallURL = p
					}
					p, e = downloadImg(r.ImageIngredientsURL, offi.ImageIngredientsURL)
					if e != nil {
						log.Println("warning while downloading image:", e)
					} else {
						offi.ImageIngredientsURL = p
					}
					p, e = downloadImg(r.ImageIngredientsSmallURL, offi.ImageIngredientsSmallURL)
					if e != nil {
						log.Println("warning while downloading image:", e)
					} else {
						offi.ImageIngredientsSmallURL = p
					}
					offi.LastImageT = r.LastImageT
					offi.OpenFoodFactID = r.Code
					buf = append(buf, offi)
					res.synced += 1
				} else {
					res.noSyncNeeded += 1
				}
				res.processed += 1

			}
			if len(buf) != 0 {
				if err = db.InsertOrUpdateMany(wp.ctx, buf, imgCollection); err != nil {
					wp.errCh <- err
				}
			}
			buf = buf[:0]
			currBatch += pageSize
			time.Sleep(time.Millisecond * sleepBetweenBatchesMs)
		}
	}
	wp.resCh <- res
}

func downloadImg(uri string, oldImageURI string) (string, error) {
	if uri != "" {
		u, err := url.Parse(uri)
		if err != nil {
			return "", err
		}
		p := strings.Split(u.Path, "/")
		filename := ""
		for i := len(p) - 1; i != 0; i-- {
			ext := filepath.Ext(p[i])
			ct := mime.TypeByExtension(ext)
			if strings.HasPrefix(ct, "image/") {
				filename = uuid.New().String() + ext
				break
			}
		}
		if filename == "" {
			return "", fmt.Errorf("could not guess filename of %s", u)
		}
		fp := filepath.Join(config.StaticDirectory, filename)
		_, err = jobs.DownloadFile(uri, fp, false)
		if err != nil {
			return "", err
		}
		oldImagePath := strings.ReplaceAll(oldImageURI, "/static", config.StaticDirectory)
		if oldImagePath != "" {
			os.Remove(oldImagePath)
		}
		return path.Join("/static", filename), err
	}
	return "", errors.New("uri is empty")
}
