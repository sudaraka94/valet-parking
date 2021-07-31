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
	vehicleMap		map[string]vehicleInfo
	slotArrayMap	map[string][]bool
	fairMap			map[string]float64
	logger			logging.Logger
}

// newVehiclePark
func newVehiclePark(slotArrayMap map[string][]bool, config *config.Config, logger logging.Logger) vehiclePark {
	// prepare fairMap, so that lookup complexity can be reduced to O(1)
	fairMap := map[string]float64{}
	for _, vType := range config.VehicleTypes {
		fairMap[vType.Name] = vType.PricePerHour
	}

	return &vehicleParkImpl{
		vehicleMap: map[string]vehicleInfo{},
		slotArrayMap: slotArrayMap,
		fairMap: fairMap,
		logger: logger,

	}
}

func (v *vehicleParkImpl) execEnter(vehicleType string, regNo string, rawTimestamp string) error {
	if timestamp, err := strconv.Atoi(rawTimestamp); err == nil {
		lotNumber, err := v.addVehicle(vehicleType, regNo, timestamp)
		if err != nil {
			return err
		}

		if lotNumber < 0 {
			fmt.Println("Reject")
		} else {
			fmt.Printf("Accept %sLot%d\n", strings.Title(vehicleType), lotNumber + 1)
		}
	} else {
		return errors.New("failed to execute command")
	}
	return nil
}

func (v *vehicleParkImpl) execExit(regNo string, rawTimestamp string) error {
	if timestamp, err := strconv.Atoi(rawTimestamp); err == nil {
		info := v.removeVehicle(regNo)
		fair, err := v.calculateFair(info.Timestamp, timestamp, info.VehicleType)
		if err != nil {
			return err
		}

		v.logger.Logf("%sLot%d %d\n", strings.Title(info.VehicleType), info.SlotNumber + 1, fair)
	} else {
		errors.New("failed to execute command")
	}
	return nil
}

func (v *vehicleParkImpl) addVehicle(vehicleType string, regNo string, timestamp int) (int, error) {
	// select slot array by vehicle type
	slotArray, ok := v.slotArrayMap[vehicleType]
	if !ok {
		return 0, errors.New("invalid vehicle type")
	}

	// get the least numbered slot
	slotNumber := util.FindAndUpdate(slotArray)

	// check if parking is full
	if slotNumber == -1 {
		return -1, nil
	}

	// Update the hashmap for tracking
	v.vehicleMap[regNo] = vehicleInfo{
		SlotNumber: slotNumber,
		Timestamp: timestamp,
		VehicleType: vehicleType,
	}

	return slotNumber, nil
}

func (v *vehicleParkImpl) removeVehicle(regNo string) vehicleInfo {
	// get vehicle record from the map
	vehicleInfo := v.vehicleMap[regNo]

	// clear the vehicle info from the map
	delete(v.vehicleMap, regNo)

	// get the slotArray
	slotArray := v.slotArrayMap[vehicleInfo.VehicleType]

	// update the slotArray
	slotArray[vehicleInfo.SlotNumber] = false

	return vehicleInfo
}

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

func (v *vehicleParkImpl) calculateFair(start int, end int, vehicleType string) (float64, error) {
	time, err := util.CalculateTime(start, end)
	if err != nil {
		return 0, err
	}

	fair, ok := v.fairMap[vehicleType]
	if !ok {
		return 0, errors.New("invalid vehicle type")
	}

	return float64(time) * fair, nil
}