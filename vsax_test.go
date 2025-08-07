package vsax

import "testing"

func TestGetAllDevices(t *testing.T) {
	vc := NewClient("https://stabilityit.vsax.net",
		"KEY",
		"TOKEN")
	ad, err := vc.GetAllDevices()
	if err != nil {
		t.FailNow()
	}
	if ad.Meta.ResponseCode != 200 {
		t.FailNow()
	}
}
