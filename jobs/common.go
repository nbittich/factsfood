package jobs

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/nbittich/factsfood/types"
	"github.com/nbittich/factsfood/types/job"
)

func printProgressDownloadFile(done chan int64, path string, total int64) {
	stop := false
	// give it some time to kick in
	log.Println("Download started.")
	time.Sleep(time.Second * 3)
	file, err := os.Open(path)
	if err != nil {
		log.Println("cannot show progress bar:", err)
		return
	}
	for {
		select {
		case downloaded := <-done:
			if total != downloaded {
				fmt.Printf("warning! expected total: %d, actual total: %d\n", total, downloaded)
			}
			stop = true
			fmt.Println()
		default:
			fi, err := file.Stat()
			if err != nil {
				continue
			}
			size := fi.Size()
			percent := float64(size) / float64(total) * 100

			fmt.Printf("\rDownload Progress: %.0f%%", percent)
		}
		time.Sleep(time.Second * 1)
		if stop {
			break
		}
	}
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

	reader := resp.Body

	if gzipped {
		if reader, err = gzip.NewReader(resp.Body); err != nil {
			return 0, err
		}
	}
	defer reader.Close()
	done := make(chan int64)
	go printProgressDownloadFile(done, filepath, resp.ContentLength)
	s, err := io.Copy(out, reader)
	done <- s
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
	log.Println(msg)
	return job.Log{
		Timestamp: time.Now(),
		Message:   msg,
	}
}
