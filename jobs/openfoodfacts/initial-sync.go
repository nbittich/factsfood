package openfoodfacts

import (
	"fmt"

	"github.com/nbittich/factsfood/types"
)

const (
	jobKey           = "OFF_INITIAL_SYNC_JOB"
	endpointParamKey = "endpoint"
)

type InitialSync struct{}

func (is InitialSync) process(job *types.Job) (*types.JobResult, error) {
	if job.Disabled {
		return nil, types.DISABLED
	}
	endpoint, ok := job.Params[endpointParamKey].(string)
	if !ok {
		fmt.Println("missing or invalid endpoint param: ", job.Params)
		return nil, types.INVALIDPARAM
	}
	fmt.Println(endpoint)
	return nil, nil
}
