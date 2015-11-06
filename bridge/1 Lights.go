package hue

import (
	"fmt"
	"time"
)

// LightID ...
type LightID string

// Light ...
type Light struct {
	Type             string
	Name             string
	ModelID          string
	ManufacturerName string
	UniqueID         string
	SwVersion        string
	PointSymbol      map[string]string

	State struct {
		LightState
		Reachable bool
	}
}

// LightState ...
type LightState struct {
	On        bool
	Bri       uint8
	Hue       uint16
	Sat       uint8
	Effect    string
	Xy        [2]float64
	Ct        uint16
	Alert     string
	Colormode string
}

// GetAllLights ...
func (b Bridge) GetAllLights() (map[LightID]Light, error) {
	var lights map[LightID]Light
	err := b.get("lights", &lights)
	return lights, err
}

// GetNewLights ...
func (b Bridge) GetNewLights() (map[LightID]Light, time.Time, error) {
	var reply map[string]interface{}
	err := b.get("lights/new", &reply)

	var lastScan time.Time
	lights := map[LightID]Light{}

	for k, v := range reply {
		if k == "lastscan" {
			lastScan = convTime(v.(string))
		} else {
			var light Light
			if err := convert(v, &light); err != nil {
				return nil, time.Time{}, err
			}
			lights[LightID(k)] = light
		}
	}
	return lights, lastScan, err
}

// SearchForNewLights ...
func (b Bridge) SearchForNewLights() error {
	var reply interface{}
	err := b.post("lights", nil, &reply)
	fmt.Printf("%+v\n", reply)
	return err
}

// GetLightAttributesAndState ...
func (b Bridge) GetLightAttributesAndState(id LightID) (Light, error) {
	var light Light
	err := b.get(fmt.Sprintf("lights/%s", id), &light)
	return light, err
}

// RenameLight ...
func (b Bridge) RenameLight(id LightID, name string) error {
	return b.put(fmt.Sprintf("lights/%s", id), map[string]string{"name": name}, nil)
}

// SetLightState ...
func (b Bridge) SetLightState(id LightID, state map[string]interface{}) error {
	return b.put(fmt.Sprintf("lights/%s/state", id), state, nil)
}

// DeleteLight ...
func (b Bridge) DeleteLight(id LightID) error {
	return b.delete(fmt.Sprintf("lights/%s", id))
}
