package utils

import (
	"os"

	"github.com/go-yaml/yaml"
)

type Configuration struct {
	DbFileName  string
	ServiceName string
}

func LoadConfiguration(fileName string) Configuration {
	var config Configuration
	var confFileContent, err = os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(confFileContent, config)
	if err != nil {
		panic(err)
	}

	return config
}
