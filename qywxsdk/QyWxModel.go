package qywxsdk

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
