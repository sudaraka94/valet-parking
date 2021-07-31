package main

import (
	"flag"
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

	//start app execution
	app.Start(*dataFilePath)
}

// initialize inits an App instance
// with given configurations
func initialize(configPath string) (App, error) {
	// load config
	config, err := parseConfigFile(configPath)
	if err != nil {
		return nil, err
	}

	// validate configs
	err = config.Validate()
	if err != nil {
		return nil, err
	}

	// initialize logging
	logger, err := initializeLogger("cli")
	if err != nil {
		return nil, err
	}

	// initialize app instance
	return NewApp(config, logger), nil
}
