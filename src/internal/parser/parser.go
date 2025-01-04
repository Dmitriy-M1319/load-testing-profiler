package parser

import (
	"encoding/json"
	"log"
	"os"
)

type BaseTestingMetadata struct {
	Type        string   `json:"type"`
	URL         string   `json:"url"`
	AuthData    []string `json:"auth_data"`
	TesterCount int64    `json:"tester_count"`
	Timeout     int32    `json:"timeout"`
}

type HttpTestingMetadata struct {
	BaseTestingMetadata
	Method      string            `json:"method"`
	Headers     map[string]string `json:"headers"`
	Body        map[string]string `json:"body"`
	QueryParams map[string]any    `json:"params"`
}

type IHttpParser interface {
	ParseFromText(text string) (*HttpTestingMetadata, error)
	ParseFromFile(filepath string) (*HttpTestingMetadata, error)
}

type JsonHttpParser struct{}

func NewJsonHttpParser() *JsonHttpParser {
	return &JsonHttpParser{}
}

func (jp *JsonHttpParser) ParseFromText(text string) (*HttpTestingMetadata, error) {
	var result HttpTestingMetadata
	err := json.Unmarshal([]byte(text), &result)
	if err != nil {
		// TODO: Пока базовый логгер
		log.Printf("Failed to parse json from text: %s", err)
		return nil, err
	}

	return &result, nil
}

func (jp *JsonHttpParser) ParseFromFile(filepath string) (*HttpTestingMetadata, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	result, err := jp.ParseFromText(string(data))
	return result, nil
}
