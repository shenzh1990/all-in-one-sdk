package qywxsdk

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

type QyWxMcClient struct {
	Token          string
	EncodingAESKey string
}

var WCMclient *QyWxMcClient

func NewQyWxMcClient(token, encodingAESKey string) *QyWxMcClient {
	wc := &QyWxMcClient{
		Token:          token,
		EncodingAESKey: encodingAESKey,
	}
	return wc
}

/**
按照微信接入规则生成signature
*/
func (d *QyWxMcClient) GetSignature(timestamp, nonce, msg string) string {
	return MakeSignature(d.Token, timestamp, nonce, msg)
}

/**
解密消息
*/
func (d *QyWxMcClient) DecryptMsg(receiverId, echostr string) (string, error) {
	p_msg, err := Decrypt(receiverId, d.EncodingAESKey, echostr)
	if err != nil {
		return "", err
	}
	return string(p_msg), nil
}

/**
加密消息
*/
func (d *QyWxMcClient) EncryptMsg(receiverId, msg string) (string, error) {
	echostr, err := Encrypt(receiverId, d.EncodingAESKey, msg)
	if err != nil {
		return "", err
	}
	return echostr, nil
}

type Callback func(message *QyWxResMessage) (interface{}, error)

/*
 微信接收消息处理
*/
func (d *QyWxMcClient) QyWxMachineMessage(req *http.Request, callback Callback) (string, error) {
	signature := req.URL.Query().Get("msg_signature")
	timestamp := req.URL.Query().Get("timestamp")
	nonce := req.URL.Query().Get("nonce")
	//把  body 内容读入字符串
	s, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", errors.New("request body read error! detail:" + err.Error())
	}
	wrm := QyWxResMessage{}

	err = xml.Unmarshal(s, &wrm)
	if err != nil {
		return "", errors.New("request body can not unmarshal to object ! detail:" + err.Error())
	}

	if signature == "" || timestamp == "" || nonce == "" || wrm.Encrypt == "" {
		return "", errors.New("Invalid request")
	}
	check_msg_sign := d.GetSignature(timestamp, nonce, wrm.Encrypt)
	if signature != check_msg_sign {
		return "", errors.New("Invalid sign")
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
