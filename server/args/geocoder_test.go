package args

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/techidea8/restgo/utils"
)

func TestGeocodeResp(t *testing.T) {
	var resp GeocoderResp
	url := "https://apis.map.qq.com/ws/geocoder/v1/?location=39.984154,116.307490&key=X4PBZ-XHD62-7FVUT-CNELI-QFUDV-CFBOO"
	bts, err := utils.HttpGet(url)
	if err != nil {
		fmt.Sprintf(err.Error())
	} else {
		json.Unmarshal(bts, &resp)
		fmt.Printf("%v", resp)
	}
}

func TestIpcodeResp(t *testing.T) {
	var resp IpcoderResp
	url := "https://apis.map.qq.com/ws/location/v1/ip?ip=175.13.99.76&key=X4PBZ-XHD62-7FVUT-CNELI-QFUDV-CFBOO"
	bts, err := utils.HttpGet(url)
	if err != nil {
		fmt.Sprintf(err.Error())
	} else {
		json.Unmarshal(bts, &resp)
		fmt.Printf("%v", resp)
	}
}
