package dev

import (
	"github.com/xiafei571/ble"
	"github.com/xiafei571/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
