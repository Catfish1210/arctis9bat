package arctis9

import (
	"fmt"
	"time"

	"github.com/karalabe/hid"
)

type Arctis9Headset struct {
	VendorID       uint16
	DeviceID       uint16
	BatterMax      uint16
	BatteryMin     uint16
	HidInfo        *hid.DeviceInfo
	Battery        Battery
	RequestTimeout time.Duration
	DataBuffer     []byte
	Error          error
}

type Battery struct {
	Status string
	Level  int
}

var Arctis9 = Arctis9Headset{
	VendorID:   0x1038,
	DeviceID:   0x12c2,
	BatterMax:  0x9A,
	BatteryMin: 0x64,
	Battery: Battery{
		Status: "BATTERY_UNAVAILABLE",
		Level:  -1,
	},
	RequestTimeout: 5,
	DataBuffer:     []byte{},
	Error:          nil,
}

func (hs *Arctis9Headset) Init() {
	device := hid.Enumerate(hs.VendorID, hs.DeviceID)
	if len(device) == 0 {
		hs.Error = fmt.Errorf("HID device not found with vendorID: %v and DevideID: %v", hs.VendorID, hs.DeviceID)
		return
	}
	hs.HidInfo = &device[0]
}

func (hs *Arctis9Headset) GetBattery() {
	batteryRequest := []byte{0x0, 0x20}

	headset, err := hs.HidInfo.Open()
	if err != nil {
		hs.Error = fmt.Errorf("failed to open device with vendorID: %v and DevideID: %v. error: %v", hs.VendorID, hs.DeviceID, err)
		return
	}
	defer headset.Close()

	_, err = headset.Write(batteryRequest)
	if err != nil {
		hs.Error = fmt.Errorf("failed to send data to HID device with vendorID: %v and DevideID: %v", hs.VendorID, hs.DeviceID)
		return
	}
	dataCh := make(chan error, 1)
	DataBuffer := make([]byte, 12)
	go func() {
		_, err := headset.Read(DataBuffer)
		hs.DataBuffer = DataBuffer
		dataCh <- err
	}()

	select {
	case err := <-dataCh:
		if err != nil {
			hs.Battery.Status = fmt.Sprintln("BATTERY_HIDERROR: ", err)
			hs.Battery.Level = -1
			hs.Error = fmt.Errorf("error: %v", err)
			return
		}
	case <-time.After(hs.RequestTimeout * time.Second):
		hs.Battery.Status = fmt.Sprintln("BATTERY_HIDERROR: ", err)
		hs.Battery.Level = -1
		hs.Error = fmt.Errorf("data read timeout error: %v", err)
		return
	}

	if hs.DataBuffer[4] == 0x01 {
		hs.Battery.Status = "BATTERY_CHARGING"
	} else {
		hs.Battery.Status = "BATTERY_AVAILABLE"
	}

	batteryLevel := int(hs.DataBuffer[3])
	if batteryLevel > int(hs.BatterMax) {
		hs.Battery.Level = 100
	} else {
		hs.Battery.Level = (batteryLevel - int(hs.BatteryMin)) * (100) / (int(hs.BatterMax) - (int(hs.BatteryMin)))
	}
}
