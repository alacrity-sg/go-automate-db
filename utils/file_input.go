package utils

import (
	"gopkg.in/yaml.v3"
	"os"
)

type FileInput struct {
	Version int `yaml:"version"`
	SQL     struct {
		Type     string `yaml:"type"`
		Mode     string `yaml:"mode"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"sql"`
	Databases []struct {
		Name     string `yaml:"name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"databases"`
}

func ParseInputFile(fileName string) (any, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	input := &FileInput{}
	err = yaml.Unmarshal(data, &input)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
