package http

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/Dmitriy-M1319/load-testing-profiler/internal/runner"
)

type Runner struct {
	Metadata Metadata
}

func NewRunner(metadata Metadata) *Runner {
	return &Runner{Metadata: metadata}
}

func (r *Runner) Run(ctx context.Context) (runner.RunningInfo, error) {
	var timeoutCtx context.Context = nil
	newCtx, cancel := context.WithTimeout(ctx, time.Millisecond*time.Duration(r.Metadata.Timeout))
	defer cancel()
	timeoutCtx = newCtx

	var body io.Reader = nil
	if len(r.Metadata.Body) > 0 {
		// TODO: перегон мапы в поток байтов
	}
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, r.Metadata.Method, r.Metadata.URL, body)
	if err != nil {
		return runner.RunningInfo{Status: 500}, err
	}

	for k, v := range r.Metadata.Headers {
		req.Header.Add(k, v)
	}

	// Замер времени выполнения
	start := time.Now()
	resp, err := client.Do(req)
	end := time.Since(start)
	if err != nil {
		var result runner.RunningInfo = runner.RunningInfo{Status: 500, RequestDuration: end}
		if r.Metadata.Timeout != 0 {
			result.IsCancelled = timeoutCtx.Err() != nil
		}
		return result, err
	}

	resp.Body.Close()
	return runner.RunningInfo{Status: int32(resp.StatusCode), RequestDuration: end}, nil
}
