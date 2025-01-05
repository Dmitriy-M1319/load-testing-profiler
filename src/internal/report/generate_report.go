package report

import (
	"fmt"
	"time"

	"github.com/Dmitriy-M1319/load-testing-profiler/internal/http"
	"github.com/Dmitriy-M1319/load-testing-profiler/internal/runner"
)

// Базовая печать отчета
func PrintHttpReport(rChan chan runner.RunningInfo, metadata *http.Metadata) {
	fmt.Printf("URL: %s\n", metadata.URL)
	fmt.Printf("Count of tests: %d\n", metadata.TesterCount)
	fmt.Printf("Auth: %v\n", metadata.AuthData)
	fmt.Printf("Cancel timeout: %dms\n", metadata.Timeout)
	fmt.Printf("Method: %s\n", metadata.Method)
	fmt.Printf("Headers: %v\n", metadata.Headers)
	fmt.Printf("Body: %v\n", metadata.Body)
	fmt.Printf("Query params: %v\n", metadata.QueryParams)

	statusMap := make(map[int32]int32)
	timeoutSlice := make([]time.Duration, 0, len(rChan))
	cancelledCount := 0

	for res := range rChan {
		statusMap[res.Status] = statusMap[res.Status] + 1
		timeoutSlice = append(timeoutSlice, res.RequestDuration)
		if res.IsCancelled {
			cancelledCount++
		}
	}

	fmt.Println("")
	fmt.Println("REPORT:")
	fmt.Println("1. Status codes:")
	for status, count := range statusMap {
		fmt.Printf("Code: %d - count %d/%d (%f)\n", status, count, cap(rChan), float64(count)/float64(cap(rChan)))
	}

	fmt.Println("")
	fmt.Println("2. Request Durations:")

	minDuration := func(timeouts []time.Duration) time.Duration {
		min := timeouts[0]
		for _, t := range timeouts {
			if t < min {
				min = t
			}
		}
		return min
	}

	maxDuration := func(timeouts []time.Duration) time.Duration {
		max := timeouts[0]
		for _, t := range timeouts {
			if t > max {
				max = t
			}
		}
		return max
	}

	fmt.Printf("Min: %d ms\n", minDuration(timeoutSlice).Milliseconds())
	fmt.Printf("Max: %d ms\n", maxDuration(timeoutSlice).Milliseconds())

	fmt.Println("")
	fmt.Println("3. Cancelled:")
	fmt.Printf("Count: %d/%d\n", cancelledCount, cap(rChan))
}
