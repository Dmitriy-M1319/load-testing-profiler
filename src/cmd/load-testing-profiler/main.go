package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Dmitriy-M1319/load-testing-profiler/internal/http"
	"github.com/Dmitriy-M1319/load-testing-profiler/internal/parser"
	"github.com/Dmitriy-M1319/load-testing-profiler/internal/runner"
	"github.com/Dmitriy-M1319/load-testing-profiler/internal/server"
)

var hParser http.IHttpParser
var bParser parser.IParser
var run runner.IRunner

func main() {
	errMsg := "usage: load-testing-profiler <setup_file>"
	if len(os.Args) < 2 {
		log.Fatal(errMsg)
	}

	//filename := "../data.json"
	filename := os.Args[1]
	bParser = parser.NewJsonParser()
	hParser = http.NewJsonHttpParser()
	preparedData, bSlice, err := bParser.ParseFromFile(filename)

	if err != nil {
		log.Fatalf("Failed to parse metadata file: %s", err)
	}
	fmt.Println(preparedData)

	if preparedData.Type == "http" {
		httpData, err := hParser.ParseFromBytes(bSlice)
		if err != nil {
			log.Fatalf("Failed to parse http metadata: %s", err)
		}
		run = http.NewRunner(*httpData)
	}

	ctx := context.Background()

	srv := server.NewRunningServer(int32(preparedData.TesterCount), run)
	srv.StartLoadTesting(ctx)

	for res := range srv.Result {
		log.Printf("Result: %d, duration: %dms", res.Status, res.RequestDuration.Milliseconds())
	}
}
