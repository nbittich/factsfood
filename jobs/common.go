package jobs

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/nbittich/factsfood/types"
	"github.com/nbittich/factsfood/types/job"
)

func DownloadFile(endpoint string, filepath string, gzipped bool) (int64, error) {
	resp, err := http.Get(endpoint)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bad status: %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return 0, err
	}
	defer out.Close()

	reader := resp.Body

	if gzipped {
		if reader, err = gzip.NewReader(resp.Body); err != nil {
			return 0, err
		}
	}
	defer reader.Close()

	_, err = io.Copy(out, reader)
	if err != nil {
		return 0, err
	}

	return resp.ContentLength, nil
}

func StatusError(jr *job.JobResult, err error) (*job.JobResult, error) {
	jr.Status = types.ERROR
	jr.UpdatedAt = time.Now()
	jr.Logs = append(jr.Logs, NewLog(fmt.Sprintf("error while running job %s: %v", jr.Key, err)))
	return jr, err
}

func NewLog(msg string) job.Log {
	return job.Log{
		Timestamp: time.Now(),
		Message:   msg,
	}
}
