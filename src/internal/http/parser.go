package http

import (
	"encoding/json"
	"os"

	"github.com/Dmitriy-M1319/load-testing-profiler/internal/parser"
	"github.com/rs/zerolog"
)

type Metadata struct {
	parser.BaseTestingMetadata
	Method      string            `json:"method"`
	Headers     map[string]string `json:"headers"`
	Body        map[string]string `json:"body"`
	QueryParams map[string]any    `json:"params"`
}

type IHttpParser interface {
	ParseFromBytes(data []byte) (*Metadata, error)
}

type JsonHttpParser struct{}

func NewJsonHttpParser() *JsonHttpParser {
	return &JsonHttpParser{}
}

func (jp *JsonHttpParser) ParseFromBytes(data []byte) (*Metadata, error) {
	var result Metadata
	err := json.Unmarshal(data, &result)
	if err != nil {
		logger := zerolog.New(os.Stdout)
		logger.Error().Msgf("Failed to parse json from text: %s", err)
		return nil, err
	}

	return &result, nil
}
