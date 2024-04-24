package govoom

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type findDeviceResult struct {
	ReturnCode    int       `json:"ReturnCode"`
	ReturnMessage string    `json:"ReturnMessage"`
	DeviceList    []*Device `json:"DeviceList"`
}

type Device struct {
	DeviceName      string `json:"DeviceName"`
	DeviceID        int    `json:"DeviceId"`
	DevicePrivateIP string `json:"DevicePrivateIP"`
	DeviceMAC       string `json:"DeviceMac"`
}

func (d *Device) GetClient() *Client {
	return NewClient(d.DevicePrivateIP)
}

func FindDevices() ([]*Device, error) {
	url := "https://app.divoom-gz.com/Device/ReturnSameLANDevice"
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data findDeviceResult
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	if data.ReturnCode != 0 {
		return nil, fmt.Errorf("failed to find devices: %s", data.ReturnMessage)
	}

	return data.DeviceList, err
}

func FindDeviceByMac(mac string) (*Device, error) {
	devices, err := FindDevices()
	if err != nil {
		return nil, err
	}
	for _, device := range devices {
		if device.DeviceMAC == mac {
			return device, nil
		}
	}
	return nil, fmt.Errorf("could not find device with mac '%s'", mac)
}
