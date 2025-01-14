package qywxsdk

import (
	"encoding/xml"
	"strconv"
	"time"
)

const (
	MsgType_Text       = "text"
	MsgType_Iamge      = "image"
	MsgType_Voice      = "voice"
	MsgType_Video      = "video"
	MsgType_Shortvideo = "shortvideo"
	MsgType_Location   = "location"
	MsgType_Link       = "link"
)

type CDATA struct {
	Value string `xml:",cdata"`
}

/*
微信接收消息
*/
type QyWxResMessage struct {
	Tousername string `xml:"ToUserName"`
	Encrypt    string `xml:"Encrypt"`
	Agentid    string `xml:"AgentID"`
}

/*
微信回复消息
*/
type QWxRepMessage struct {
	XMLName   xml.Name `xml:"xml"`
	Encrypt   CDATA    `xml:"Encrypt"`
	Signature CDATA    `xml:"MsgSignature"`
	Timestamp string   `xml:"TimeStamp"`
	Nonce     CDATA    `xml:"Nonce"`
}

func NewQWxRepMessage(encrypt, signature, timestamp, nonce string) *QWxRepMessage {
	return &QWxRepMessage{Encrypt: CDATA{Value: encrypt}, Signature: CDATA{Value: signature}, Timestamp: timestamp, Nonce: CDATA{Value: nonce}}
}

// 消息的XML包
// <xml>
//
//	<ToUserName>
//	    <![CDATA[ww9bceb0901187d938]]>
//	</ToUserName>
//	<FromUserName>
//	    <![CDATA[ShenZeHua]]>
//	</FromUserName>
//	<CreateTime>1679509951</CreateTime>
//	<MsgType>
//	    <![CDATA[text]]>
//	</MsgType>
//	<Content>
//	    <![CDATA[你好]]>
//	</Content>
//	<MsgId>7213440312870231572</MsgId>
//	<AgentID>1000003</AgentID>
//
// </xml>
type QyWxMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`
	FromUserName CDATA    `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      CDATA    `xml:"MsgType"`
	Content      CDATA    `xml:"Content"`
	MsgId        string   `xml:"MsgId"`
	AgentID      string   `xml:"AgentID"`
}

// 扩展原有消息库 同时使用企业微信机器人和企业微信客服消息
type QyWxKfMessage struct {
	QyWxMessage
	Event    CDATA `xml:"event"`
	Token    CDATA `xml:"Token"`
	OpenKfId CDATA `xml:"OpenKfId"`
}

/*
*
设置消息内容 content内容，msgid 序号 MsgType 消息类型
*/
func NewQyWxMessage(ToUserName, FromUserName, MsgType, content, MsgId, AgentID string) *QyWxMessage {
	return &QyWxMessage{ToUserName: CDATA{Value: ToUserName}, FromUserName: CDATA{Value: FromUserName}, CreateTime: strconv.FormatInt(time.Now().Unix(), 10), MsgType: CDATA{Value: MsgType}, Content: CDATA{Value: content}, MsgId: MsgId, AgentID: AgentID}
}
