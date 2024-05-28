package job

import (
	"time"

	"github.com/nbittich/factsfood/types"
)

type JobError uint8

const (
	DISABLED JobError = iota + 1
	INVALIDPARAM
	BADKEY
)

func (j JobError) Error() string {
	switch j {
	case DISABLED:
		return "job disabled"
	case INVALIDPARAM:
		return "missing or invalid param in job config"
	case BADKEY:
		return "bad key"
	default:
		return "unknown job error"

	}
}

type JobResult struct {
	ID        string                 `bson:"_id" json:"_id"`
	Key       string                 `json:"key"`
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
	Status    types.StatusType       `json:"status"`
	Logs      []Log                  `json:"logs"`
	Metadata  map[string]interface{} `json:"metadata"`
}

type Log struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

func (jobResult JobResult) GetID() string {
	return jobResult.ID
}

func (jobResult *JobResult) SetID(id string) {
	jobResult.ID = id
}

type JobParams map[string]interface{}

type Job struct {
	ID             string    `bson:"_id" json:"_id"`
	CronExpression string    `json:"cronExpression"`
	SpecificDate   time.Time `json:"specificDate"`
	NextSchedule   time.Time `json:"nextSchedule"`
	Running        bool      `json:"isRunning"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Key            string    `json:"key"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Disabled       bool      `json:"disabled"`
	Params         JobParams `json:"params"`
}

func (job Job) GetID() string {
	return job.ID
}

func (job *Job) SetID(id string) {
	job.ID = id
}

type JobProcessor interface {
	Process(job *Job) (*JobResult, error)
}
