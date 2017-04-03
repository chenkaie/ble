package dev

import (
	"github.com/chenkaie/ble"
	"github.com/chenkaie/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return darwin.NewDevice()
}
