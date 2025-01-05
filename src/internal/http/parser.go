package http

import (
	"encoding/json"
	"log"

	"github.com/Dmitriy-M1319/load-testing-profiler/internal/parser"
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
		// TODO: Пока базовый логгер
		log.Printf("Failed to parse json from text: %s", err)
		return nil, err
	}

	return &result, nil
}
