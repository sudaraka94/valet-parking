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
	return &appImpl{
		config: config,
		logger: logger,
	}
}

// LoadData reads the contents of the data file located in the given
// filepath and parse the content
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

// parseDataContent parses the stringified contents of the datafile into
// number of vehicle slots and an array of operations to execute
func (a *appImpl) parseDataContent(content string) (map[string][]bool, []string, error) {
	// split data by line breaks
	// first line contains the number of slots and other lines contains operations
	contentArr := strings.Split(content, "\n")
	// there has to be at least a single operation
	if len(contentArr) <= 1 {
		return map[string][]bool{}, []string{}, errors.New("add at least a single operation for execution")
	}

	// split first line to get the number of slots for each vehicle type
	numberOfSlots := strings.Split(contentArr[0], " ")
	// validate if the number of vehicle types defined equals to the number of slot sizes
	if len(numberOfSlots) != len(a.config.VehicleTypes) {
		return map[string][]bool{}, []string{}, errors.New("include slot numbers strictly for defined vehicle types")
	}

	// iterate through vehicle types and assign arrays to keep track of slot occupancy
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

// Start Executes each of the operations on the vehiclePark serially
func (a *appImpl) Start() {
	for _, operation := range a.operations {
		a.vehiclePark.ExecuteCmd(operation)
	}
}
