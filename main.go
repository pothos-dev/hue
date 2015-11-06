package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/redefiance/hue/bridge"
)

func main() {
	bridge := hue.Connect("192.168.0.100", "1b30b4a42c52e30f38e4a8cc138fa497")
	lights, err := bridge.GetAllLights()
	deny(err)

	for id := range lights {
		deny(bridge.SetLightState(id, map[string]interface{}{"on": true}))
	}
}

func deny(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("error at %s: %d: %s\n", file, line, err)
		os.Exit(0)
	}
}

func assert(cond bool) {
	if !cond {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("assertion failed at %s: %d\n", file, line)
		os.Exit(0)
	}
}
