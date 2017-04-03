package dev

import (
	"github.com/chenkaie/ble"
	"github.com/chenkaie/ble/linux"
)

// DefaultDevice ...
func DefaultDevice() (d ble.Device, err error) {
	return linux.NewDevice()
}
