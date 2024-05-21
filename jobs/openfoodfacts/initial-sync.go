package openfoodfacts

import (
	"fmt"
	"path"
	"time"

	"github.com/google/uuid"
	"github.com/nbittich/factsfood/config"
	"github.com/nbittich/factsfood/jobs"
	"github.com/nbittich/factsfood/types"
	jobTypes "github.com/nbittich/factsfood/types/job"
)

const (
	jobKey           = "OFF_INITIAL_SYNC_JOB"
	endpointParamKey = "endpoint"
)

type InitialSync struct{}

func (is InitialSync) process(job *jobTypes.Job) (*jobTypes.JobResult, error) {
	jr := jobTypes.JobResult{
		Key:       job.Key,
		Status:    types.SUCCESS,
		CreatedAt: time.Now(),
		Logs:      make([]jobTypes.Log, 10),
	}
	if job.Disabled {
		return nil, jobTypes.DISABLED
	}
	endpoint, ok := job.Params[endpointParamKey].(string)

	if !ok {
		fmt.Println("missing or invalid endpoint param: ", job.Params)
		return nil, jobTypes.INVALIDPARAM
	}

	tempPath := path.Join(config.TempDir, fmt.Sprintf("%s.csv.gz", uuid.New().String()))
	jr.Logs = append(jr.Logs, jobTypes.Log{Timestamp: time.Now(), Message: fmt.Sprintf("downloading CSV from %s and saved to %s", endpoint, tempPath)})

	if err := jobs.DownloadFile(endpoint, tempPath); err != nil {
		return jobs.StatusError(&jr, err)
	}

	return &jr, nil
}
