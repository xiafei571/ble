package darwin

import (
	"errors"
	"time"

	"github.com/xiafei571/ble/linux/hci/cmd"
	"github.com/xiafei571/ble/linux/hci/evt"
)

// SetConnectedHandler sets handler to be called when new connection is established.
func (d *Device) SetConnectedHandler(f func(evt.LEConnectionComplete)) error {
	return errors.New("Not supported")
}

// SetDisconnectedHandler sets handler to be called on disconnect.
func (d *Device) SetDisconnectedHandler(f func(evt.DisconnectionComplete)) error {
	return errors.New("Not supported")
}

// SetPeripheralRole configures the device to perform Peripheral tasks.
func (d *Device) SetPeripheralRole() error {
	d.role = 1
	return nil
}

// SetCentralRole configures the device to perform Central tasks.
func (d *Device) SetCentralRole() error {
	d.role = 0
	return nil
}

// SetDeviceID sets HCI device ID.
func (d *Device) SetDeviceID(id int) error {
	return errors.New("Not supported")
}

// SetDialerTimeout sets dialing timeout for Dialer.
func (d *Device) SetDialerTimeout(dur time.Duration) error {
	return errors.New("Not supported")
}

// SetListenerTimeout sets dialing timeout for Listener.
func (d *Device) SetListenerTimeout(dur time.Duration) error {
	return errors.New("Not supported")
}

// SetConnParams overrides default connection parameters.
func (d *Device) SetConnParams(param cmd.LECreateConnection) error {
	return errors.New("Not supported")
}

// SetScanParams overrides default scanning parameters.
func (d *Device) SetScanParams(param cmd.LESetScanParameters) error {
	return errors.New("Not supported")
}

// SetAdvParams overrides default advertising parameters.
func (d *Device) SetAdvParams(param cmd.LESetAdvertisingParameters) error {
	return errors.New("Not supported")
}
