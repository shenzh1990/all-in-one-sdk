package wxsdk

import (
	"crypto/sha1"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
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

/**
按照微信接入规则生成signature
*/
func (d *WxClient) MakeSignature(token, timestamp, nonce string) string {

	//1. 将 plat_token、timestamp、nonce三个参数进行字典序排序
	sl := []string{token, timestamp, nonce}
	sort.Strings(sl)
	//2. 将三个参数字符串拼接成一个字符串进行sha1加密
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))

	return fmt.Sprintf("%x", s.Sum(nil))
}

type Callback func(message *WxResMessage) (interface{}, error)

/*
 微信接收消息处理
*/
func (d *WxClient) WxMachineMessage(req *http.Request, callback Callback) (string, error) {
	//把  body 内容读入字符串
	s, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", errors.New("request body read error! detail:" + err.Error())
	}
	wrm := WxResMessage{}

	err = xml.Unmarshal(s, &wrm)
	if err != nil {
		return "", errors.New("request body can not unmarshal to object ! detail:" + err.Error())
	}
	obj, err := callback(&wrm)
	if err != nil {
		return "", errors.New("callback error ! detail:" + err.Error())
	}
	res, err := xml.Marshal(&obj)
	if err != nil {
		return "", errors.New("xml marshal error  ! detail:" + err.Error())
	}
	return string(res), nil
}

/**
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

/**
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

/**
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

/**
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
