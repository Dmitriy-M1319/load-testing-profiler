package runner

import (
	"context"
	"io"
	"net/http"

	"github.com/Dmitriy-M1319/load-testing-profiler/internal/parser"
)

type RunningInfo struct {
	Status int32
}

type IRunner interface {
	Run(ctx context.Context) (RunningInfo, error)
}

type HttpRunner struct {
	Metadata parser.HttpTestingMetadata
}

func NewHttpRunner(metadata parser.HttpTestingMetadata) *HttpRunner {
	return &HttpRunner{Metadata: metadata}
}

func (r *HttpRunner) Run(ctx context.Context) (RunningInfo, error) {
	var body io.Reader = nil
	if len(r.Metadata.Body) > 0 {
		// TODO: перегон мапы в поток байтов
	}
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, r.Metadata.Method, r.Metadata.URL, body)
	if err != nil {
		return RunningInfo{Status: 500}, err
	}

	for k, v := range r.Metadata.Headers {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return RunningInfo{Status: 500}, err
	}
	return RunningInfo{Status: int32(resp.StatusCode)}, nil
}
