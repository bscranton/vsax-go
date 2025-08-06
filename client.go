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

func (vc *Client) GetAllDevices() (AllDevices, error) {
	alldevs := AllDevices{}
	_, err := vc.c.R().
		SetResult(&alldevs).
		Get(vc.server + "/api/v3/devices")
	if err != nil {
		fmt.Println(err)
		return AllDevices{}, err
	}
	return alldevs, nil
}
