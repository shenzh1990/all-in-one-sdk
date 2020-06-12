package wxsdk

import (
	"encoding/xml"
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

/*
 微信接收消息
*/
type WxResMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Url          string
	PicUrl       string
	MediaId      string
	ThumbMediaId string
	Format       string
	Content      string
	MsgId        int
	Scale        int
	Location_X   string
	Location_Y   string
	Label        string
	Title        string
	Description  string
}

/*
 微信回复消息
*/
type WxRepMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Content      string
}
