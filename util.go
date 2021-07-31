package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"main/logging"
)

// parseConfigFile reads the yml file in the given path
// and unmarshalls the content in to a Config struct
func parseConfigFile(configFilePath string) (*Config, error) {
	rawFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		// handle any file read errors
		return nil, fmt.Errorf("failed to read config with error, %w", err)
	}

	// try and parse yml the content into Config struct
	var config Config
	err = yaml.Unmarshal(rawFile, &config)
	if err != nil {
		// handle the errors while parsing yml
		return nil, fmt.Errorf("failed to parse config file: %s, with error, %w", configFilePath, err)
	}

	return &config, nil
}

// initializeLogger returns a logger depending on the logger type
func initializeLogger(loggerType string) (logging.Logger, error) {
	switch loggerType {
	case "cli":
		return logging.NewCLILogger(), nil
	default:
		return nil, errors.New("invalid logger type")
	}
}