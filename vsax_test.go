package vsax

import "testing"

func TestGetAllDevices(t *testing.T) {
	vc := NewClient("https://stabilityit.vsax.net",
		"541b4c2529ba4605931cbea99cfb0a44",
		"60851f18d19c4d77a6a81ac2fa7d5046863248c629ad45dc80669bc71710cb09")
	ad, err := vc.GetAllDevices()
	if err != nil {
		t.FailNow()
	}
	if ad.Meta.ResponseCode != 200 {
		t.FailNow()
	}
}
