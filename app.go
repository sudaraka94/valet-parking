package main

import (
	"main/logging"
)

type App interface {
	Start(filePath string)
}

type appImpl struct {
	config *Config
	logger logging.Logger
}

func NewApp(config *Config, logger logging.Logger) App {
	logger.Log("Initializing the app...")
	return &appImpl{
		config: config,
		logger: logger,
	}
}

func (a *appImpl) Start(filePath string) {

}
