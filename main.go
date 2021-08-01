package main

import (
	"flag"
	"fmt"
	"main/app"
	"main/util"
)

// main is the entrypoint to the application
// it parse the flags passed and initializes the application
// accordingly
func main()  {
	// define flags eg:
	// -config=config.yml -data=datafile
	configFilePath := flag.String("config", "config.yml", "Config file path")
	dataFilePath := flag.String("data", "datafile", "Data file path")

	//parse the flags
	flag.Parse()

	//initializes the application
	app, err := initialize(*configFilePath)
	if err != nil {
		fmt.Printf("failed to initialize the application with error: %s", err.Error())
		return
	}

	// load data from data file and initialize the vehicle park
	err = app.LoadData(*dataFilePath)
	if err != nil {
		fmt.Printf("failed to load data from the data file with err, %s", err.Error())
		return
	}

	//start app execution
	app.Start()
}

// initialize inits an App instance
// with given configurations
func initialize(configPath string) (app.App, error) {
	// load config from config.yml
	config, err := util.ParseConfigFile(configPath)
	if err != nil {
		return nil, err
	}

	// validate configs
	err = config.Validate()
	if err != nil {
		return nil, err
	}

	// initialize logging
	logger, err := util.InitializeLogger(config.LoggerConfig.LoggerType)
	if err != nil {
		return nil, err
	}

	// initialize app instance
	return app.NewApp(config, logger), nil
}
