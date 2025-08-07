package vsax

import (
	"fmt"
	"resty.dev/v3"
)

type Client struct {
	server string
	c      *resty.Client
}

func NewClient(server string, key string, token string) Client {
	c := resty.New()
	c.SetBasicAuth(key, token)
	return Client{server, c}
}

func (vc *Client) GetAllDevices() (AllDevicesResult, error) {
	alldevs := AllDevicesResult{}
	_, err := vc.c.R().
		SetResult(&alldevs).
		Get(vc.server + "/api/v3/devices")
	if err != nil {
		fmt.Println(err)
		return AllDevicesResult{}, err
	}
	return alldevs, nil
}

func (vc *Client) GetDevice(deviceId string) (DeviceResult, error) {
	device := DeviceResult{}
	_, err := vc.c.R().
		SetResult(&device).
		Get(vc.server + "/api/v3/devices/" + deviceId)
	if err != nil {
		fmt.Println(err)
		return DeviceResult{}, err
	}
	return device, nil
}
