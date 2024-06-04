package manager

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/adhocore/gronx"
	"github.com/nbittich/factsfood/config"
	"github.com/nbittich/factsfood/jobs/openfoodfacts"
	"github.com/nbittich/factsfood/services/db"
	"github.com/nbittich/factsfood/types/job"
	"go.mongodb.org/mongo-driver/bson"
)

var jobProcessors map[string]job.JobProcessor

var (
	jobCollection       = db.GetCollection("job")
	jobResultCollection = db.GetCollection("jobResult")
	started             = false
)

func init() {
	jobProcessors = make(map[string]job.JobProcessor, 3)
	Register(&openfoodfacts.Sync{}, openfoodfacts.InitialSyncJobKey, openfoodfacts.SyncJobKey)
	Register(&openfoodfacts.SyncImg{}, openfoodfacts.SyncImgJobKey)
	initJobs()
}

func initJobs() {
	jobs, err := getAllNonDisabledJobs()
	if err != nil {
		log.Fatal("could not load jobs:", err)
	}
	for _, j := range jobs {
		j.Running = false
		err := setNextSchedule(&j)
		if err != nil {
			log.Println("could not set next schedule for job", j.Name, "err", err)
			j.Disabled = true
		}
		if _, err := db.Save(&j, jobCollection); err != nil {
			log.Println("could not persist job : ", err)
		}
	}
}

func Register(processor job.JobProcessor, keys ...string) {
	for _, key := range keys {
		log.Println("register job processor", key)
		jobProcessors[key] = processor
	}
}

func getEnabledAndNonRunningJobs() ([]job.Job, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.MongoCtxTimeout)
	defer cancel()
	return db.Find[job.Job](ctx, &bson.M{"disabled": false, "running": false}, jobCollection, nil)
}

func getAllNonDisabledJobs() ([]job.Job, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.MongoCtxTimeout)
	defer cancel()
	return db.Find[job.Job](ctx, &bson.M{"disabled": false}, jobCollection, nil)
}

func Start() {
	if started {
		log.Fatal("You cannot start the manager twice")
	}
	started = true
	var wg sync.WaitGroup

	for {
		jobs, err := getEnabledAndNonRunningJobs()
		if err != nil {
			log.Println("could not load jobs:", err)
		}

		for _, j := range jobs {
			processor, ok := jobProcessors[j.Key]
			if !ok {
				log.Println("processor not found: ", j.Key)
				continue
			}
			wg.Add(1)
			go process(&wg, &j, processor)
		}
		wg.Wait() // make sure that all jobs are set to running
		time.Sleep(time.Second * 5)
	}
}

func setNextSchedule(j *job.Job) error {
	if !j.SpecificDate.IsZero() {
		j.NextSchedule = j.SpecificDate
	} else {
		nextTick, err := gronx.NextTick(j.CronExpression, true)
		if err != nil {
			return err
		}
		j.NextSchedule = nextTick
	}
	return nil
}

func process(wg *sync.WaitGroup, j *job.Job, processor job.JobProcessor) {
	if j.NextSchedule.Before(time.Now()) {
		log.Println("starting job", j.Key)
		j.Running = true
		j.UpdatedAt = time.Now()
		_, err := db.Save(j, jobCollection)
		wg.Done()
		if err != nil {
			log.Println("could not set job to running: ", err)
			return
		}
		res, err := processor.Process(j)
		j.UpdatedAt = time.Now()
		j.Running = false
		switch {
		case err == job.INVALIDPARAM, err == job.DISABLED, !j.SpecificDate.IsZero():
			j.Disabled = true
		case err != nil:
			log.Println("job error: ", err)
			fallthrough
		default:
			if err = setNextSchedule(j); err != nil {
				log.Println("could not set next schedule", err)
				j.Disabled = true
			}
		}
		if _, err := db.Save(res, jobResultCollection); err != nil {
			log.Println("could not persist job result: ", err)
		}
		if _, err := db.Save(j, jobCollection); err != nil {
			log.Println("could not persist job : ", err)
		}
	} else {
		wg.Done()
	}
}
