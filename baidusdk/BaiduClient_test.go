package baidusdk

import "testing"

func TestBaiduMapClient_GetGeoCode(t *testing.T) {

	d := &BaiduMapClient{
		BaseUrl: "http://api.map.baidu.com",
		SK:      "j0kOCLIvS2mR5pRLSW39Sao6HgNZwFD8",
	}
	got, err := d.GetGeoCode("杭州大厦")
	if err != nil {
		t.Errorf("BaiduMapClient.GetGeoCode() error = %v", err)
		return
	}
	t.Log(got)

}
