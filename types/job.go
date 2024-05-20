package types

import "time"

type JobResult struct {
	ID        string     `bson:"_id" json:"_id"`
	Key       string     `json:"key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Status    StatusType `json:"status"`
	Logs      []Log      `json:"logs"`
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
	NextSchedule   time.Time `json:"nextSchedule"`
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
