package darwin

import "github.com/xiafei571/ble"

func uuidSlice(uu []ble.UUID) [][]byte {
	us := [][]byte{}
	for _, u := range uu {
		us = append(us, ble.Reverse(u))
	}
	return us
}
