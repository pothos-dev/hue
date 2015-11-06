package hue

import "time"

// SensorID ...
type SensorID string

// Sensor ...
type Sensor struct {
	State struct {
		DayLight    bool
		Lastupdated string
	}
	Config struct {
		On            bool
		Long          string
		Lat           string
		RunriseOffset int
		SunsetOffset  int
	}
	Name             string
	Type             string
	ModelID          string
	Manufacturername string
	SwVersion        string
}

// GetAllSensors ...
func (b Bridge) GetAllSensors() (map[SensorID]Sensor, error) {
	panic("TODO")
}

// CreateSensor ...
func (b Bridge) CreateSensor(attributes map[string]interface{}) (SensorID, error) {
	panic("TODO")
}

// FindNewSensors ...
func (b Bridge) FindNewSensors() error {
	panic("TODO")
}

// GetNewSensors ...
func (b Bridge) GetNewSensors() (map[SensorID]Sensor, time.Time, error) {
	panic("TODO")
}

// GetSensor ...
func (b Bridge) GetSensor(id SensorID) (Sensor, error) {
	panic("TODO")
}

// DeleteSensor ...
func (b Bridge) DeleteSensor(id SensorID) error {
	panic("TODO")
}

// ChangeSensorConfig ...
func (b Bridge) ChangeSensorConfig(id SensorID, state map[string]interface{}) error {
	panic("TODO")
}
