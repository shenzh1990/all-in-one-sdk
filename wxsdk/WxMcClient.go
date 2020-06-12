package wxsdk

import (
	"crypto/sha1"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

type WxMcClient struct {
	Token          string
	EncodingAESKey string
}

var WCMclient *WxMcClient

func NewWxMcClient(token, encodingAESKey string) *WxMcClient {
	wc := &WxMcClient{
		Token:          token,
		EncodingAESKey: encodingAESKey,
	}
	return wc
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
