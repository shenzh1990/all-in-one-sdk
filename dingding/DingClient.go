package dingding

import (
	"errors"
	"github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
	"sync"
	"time"
)

type DingClient struct {
	DingUrl         string
	AppKey          string
	AppSecret       string
	AccessToken     AccessTokenRtn
	AccessTokenMute sync.RWMutex
}

func NewDingClient(DingUrl, AppKey, AppSecret string) *DingClient {
	dc := &DingClient{
		DingUrl:   DingUrl,
		AppKey:    AppKey,
		AppSecret: AppSecret,
	}
	dc.initAccessToken()
	return dc
}
func (d *DingClient) initAccessToken() error {
	_, _, errs := gorequest.New().Get(d.DingUrl + "/gettoken?appkey=" +
		d.AppKey + "&appsecret=" + d.AppSecret + "").EndStruct(&d.AccessToken)
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
func (d *DingClient) GetAccessToken() string {
	d.AccessTokenMute.Lock()
	defer d.AccessTokenMute.Unlock()
	//获得成功或者获得的时间比现在的时间早，则重新获取
	if d.AccessToken.GetTime.Before(time.Now()) {
		d.initAccessToken()
	}
	return d.AccessToken.AccessToken
}

/**
发送群消息 对应钉钉 =chat/send
*/
func (d *DingClient) SendGroupMessage(message string, chatId string) error {
	groupMessage := GroupMessage{
		Chatid:  chatId,
		Msgtype: "text",
		Text: struct {
			Content string `json:"content"`
		}{message},
	}
	//groupMessageJson, _ := jsoniter.Marshal(groupMessage)
	accesstoken := d.GetAccessToken()

	_, body, errs := gorequest.New().Post(d.DingUrl + "/chat/send?access_token=" +
		accesstoken).Send(groupMessage).EndBytes()
	if len(errs) > 0 {
		return errors.New("access get error")
	}

	if jsoniter.Get(body, "errcode").ToInt32() != 0 {
		return errors.New(string(body))
	}
	return nil
}
