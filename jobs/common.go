package jobs

import (
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/mitchellh/mapstructure"
	"github.com/nbittich/factsfood/config"
	"github.com/nbittich/factsfood/services/utils"
	"github.com/nbittich/factsfood/types"
	"github.com/nbittich/factsfood/types/job"
	"golang.org/x/time/rate"
)

type progress struct {
	startedAt time.Time
	total     uint64
}

func (p *progress) Write(b []byte) (int, error) {
	duration := uint64(time.Since(p.startedAt).Seconds())
	p.total += uint64(len(b))
	if duration <= 0 {
		duration = 1
	}
	inSec := p.total / duration
	fmt.Printf("\033[2K\rDownloaded %s...(rate %s/s)", humanize.Bytes(p.total), humanize.Bytes(inSec))
	return len(b), nil
}

type ThrottledReader struct {
	r       io.Reader
	limiter *rate.Limiter
}

func (tr *ThrottledReader) Read(p []byte) (int, error) {
	n, err := tr.r.Read(p)
	if err != nil {
		return n, err
	}
	err = tr.limiter.WaitN(context.TODO(), n)
	return n, err
}

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

	defer resp.Body.Close()

	var rateLimit int
	var reader io.Reader

	if gzipped {
		if reader, err = gzip.NewReader(resp.Body); err != nil {
			return 0, err
		}
		rateLimit = config.HTTPGzippedDownloadRateLimitInMegaBytes * 1024 * 1024
	} else {
		rateLimit = config.HTTPDownloadRateLimitInMegaBytes * 1024 * 1024
		reader = resp.Body
	}

	limiter := rate.NewLimiter(rate.Limit(rateLimit), rateLimit)

	reader = &ThrottledReader{
		r:       reader,
		limiter: limiter,
	}

	progressReader := io.TeeReader(reader, &progress{startedAt: time.Now()})
	contentLength, err := io.Copy(out, progressReader)
	if err != nil {
		return 0, err
	}

	fmt.Println()
	return contentLength, nil
}

func StatusError(jr *job.JobResult, err error) (*job.JobResult, error) {
	jr.Status = types.ERROR
	jr.UpdatedAt = time.Now()
	jr.Logs = append(jr.Logs, NewLog(fmt.Sprintf("error while running job %s: %v", jr.Key, err)))
	return jr, err
}

func NewLog(msg string) job.Log {
	log.Println(msg)
	return job.Log{
		Timestamp: time.Now(),
		Message:   msg,
	}
}

func ValidateJobAndGetParam[T any](
	j *job.Job,
	jr *job.JobResult,
	modifier func(t *T) (*T, error),
	validKeys ...string,
) (*T, error) {
	if j.Disabled {
		return nil, job.DISABLED
	}
	if !slices.Contains(validKeys, j.Key) {
		return nil, job.BADKEY
	}
	it := new(T)

	if err := mapstructure.Decode(j.Params, it); err != nil {
		jr.Logs = append(jr.Logs, NewLog(fmt.Sprintf("error: %s", err)))
		return nil, job.INVALIDPARAM
	}
	if err := utils.ValidateStruct(it); err != nil {
		jr.Logs = append(jr.Logs, NewLog(fmt.Sprintf("error: %s", err)))
		return nil, job.INVALIDPARAM
	}

	if modifier != nil {
		return modifier(it)
	}

	return it, nil
}
