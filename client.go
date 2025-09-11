package vsax

import (
	"fmt"
	"resty.dev/v3"
	"strconv"
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
func (vc *Client) GetAllAssets(include string) (AllAssetsResult, error) {
	assets := []Asset{}
	results := AllAssetsResult{}
	total := 0

	top := 100
	skip := 0
	_, err := vc.c.R().
		SetResult(&results).
		Get(vc.server + "/api/v3/assets?$top=" + strconv.Itoa(top) +
			"&$skip=" + strconv.Itoa(skip) +
			"&include=" + include)
	if err != nil {
		fmt.Println(err)
		return AllAssetsResult{}, err
	}

	total = results.Meta.TotalCount
	assets = append(assets, results.Data...)
	iterations := total / top
	if total%top > 0 {
		iterations += 1
	}

	for i := 1; i < iterations; i++ {
		skip = top * i
		_, err := vc.c.R().
			SetResult(&results).
			Get(vc.server + "/api/v3/assets?$top=" + strconv.Itoa(top) +
				"&$skip=" + strconv.Itoa(skip) +
				"&include=" + include)
		if err != nil {
			fmt.Println(err)
			return AllAssetsResult{}, err
		}
		assets = append(assets, results.Data...)
	}

	return AllAssetsResult{assets, results.Meta}, nil
}

func (vc *Client) GetAsset(assetId string, include string) (AssetResult, error) {
	asset := AssetResult{}
	_, err := vc.c.R().
		SetResult(&asset).
		Get(vc.server + "/api/v3/assets/" + assetId + "?include=" + include)
	if err != nil {
		fmt.Println(err)
		return AssetResult{}, err
	}
	return asset, nil
}

func (vc *Client) GetAllOrganizations() (AllOrganizationsResult, error) {
	organizations := []Organization{}
	result := AllOrganizationsResult{}
	total := 0

	top := 100
	skip := 0
	_, err := vc.c.R().
		SetResult(&result).
		Get(vc.server + "/api/v3/organizations?$top=" + strconv.Itoa(top) +
			"&$skip=" + strconv.Itoa(skip))
	if err != nil {
		fmt.Println(err)
		return AllOrganizationsResult{}, err
	}

	total = result.Meta.TotalCount
	organizations = append(organizations, result.Data...)

	if total < top {
		return AllOrganizationsResult{organizations, result.Meta}, nil
	}

	iterations := total / top
	if total%top > 0 {
		iterations += 1
	}
	for i := 1; i < iterations; i++ {
		skip = top * i
		_, err := vc.c.R().
			SetResult(&result).
			Get(vc.server + "/api/v3/organizations?$top=" + strconv.Itoa(top) +
				"&$skip=" + strconv.Itoa(skip))
		if err != nil {
			fmt.Println(err)
			return AllOrganizationsResult{}, err
		}
		organizations = append(organizations, result.Data...)
	}

	return AllOrganizationsResult{organizations, result.Meta}, nil
}

func (vc *Client) GetOrganization(organizationId string) (OrganizationResult, error) {
	organization := OrganizationResult{}
	_, err := vc.c.R().
		SetResult(&organization).
		Get(vc.server + "/api/v3/organizations/" + organizationId)
	if err != nil {
		fmt.Println(err)
		return OrganizationResult{}, err
	}
	return organization, nil
}
