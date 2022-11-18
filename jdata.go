package main

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

type JsonData map[string]interface{}

var ErrorInvalidPath = errors.New("invalid path")
var ErrorInvalidType = errors.New("invalid value type")
var ErrorUnmarshal = errors.New("unmarshal error")

func NewJsonData(data string) (*JsonData, error) {
	j := &JsonData{}
	err := json.Unmarshal([]byte(data), j)
	if err != nil {
		return nil, ErrorUnmarshal
	}

	return j, nil
}

func (j JsonData) Get(path string) (interface{}, error) {
	keys := strings.Split(path, ".")
	var node interface{}
	node = j
	for i, key := range keys {
		switch node.(type) {
		case JsonData:
			node = node.(JsonData)[key]
		case map[string]interface{}:
			node = node.(map[string]interface{})[key]
		case []interface{}:
			id, _ := strconv.Atoi(key)
			if id >= 0 && id <= len(node.([]interface{}))-1 {
				node = node.([]interface{})[id]
			} else {
				return nil, ErrorInvalidPath
			}
		case interface{}:
			if i == len(keys)-1 {
				return node, nil
			}
			return nil, ErrorInvalidPath
		}
	}
	return node, nil
}

func (j JsonData) GetInt(path string) (int, error) {
	v, err := j.Get(path)
	if err != nil {
		return 0, err
	}
	switch v.(type) {
	case int:
		return v.(int), nil
	default:
		return 0, ErrorInvalidType
	}
}

func (j JsonData) GetFloat32(path string) (float32, error) {
	v, err := j.Get(path)
	if err != nil {
		return 0, err
	}
	switch v.(type) {
	case float32:
		return v.(float32), nil
	default:
		return 0, ErrorInvalidType
	}
}

func (j JsonData) GetFloat64(path string) (float64, error) {
	v, err := j.Get(path)
	if err != nil {
		return 0, err
	}
	switch v.(type) {
	case float64:
		return v.(float64), nil
	default:
		return 0, ErrorInvalidType
	}
}

func (j JsonData) GetString(path string) (string, error) {
	v, err := j.Get(path)
	if err != nil {
		return "", err
	}
	switch v.(type) {
	case string:
		return v.(string), nil
	default:
		return "", ErrorInvalidType
	}
}

func (j JsonData) GetBool(path string) (bool, error) {
	v, err := j.Get(path)
	if err != nil {
		return false, err
	}
	switch v.(type) {
	case bool:
		return v.(bool), nil
	default:
		return false, ErrorInvalidType
	}
}

func (j JsonData) GetSlice(path string) ([]any, error) {
	v, err := j.Get(path)
	if err != nil {
		return nil, err
	}
	switch v.(type) {
	case []any:
		return v.([]any), nil
	default:
		return nil, ErrorInvalidType
	}
}
