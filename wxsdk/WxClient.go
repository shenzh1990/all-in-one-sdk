package wxsdk

import (
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
	"sync"
	"time"
)

type WxClient struct {
	WxUrl           string
	AppKey          string
	AppSecret       string
	AccessToken     AccessTokenRtn
	AccessTokenMute sync.RWMutex
}

var WClient *WxClient

func NewWxClient(WxUrl, AppKey, AppSecret string) *WxClient {
	wc := &WxClient{
		WxUrl:     WxUrl,
		AppKey:    AppKey,
		AppSecret: AppSecret,
	}
	wc.initAccessToken()
	return wc
}
func (d *WxClient) initAccessToken() error {
	_, _, errs := gorequest.New().Get(d.WxUrl + "/cgi-bin/token?grant_type=client_credential&appid=" +
		d.AppKey + "&secret=" + d.AppSecret + "").EndStruct(&d.AccessToken)
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
func (d *WxClient) GetAccessToken() string {
	d.AccessTokenMute.Lock()
	defer d.AccessTokenMute.Unlock()
	//获得成功或者获得的时间比现在的时间早，则重新获取
	if d.AccessToken.GetTime.Before(time.Now()) {
		d.initAccessToken()
	}
	return d.AccessToken.AccessToken
}

/*
*
自定义菜单创建 https://developers.weixin.qq.com/doc/offiaccount/Custom_Menus/Creating_Custom-Defined_Menu.html
*/
func (d *WxClient) CreateMenu(wxMenu WxMenu) error {

	//groupMessageJson, _ := jsoniter.Marshal(groupMessage)
	accesstoken := d.GetAccessToken()

	_, body, errs := gorequest.New().Post(d.WxUrl + "/cgi-bin/menu/create?access_token=" +
		accesstoken).Send(wxMenu).EndBytes()
	if len(errs) > 0 {
		return errors.New("create menu error")
	}

	if jsoniter.Get(body, "errcode").ToInt32() != 0 {
		fmt.Print(string(body))
		return errors.New(string(body))
	}
	return nil
}

/*
*
自定义菜单包括非自定义菜单查询
*/
func (d *WxClient) GetCurrentMenu() (WxMenu, error) {

	wxMenu := WxMenu{}
	//groupMessageJson, _ := jsoniter.Marshal(groupMessage)
	accesstoken := d.GetAccessToken()

	_, _, errs := gorequest.New().Get(d.WxUrl + "/cgi-bin/get_current_selfmenu_info?access_token=" +
		accesstoken).EndStruct(&wxMenu)
	if len(errs) > 0 {
		return wxMenu, errors.New("get_current_selfmenu_info get error")
	}
	return wxMenu, nil
}

/*
*
自定义菜单删除
*/
func (d *WxClient) DeleteMenu() error {

	wxMenu := WxMenu{}
	//groupMessageJson, _ := jsoniter.Marshal(groupMessage)
	accesstoken := d.GetAccessToken()

	_, body, errs := gorequest.New().Get(d.WxUrl + "/cgi-bin/menu/delete?access_token=" +
		accesstoken).EndStruct(&wxMenu)
	if len(errs) > 0 {
		return errors.New("delete menu get error")
	}
	if jsoniter.Get(body, "errcode").ToInt32() != 0 {
		fmt.Print(string(body))
		return errors.New(string(body))
	}
	return nil
}

/*
*
自定义菜单包括非自定义菜单查询
*/
func (d *WxClient) GetMenu() (WxMenu, error) {

	wxMenu := WxMenu{}
	//groupMessageJson, _ := jsoniter.Marshal(groupMessage)
	accesstoken := d.GetAccessToken()

	_, _, errs := gorequest.New().Get(d.WxUrl + "/cgi-bin/menu/get?access_token=" +
		accesstoken).EndStruct(&wxMenu)
	if len(errs) > 0 {
		return wxMenu, errors.New(" menu get error")
	}
	return wxMenu, nil
}

/*
*
发送模板消息
*/
func (d *WxClient) SendTemplateMessage(message TemplateMessage) error {
	accesstoken := d.GetAccessToken()
	_, body, errs := gorequest.New().Post(d.WxUrl + "/cgi-bin/message/template/send?access_token=" +
		accesstoken).Send(message).EndBytes()
	if len(errs) > 0 {
		return errors.New("create menu error")
	}

	if jsoniter.Get(body, "errcode").ToInt32() != 0 {
		fmt.Print(string(body))
		return errors.New(string(body))
	}
	return nil
}
