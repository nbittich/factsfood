package openfoodfacts

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/edsrzf/mmap-go"
	"github.com/google/uuid"
	"github.com/nbittich/factsfood/config"
	"github.com/nbittich/factsfood/jobs"
	"github.com/nbittich/factsfood/types"
	jobTypes "github.com/nbittich/factsfood/types/job"
)

const (
	jobKey           = "OFF_INITIAL_SYNC_JOB"
	endpointParamKey = "endpoint"
	csvSeparatorKey  = "separator"
	gzipKey          = "gzip"
	parallelismKey   = "parallelism"
)

type InitialSync struct{}

func (is InitialSync) process(job *jobTypes.Job) (*jobTypes.JobResult, error) {
	jr := jobTypes.JobResult{
		Key:       job.Key,
		Status:    types.ERROR,
		CreatedAt: time.Now(),
		Logs:      make([]jobTypes.Log, 10),
	}
	if job.Disabled {
		return jobs.StatusError(&jr, jobTypes.DISABLED)
	}
	endpoint, ok := job.Params[endpointParamKey].(string)

	if !ok {
		jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("missing or invalid endpoint param %s: %s", endpointParamKey, job.Params)))
		return jobs.StatusError(&jr, jobTypes.INVALIDPARAM)
	}

	separator, ok := job.Params[csvSeparatorKey].(string)
	if !ok {
		jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("missing or invalid endpoint param %s: %s", csvSeparatorKey, job.Params)))
		return jobs.StatusError(&jr, jobTypes.INVALIDPARAM)
	}

	gzipped, ok := job.Params[gzipKey].(bool)
	if !ok {
		jr.Logs = append(jr.Logs, jobs.NewLog("warning! missing or invalid gzip param. fallback to false"))
		gzipped = false
	}

	parallelism, ok := job.Params[parallelismKey].(int)
	if !ok {
		jr.Logs = append(jr.Logs, jobs.NewLog("warning! missing or invalid parallelism param. fallback to thread counts"))
		parallelism = runtime.NumCPU()
	}

	tempPath := path.Join(config.TempDir, fmt.Sprintf("%s.csv", uuid.New().String()))
	jr.Logs = append(jr.Logs, jobTypes.Log{Timestamp: time.Now(), Message: fmt.Sprintf("downloading CSV from %s and saved to %s", endpoint, tempPath)})

	fileLen, err := jobs.DownloadFile(endpoint, tempPath, gzipped)
	if err != nil {
		return jobs.StatusError(&jr, err)
	}

	jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("CSV file %s downloaded, len: %d", tempPath, fileLen)))

	// CSV shenaningans
	// TODO extract to a separate func as it will be used also in delta-sync.go
	jr.Logs = append(jr.Logs, jobs.NewLog("mmap csv file"))
	f, err := os.Open(tempPath)
	if err != nil {
		return jobs.StatusError(&jr, err)
	}
	defer f.Close()
	data, err := mmap.Map(f, mmap.RDONLY, 0)
	if err != nil {
		return jobs.StatusError(&jr, err)
	}

	var offset int64 = 0
	maxChunkSize := fileLen / int64(parallelism)
	maxChunkSize -= maxChunkSize - 4096 // 4k page

	for offset < fileLen {
		end := offset + maxChunkSize
		if end > fileLen {
			end = fileLen
		}

		currentLastNewLine := offset + int64(lastNewline(data[offset:end]))

		go func() {
			log.Println(currentLastNewLine) // todo
		}()
	}

	jr.Logs = append(jr.Logs, jobs.NewLog(fmt.Sprintf("Extracting CSV using '%s' separator", separator)))

	return &jr, nil
}

func lastNewline(s []byte) int {
	i := len(s) - 1
	for i > 0 {
		if s[i] == '\n' {
			return i + 1
		}
		i -= 1
	}
	return len(s)
}
