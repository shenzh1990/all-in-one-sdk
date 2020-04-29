package dingding

import (
	"encoding/base64"
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

type DingMcClient struct {
	appKey    string
	appSecret string
}

func NewDingMcClient(AppKey, AppSecret string) *DingMcClient {
	dc := &DingMcClient{
		appKey:    AppKey,
		appSecret: AppSecret,
	}
	return dc
}

type Callback func(message *McInMessage) (interface{}, error)

func (d *DingMcClient) DingMachineMessage(req *http.Request, callback Callback) (string, error) {

	timestamp := req.Header.Get("timestamp")
	signkey := GetHmacByte(timestamp+"\n"+d.appSecret, d.appSecret)
	signData := base64.StdEncoding.EncodeToString([]byte(signkey))
	sign := req.Header.Get("sign")
	if sign != signData {
		return "", errors.New("Sign do not passed!")
	}
	mcInMessage := McInMessage{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", errors.New("request body read error! detail:" + err.Error())
	}
	err = jsoniter.Unmarshal(body, &mcInMessage)
	if err != nil {
		return "", errors.New("request body is not McInMessage! detail:" + err.Error())
	}
	obj, err := callback(&mcInMessage)
	if err != nil {
		return "", errors.New("callback is err! detail:" + err.Error())
	}

	res, err := jsoniter.MarshalToString(&obj)
	if err != nil {
		return "", errors.New("json marshal error  ! detail:" + err.Error())
	}
	return res, nil

}
