package types

import "time"

type JobResult struct {
	ID          string     `bson:"_id" json:"_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	Status      StatusType `json:"status"`
	Logs        []string   `json:"logs"`
}

type Job struct {
	CronExpression string
	NextSchedule   time.Time
	Name           string
	Description    string
	Processor      JobProcessor
}

type JobProcessor interface {
	Process() (JobResult, error)
}
