package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	config2 "main/config"
	"main/logging"
	"strconv"
	"strings"
)

type App interface {
	Start()
	LoadData(filePath string) error
}

type appImpl struct {
	config			*config2.Config
	logger			logging.Logger
	operations		[]string
	vehiclePark		vehiclePark
}

func NewApp(config *config2.Config, logger logging.Logger) App {
	logger.Log("Initializing the app...")
	return &appImpl{
		config: config,
		logger: logger,
	}
}

func (a *appImpl) LoadData(filePath string) error {
	rawFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		// handle any file read errors
		return fmt.Errorf("failed to read data file with error, %w", err)
	}
	// convert content to string
	content := string(rawFile)

	slotArrays, oprations, err := a.parseDataContent(content)
	if err != nil {
		return err
	}
	// set slotArrays and operations
	a.operations = oprations
	a.vehiclePark = newVehiclePark(slotArrays, a.config, a.logger)
	return nil
}

func (a *appImpl) parseDataContent(content string) (map[string][]bool, []string, error) {
	// split data into lines
	contentArr := strings.Split(content, "\n")
	// validate if there are more than 1 line
	if len(contentArr) <= 1 {
		return map[string][]bool{}, []string{}, errors.New("add at least a single operation for execution")
	}

	// split first line to get the number of slots for each type
	numberOfSlots := strings.Split(contentArr[0], " ")
	// validate if there are more than 1 line
	if len(numberOfSlots) != len(a.config.VehicleTypes) {
		return map[string][]bool{}, []string{}, errors.New("include slot numbers strictly for defined vehicle types")
	}

	// iterate through vehicle types and assign arrays
	slotArrays := map[string][]bool{}
	for i, vType := range a.config.VehicleTypes {
		slotCount, err := strconv.Atoi(numberOfSlots[i])
		if err != nil {
			return map[string][]bool{}, []string{}, errors.New("invalid data file format")
		}
		slotArrays[vType.Name] = make([]bool, slotCount)
	}

	return slotArrays, contentArr[1:], nil
}

func (a *appImpl) Start() {
	// Execute each of the operations on the vehiclePark
	for _, operation := range a.operations {
		a.vehiclePark.ExecuteCmd(operation)
	}
}
