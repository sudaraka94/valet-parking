package app

import (
	"errors"
	"fmt"
	"main/config"
	"main/logging"
	"main/util"
	"strconv"
	"strings"
)

// type definitions
type vehicleInfo struct {
	SlotNumber	int
	Timestamp	int
	VehicleType	string
}

type vehiclePark interface {
	ExecuteCmd(command string) error
}

type vehicleParkImpl struct {
	vehicleMap   map[string]vehicleInfo
	slotArrayMap map[string][]bool
	fareMap      map[string]float64
	logger       logging.Logger
}

// newVehiclePark is the constructor which encapsulates the internals of vehicle park initialization
func newVehiclePark(slotArrayMap map[string][]bool, config *config.Config, logger logging.Logger) vehiclePark {
	// prepare fareMap, so that lookup complexity can be reduced to O(1)
	// could be improved with directly loading fares as a map from config.yml
	// used float as the datatype for price per hour property to allow decimal points in case required
	fareMap := map[string]float64{}
	for _, vType := range config.VehicleTypes {
		fareMap[vType.Name] = vType.PricePerHour
	}

	return &vehicleParkImpl{
		vehicleMap:   map[string]vehicleInfo{},
		slotArrayMap: slotArrayMap,
		fareMap:      fareMap,
		logger:       logger,

	}
}

// ExecuteCmd calls execEnter or execExit depending on the fist word of the operation
func (v *vehicleParkImpl) ExecuteCmd(command string) error {
	ops := strings.Split(command, " ")
	switch ops[0] {
	case "Enter":
		return v.execEnter(ops[1], ops[2], ops[3])
	case "Exit":
		return v.execExit(ops[1], ops[2])
	default:
		return errors.New("invalid command")
	}
}

// execEnter simulates the scenario of accepting a vehicle to the vehicle park
func (v *vehicleParkImpl) execEnter(vehicleType string, regNo string, rawTimestamp string) error {
	if timestamp, err := strconv.Atoi(rawTimestamp); err == nil {
		lotNumber, err := v.addVehicle(vehicleType, regNo, timestamp)
		if err != nil {
			return err
		}

		if lotNumber < 0 {
			v.logger.Logf("Reject\n")
		} else {
			// convert the array index to lot number by adding 1 (lot numbers start with 1)
			v.logger.Logf("Accept %sLot%d\n", strings.Title(vehicleType), lotNumber + 1)
		}
	} else {
		return fmt.Errorf("failed to read the timestamp %s", rawTimestamp)
	}
	return nil
}


func (v *vehicleParkImpl) execExit(regNo string, rawTimestamp string) error {
	if timestamp, err := strconv.Atoi(rawTimestamp); err == nil {
		info, err := v.removeVehicle(regNo)
		if err != nil {
			return err
		}

		fare, err := v.calculateFare(info.Timestamp, timestamp, info.VehicleType)
		if err != nil {
			return err
		}
		// I am printing floats without any decimal points since that's what I saw in the example given
		v.logger.Logf("%sLot%d %.f\n", strings.Title(info.VehicleType), info.SlotNumber + 1, fare)
	} else {
		errors.New("failed to execute command")
	}
	return nil
}

// function imported this to allow mocking package level functions
var findAndUpdate = util.FindAndUpdate

func (v *vehicleParkImpl) addVehicle(vehicleType string, regNo string, timestamp int) (int, error) {
	// select slot array by vehicle type
	slotArray, ok := v.slotArrayMap[vehicleType]
	if !ok {
		return 0, errors.New("invalid vehicle type")
	}

	// get the least numbered slot
	slotNumber := findAndUpdate(slotArray)

	// update info if parking is not full
	if slotNumber >= 0 {
		// Update the hashmap for tracking
		v.vehicleMap[regNo] = vehicleInfo{
			SlotNumber: slotNumber,
			Timestamp: timestamp,
			VehicleType: vehicleType,
		}
	}

	return slotNumber, nil
}

func (v *vehicleParkImpl) removeVehicle(regNo string) (vehicleInfo, error) {
	// get vehicle record from the map
	currentVehicleInfo, ok := v.vehicleMap[regNo]
	if !ok {
		return vehicleInfo{}, errors.New("vehicle not found")
	}

	// clear the vehicle info from the map
	delete(v.vehicleMap, regNo)

	// get the slotArray
	slotArray := v.slotArrayMap[currentVehicleInfo.VehicleType]

	// update the slotArray
	slotArray[currentVehicleInfo.SlotNumber] = false

	return currentVehicleInfo, nil
}

// using this method so that I'm able to mock the package level function
var calculateTime = util.CalculateTime
func (v *vehicleParkImpl) calculateFare(start int, end int, vehicleType string) (float64, error) {
	time, err := calculateTime(start, end)
	if err != nil {
		return 0, err
	}

	fare, ok := v.fareMap[vehicleType]
	if !ok {
		return 0, errors.New("invalid vehicle type")
	}

	return float64(time) * fare, nil
}