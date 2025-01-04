package server

import (
	"context"
	"log"
	"sync"

	"github.com/Dmitriy-M1319/load-testing-profiler/internal/runner"
)

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

func (s *RunningServer) StartLoadTesting(ctx context.Context, cancel context.CancelFunc) {
	wg := &sync.WaitGroup{}
	for i := 0; i < int(s.TotalCount); i++ {
		wg.Add(1)
		go func() {
			log.Println("goroutine start")
			defer wg.Done()

			info, err := s.Runner.Run(ctx)
			if err != nil {
				// TODO: залогировать по нормальному
				log.Printf("request fatal: %s", err)
			}

			s.Result <- info
			log.Println("goroutine end")
		}()
	}

	wg.Wait()
	log.Println("goroutines cancelled")
}
