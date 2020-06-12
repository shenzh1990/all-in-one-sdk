package wxsdk

import (
	"encoding/xml"
	"time"
)

type AccessTokenRtn struct {
	Errcode     int    `json:"errcode"`
	AccessToken string `json:"access_token"`
	Errmsg      string `json:"errmsg"`
	ExpiresIn   int    `json:"expires_in"`
	GetTime     time.Time
}

type WxMenu struct {
	Button []Button `json:"button"`
}
type Button struct {
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Key       string   `json:"key,omitempty"`
	Url       string   `json:"url,omitempty"`
	MediaID   string   `json:"media_id,omitempty"`
	AppID     string   `json:"appid,omitempty"`
	PagePath  string   `json:"pagepath,omitempty"`
	SubButton []Button `json:"sub_button,omitempty"`
}

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
