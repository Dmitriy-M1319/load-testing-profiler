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

type IParser interface {
	ParseFromBytes(data []byte) (*BaseTestingMetadata, error)
	ParseFromFile(filepath string) (*BaseTestingMetadata, []byte, error)
}

type JsonParser struct{}

func NewJsonParser() *JsonParser {
	return &JsonParser{}
}

func (j *JsonParser) ParseFromFile(filepath string) (*BaseTestingMetadata, []byte, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, nil, err
	}
	result, err := j.ParseFromBytes(data)
	return result, data, nil
}

func (j *JsonParser) ParseFromBytes(data []byte) (*BaseTestingMetadata, error) {
	var result BaseTestingMetadata
	err := json.Unmarshal(data, &result)
	if err != nil {
		// TODO: Пока базовый логгер
		log.Printf("Failed to parse json from text: %s", err)
		return nil, err
	}

	return &result, nil
}
