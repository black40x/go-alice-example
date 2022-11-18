package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Scenario struct {
	Name         string  `yaml:"name"`
	GlobalErrors []Alice `yaml:"global-errors"`
	Dialog       []Node  `yaml:"dialog"`
	Properties   map[string]string
}

type Extract struct {
	Path     string `yaml:"path"`
	Property string `yaml:"property"`
}

type Alice struct {
	Text string `yaml:"text"`
	TTS  string `yaml:"tts"`
}

type Node struct {
	Alice   Alice    `yaml:"alice"`
	User    string   `yaml:"user"`
	Extract *Extract `yaml:"extract"`
}

func OpenScenario(fileName string) (*Scenario, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	sc := &Scenario{}
	err = yaml.Unmarshal(data, sc)
	if err != nil {
		return nil, err
	}

	return sc, nil
}
