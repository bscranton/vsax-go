package vsax

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"testing"
)

func TestGetAllDevices(t *testing.T) {
	vc := NewClient("https://stabilityit.vsax.net",
		os.Getenv("VSAX_CLIENT_KEY"),
		os.Getenv("VSAX_CLIENT_TOKEN"))
	ad, err := vc.GetAllDevices()
	if err != nil {
		t.FailNow()
	}
	if ad.Meta.ResponseCode != 200 {
		t.FailNow()
	}
}

func TestGetDevice(t *testing.T) {
	vc := NewClient("https://stabilityit.vsax.net",
		os.Getenv("VSAX_CLIENT_KEY"),
		os.Getenv("VSAX_CLIENT_TOKEN"))
	device, err := vc.GetDevice("1e6aaa00-0d65-4767-a24d-71955fc57573")
	if err != nil {
		t.FailNow()
	}
	if device.Meta.ResponseCode != 200 {
		t.FailNow()
	}
}

func TestGetAllAssets(t *testing.T) {
}

func TestGetAsset(t *testing.T) {
}
