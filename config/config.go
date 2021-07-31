package config

import (
	"errors"
)

// Config structure definition and the config validation
type Config struct {
	VehicleTypes []VehicleType `yaml:"vehicle_types"`
	LoggerConfig LoggerConfig  `yaml:"logger_config"`
}

// Validate validates Config struct
func (c *Config) Validate() error {
	// check if there are at least one vehicle type defined
	if len(c.VehicleTypes) == 0 {
		return errors.New("please define at least 1 vehicle type")
	}

	// validate each of the VehicleType objects
	for _, v := range c.VehicleTypes {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	// validate logger config
	return c.LoggerConfig.Validate()
}

type VehicleType struct {
	Name 			string	`yaml:"name"`
	PricePerHour	float64	`yaml:"price_per_hour"`
}

// Validate validates VehicleType struct
func (v *VehicleType) Validate() error {
	// return an error if the vehicle type name or price per hour is not defined
	if v.Name == "" {
		return errors.New("name of a VehicleType can't be empty")
	}

	if v.PricePerHour <= 0 {
		return errors.New("pricePerHour of a vehicleType has to be > 0")
	}
	return nil
}

type LoggerConfig struct {
	LoggerType	string	`yaml:"logger_type"`
}

func (l *LoggerConfig) Validate() error {
	// logger type is mandatory
	if l.LoggerType == "" {
		return errors.New("loggerType can't be empty")
	}
	return nil
}