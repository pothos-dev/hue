package fhue

import (
	"github.com/redefiance/hue/bridge"
)

// Light ...
type Light struct{}

// Command ...
type Command struct {
	bridge hue.Bridge
	action string
	data   map[string]interface{}
}

// Execute ...
func (c *Command) Execute() {
}
