package wxsdk

import (
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
type TemplateMessage struct {
	ToUser     string                 `json:"touser"`
	TemplateID string                 `json:"template_id"`
	URL        string                 `json:"url,omitempty"`
	Data       map[string]ValueObject `json:"data"`
}

type ValueObject struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}
