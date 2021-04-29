package att

import "github.com/xiafei571/ble"

// attr is a BLE attribute.
type attr struct {
	h    uint16
	endh uint16
	typ  ble.UUID

	v  []byte
	rh ble.ReadHandler
	wh ble.WriteHandler
}
