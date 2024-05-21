package jobs

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nbittich/factsfood/types"
	"github.com/nbittich/factsfood/types/job"
)

func DownloadFile(endpoint string, filepath string) error {
	resp, err := http.Get(endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	if ct := resp.Header.Get(echo.HeaderContentType); ct != echo.MIMEOctetStream {
		return fmt.Errorf("bad content type: %s", ct)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
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
