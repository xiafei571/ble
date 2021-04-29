package dev

import (
	"github.com/xiafei571/ble"
	"github.com/xiafei571/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
