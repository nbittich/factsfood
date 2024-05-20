package openfoodfacts

import "github.com/nbittich/factsfood/types"

const (
	jobKey           = "OFF_INITIAL_SYNC_JOB"
	endpointParamKey = "OFF_CSV_GZ_ENDPOINT"
)

type InitialSync struct{}

func (is InitialSync) process(_ *types.Job) (*types.JobResult, error) {
	return nil, nil
}
