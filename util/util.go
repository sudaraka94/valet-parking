package util

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"main/config"
	"main/logging"
)

// ParseConfigFile reads the yml file in the given path
// and unmarshalls the content in to a config.Config struct
func ParseConfigFile(configFilePath string) (*config.Config, error) {
	rawFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		// handle any file read errors
		return nil, fmt.Errorf("failed to read config with error, %w", err)
	}

	// try and parse yml the content into Config struct
	var conf config.Config
	err = yaml.Unmarshal(rawFile, &conf)
	if err != nil {
		// handle the errors while parsing yml
		return nil, fmt.Errorf("failed to parse config file: %s, with error, %w", configFilePath, err)
	}

	return &conf, nil
}

// InitializeLogger returns a logger depending on the logger type
// allowed loggerTypes : cli
func InitializeLogger(loggerType string) (logging.Logger, error) {
	switch loggerType {
	case "cli":
		return logging.NewCLILogger(), nil
	default:
		return nil, errors.New("invalid logger type")
	}
}

// FindAndUpdate findAndUpdate finds the least index with -1 as value
// O(n) time complexity
func FindAndUpdate(arr []bool) int {
	for i, value := range arr {
		if !value {
			arr[i] = true
			return i
		}
	}
	return -1
}

// CalculateTime Calculates the time gap between given
// start and end timestamps. Outputs the gap in hours rounded to nearest hour
// eg: 1.2 hrs = 2hrs
func CalculateTime(start int, end int) (int, error) {
	if (end - start) < 0 {
		return 0, errors.New("invalid time input")
	}

	timeInSec := end - start
	timeInHours := timeInSec / (60*60)
	remainder := timeInSec % (60*60)
	if remainder>0 {
		return timeInHours + 1, nil
	}

	return timeInHours, nil
}