package dingding

import "time"

type AccessTokenRtn struct {
	Errcode     int    `json:"errcode"`
	AccessToken string `json:"access_token"`
	Errmsg      string `json:"errmsg"`
	ExpiresIn   int    `json:"expires_in"`
	GetTime     time.Time
}

type GroupMessage struct {
	Chatid  string `json:"chatid"`
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}
