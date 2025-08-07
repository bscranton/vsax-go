package vsax

import (
	"fmt"
	"resty.dev/v3"
)

type Client struct {
	Key    string
	Token  string
	server string
	c      *resty.Client
}

func NewClient(server string, key string, token string) Client {
	c := resty.New()
	c.SetBasicAuth(key, token)
	return Client{key, token, server, c}
}

func (vc *Client) GetAllDevices() (AllDevicesResult, error) {
	alldevs := AllDevicesResult{}
	_, err := vc.c.R().
		SetResult(&alldevs).
		Get(vc.server + "/api/v3/devices")
	if err != nil {
		fmt.Println(err)
		return nil, err
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
		return nil, err
	}
	return device, nil
}
	