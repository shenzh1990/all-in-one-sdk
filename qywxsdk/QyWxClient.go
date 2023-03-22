package qywxsdk

import (
	"errors"
	"github.com/parnurzeal/gorequest"
	"sync"
	"time"
)

type QyWxClient struct {
	WxUrl           string
	CropId          string
	CropSecret      string
	AccessToken     AccessTokenRtn
	AccessTokenMute sync.RWMutex
}

var QyWClient *QyWxClient

func NewQyWxClient(WxUrl, CropId, CropSecret string) *QyWxClient {
	wc := &QyWxClient{
		WxUrl:      WxUrl,
		CropId:     CropId,
		CropSecret: CropSecret,
	}
	wc.initAccessToken()
	return wc
}
func (d *QyWxClient) initAccessToken() error {
	_, _, errs := gorequest.New().Get(d.WxUrl + "cgi-bin/gettoken?corpid=" +
		d.CropId + "corpsecret=" + d.CropSecret + "").EndStruct(&d.AccessToken)
	if len(errs) > 0 {
		return errors.New("access get error")
	}
	if d.AccessToken.Errcode != 0 {
		return errors.New(d.AccessToken.Errmsg)
	}
	//设置token过期重新获取时间为5000秒
	d.AccessToken.GetTime = time.Now().Add(5000 * time.Second)
	return nil
}
func (d *QyWxClient) GetAccessToken() string {
	d.AccessTokenMute.Lock()
	defer d.AccessTokenMute.Unlock()
	//获得成功或者获得的时间比现在的时间早，则重新获取
	if d.AccessToken.GetTime.Before(time.Now()) {
		d.initAccessToken()
	}
	return d.AccessToken.AccessToken
}
