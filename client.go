package vsax

import (
	"fmt"
    "strconv"
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

// This is not async and may take a *very* long time to finish.
func (vc *Client) GetAllAssets() (AllAssetsResult, error) {
    assets := []Asset{}
    results := AllAssetsResult{}
    total := 0
    
    top := 100
    skip := 0
    _, err := vc.c.R().
        SetResult(&results).
        Get(vc.server + "/api/v3/assets?$top=" + strconv.Itoa(top) + 
            "&$skip=" + strconv.Itoa(skip) + 
            "&include=none")
    if err != nil {
        fmt.Println(err)
        return AllAssetsResult{}, err
    }
    
    total = results.Meta.TotalCount
    assets = append(assets, results.Data...)
    iterations := total/top
    if total % top > 0 {
        iterations += 1
    }
    
    for i := 1; i < iterations; i++ {
        skip = top * i
        _, err := vc.c.R().
            SetResult(&results).
            Get(vc.server + "/api/v3/assets?$top=" + strconv.Itoa(top) + 
                "&$skip=" + strconv.Itoa(skip) + 
                "&include=none")
        if err != nil {
            fmt.Println(err)
            return AllAssetsResult{}, err
        }
        assets = append(assets, results.Data...)
    }
    
    return AllAssetsResult{assets, results.Meta}, nil
}
    

func (vc *Client) GetAsset(assetId string) (AssetResult,error) {
    asset := AssetResult{}
    _, err := vc.c.R().
        SetResult(&asset).
        Get(vc.server + "/api/v3/assets/" + assetId)
    if err != nil {
        fmt.Println(err)
        return AssetResult{}, err
    }
    return asset, nil
}
