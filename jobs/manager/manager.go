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
)

func init() {
	jobProcessors = make(map[string]job.JobProcessor, 1)
	Register(openfoodfacts.InitialSyncJobKey, &openfoodfacts.InitialSync{})
}

func Register(key string, processor job.JobProcessor) {
	log.Println("register job processor", key)
	jobProcessors[key] = processor
}

func getEnabledAndNonRunningJobs() ([]job.Job, error) {
	ctx, cancel := context.WithTimeout(context.Background(), config.MongoCtxTimeout)
	defer cancel()
	return db.Find[job.Job](ctx, &bson.M{"disabled": false, "running": false}, jobCollection)
}

func Start() {
	var wg sync.WaitGroup
	for {
		jobs, err := getEnabledAndNonRunningJobs()
		if err != nil {
			log.Fatal("could not load jobs:", err)
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
		time.Sleep(time.Second)
	}
}

func process(wg *sync.WaitGroup, j *job.Job, processor job.JobProcessor) {
	nextSchedule := j.SpecificDate
	if nextSchedule.IsZero() {
		nextTick, err := gronx.NextTick(j.CronExpression, true)
		if err != nil {
			log.Println("could not get next tick: ", err)
			return
		}
		nextSchedule = nextTick
	}
	if nextSchedule.Before(time.Now()) {
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
