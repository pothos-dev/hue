package hue

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Bridge ...
type Bridge struct {
	ID                string // Unused
	InternalIPAddress string
	UserID            UserID
}

// Connect ...
func Connect(internalIPAddress string, id UserID) Bridge {
	return Bridge{
		"",
		internalIPAddress,
		id,
	}
}

// DiscoverBridges ...
func DiscoverBridges() ([]Bridge, error) {
	bridges, err := discoverViaUPnP()
	if err == nil {
		return bridges, nil
	}

	bridges, err = discoverViaNUPnP()
	if err == nil {
		return bridges, nil
	}
	return bridges, nil
}

func discoverViaUPnP() ([]Bridge, error) {
	return nil, errors.New("NYI")
}

func discoverViaNUPnP() ([]Bridge, error) {
	resp, err := http.Get("https://www.meethue.com/api/nupnp")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var bridges []Bridge
	err = json.NewDecoder(resp.Body).Decode(&bridges)

	return bridges, err
}

func (b Bridge) url(resource string) string {
	return fmt.Sprintf("http://%s/api/%s/%s", b.InternalIPAddress, b.UserID, resource)
}

// get ...
func (b Bridge) get(resource string, reply interface{}) error {
	resp, err := http.Get(b.url(resource))
	if err != nil {
		return fmt.Errorf("get: http error: %s", err.Error())
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(reply)
	if err != nil {
		return fmt.Errorf("get: decoding error: %s", err.Error())
	}
	return nil
}

// post ...
func (b Bridge) post(resource string, data interface{}, reply interface{}) error {
	buf := bytes.NewBuffer([]byte{})
	if data != nil {
		encoded, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("post: encoding error: %s", err.Error())
		}
		buf.Write(encoded)
	}

	resp, err := http.Post(b.url(resource), "application/json", buf)
	if err != nil {
		return fmt.Errorf("post: http error: %s", err.Error())
	}
	defer resp.Body.Close()

	if reply != nil {
		var m []struct {
			Success interface{}
			Error   *struct {
				Type        int
				Address     string
				Description string
			}
		}
		err = json.NewDecoder(resp.Body).Decode(&m)
		if err != nil {
			return fmt.Errorf("post: decoding error: %s", err.Error())
		}

		if len(m) != 1 {
			return fmt.Errorf("post: Expected 1 result, got %d", len(m))
		}

		if e := m[0].Error; e != nil {
			return fmt.Errorf("post: '%s' returned error %d: %s", e.Address, e.Type, e.Description)
		}

		err = convert(m[0].Success, reply)
		if err != nil {
			return fmt.Errorf("post: conversion error: %s", err.Error())
		}
	}

	return nil
}

// put ...
func (b Bridge) put(resource string, data interface{}, reply interface{}) error {
	buf := bytes.NewBuffer([]byte{})
	if data != nil {
		encoded, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("put: encoding error: %s", err.Error())
		}
		buf.Write(encoded)
	}

	req, err := http.NewRequest("put", b.url(resource), buf)
	if err != nil {
		return fmt.Errorf("put: http error: %s", err.Error())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("put: http error: %s", err.Error())
	}
	defer resp.Body.Close()

	if reply != nil {
		var m []struct {
			Success interface{}
			Error   *struct {
				Type        int
				Address     string
				Description string
			}
		}
		err = json.NewDecoder(resp.Body).Decode(&m)
		if err != nil {
			return fmt.Errorf("put: decoding error: %s", err.Error())
		}

		if len(m) != 1 {
			return fmt.Errorf("put: Expected 1 result, got %d", len(m))
		}

		if e := m[0].Error; e != nil {
			return fmt.Errorf("put: '%s' returned error %d: %s", e.Address, e.Type, e.Description)
		}

		err = convert(m[0].Success, reply)
		if err != nil {
			return fmt.Errorf("put: conversion error: %s", err.Error())
		}
	}

	return nil
}

// delete ...
func (b Bridge) delete(resource string) error {
	req, err := http.NewRequest("delete", b.url(resource), bytes.NewBuffer([]byte{}))
	if err != nil {
		return fmt.Errorf("delete: http error: %s", err.Error())
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("delete: http error: %s", err.Error())
	}
	resp.Body.Close()

	return err
}

func convert(in, out interface{}) error {
	buf := bytes.NewBuffer([]byte{})
	err := json.NewEncoder(buf).Encode(in)
	if err != nil {
		return err
	}
	err = json.NewDecoder(buf).Decode(out)
	if err != nil {
		return err
	}
	return nil
}

func convTime(str string) time.Time {
	panic("TODO")
}
