package server

import (
	"context"
	"os"
	"sync"

	"github.com/Dmitriy-M1319/load-testing-profiler/internal/runner"
	"github.com/rs/zerolog"
)

type threadLogger struct {
	logger zerolog.Logger
	mutex  sync.Mutex
}

var logger threadLogger = threadLogger{logger: zerolog.New(os.Stdout)}

type RunningServer struct {
	TotalCount int32
	Runner     runner.IRunner
	Result     chan runner.RunningInfo
}

func NewRunningServer(totalCount int32, r runner.IRunner) *RunningServer {
	return &RunningServer{
		TotalCount: totalCount,
		Runner:     r,
		Result:     make(chan runner.RunningInfo, totalCount)}
}

func (s *RunningServer) StartLoadTesting(ctx context.Context) {
	wg := &sync.WaitGroup{}
	for i := 0; i < int(s.TotalCount); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			info, err := s.Runner.Run(ctx)
			if err != nil {
				logger.mutex.Lock()
				logger.logger.Error().Msgf("Error: %s", err)
				logger.mutex.Unlock()
			}

			s.Result <- info
		}()
	}

	wg.Wait()
	close(s.Result)
}
