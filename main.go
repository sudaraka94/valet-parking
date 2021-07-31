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
	// define flags
	configFilePath := flag.String("config", "config.yml", "Config file path")
	dataFilePath := flag.String("data", "datafile", "Data file path")

	//parse the flags
	flag.Parse()

	//initializes the application
	app, err := initialize(*configFilePath)
	if err != nil {
		panic(err)
	}

	// load data from data file
	err = app.LoadData(*dataFilePath)
	if err != nil {
		panic(fmt.Errorf("failed to load data from the data file with err, %w", err))
	}

	//start app execution
	app.Start()
}

// initialize inits an App instance
// with given configurations
func initialize(configPath string) (app.App, error) {
	// load config
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
	logger, err := util.InitializeLogger("cli")
	if err != nil {
		return nil, err
	}

	// initialize app instance
	return app.NewApp(config, logger), nil
}
