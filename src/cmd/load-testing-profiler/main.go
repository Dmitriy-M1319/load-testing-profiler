package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Dmitriy-M1319/load-testing-profiler/internal/parser"
	"github.com/Dmitriy-M1319/load-testing-profiler/internal/runner"
	"github.com/Dmitriy-M1319/load-testing-profiler/internal/server"
)

var pr parser.IHttpParser
var run runner.IRunner

func main() {
	errMsg := "usage: load-testing-profiler <setup_file>"
	if len(os.Args) < 2 {
		log.Fatal(errMsg)
	}

	filename := os.Args[1]
	pr = parser.NewJsonHttpParser()
	preparedData, err := pr.ParseFromFile(filename)

	if err != nil {
		log.Fatalf("Failed to parse metadata file: %s", err)
	}

	run = runner.NewHttpRunner(*preparedData)
	ctx := context.Background()
	var cancel context.CancelFunc
	if preparedData.Timeout != 0 {
		ctx, cancel = context.WithTimeout(ctx, time.Second*time.Duration(preparedData.Timeout))
	}

	srv := server.NewRunningServer(int32(preparedData.TesterCount), run)
	srv.StartLoadTesting(ctx, cancel)

	for res := range srv.Result {
		log.Printf("Result: %d", res.Status)
	}

	close(srv.Result)
}
