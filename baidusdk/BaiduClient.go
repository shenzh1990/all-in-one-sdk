package baidusdk

import (
	"errors"
	"github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
)

type BaiduMapClient struct {
	BaseUrl string
	SK      string
}

func NewBaiduMapClient(BaseUrl, SK string) *BaiduMapClient {
	bmc := &BaiduMapClient{
		BaseUrl: BaseUrl,
		SK:      SK,
	}
	return bmc
}
func (d *BaiduMapClient) GetGeoCode(address string) (string, error) {

	Url := d.BaseUrl + "/geocoding/v3/?address=" + address + "&output=json&ak=" + d.SK
	_, body, errs := gorequest.New().Get(Url).EndBytes()
	if len(errs) > 0 {
		return "", errs[0]
	}
	if jsoniter.Get(body, "status").ToInt32() != 0 {
		return "", errors.New(string(body))
	}
	return string(body), nil
}
