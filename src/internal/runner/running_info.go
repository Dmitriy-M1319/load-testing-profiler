package runner

import (
	"context"
	"time"
)

type RunningInfo struct {
	Status          int32
	RequestDuration time.Duration
	IsCancelled     bool
}

type IRunner interface {
	Run(ctx context.Context) (RunningInfo, error)
}
