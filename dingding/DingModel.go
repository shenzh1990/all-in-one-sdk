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
	Text    Text   `json:"text"`
}

type WorkMessage struct {
	AgentId    int         `json:"agent_id"`
	UserIdList string      `json:"userid_list,omitempty"`
	DeptIdList string      `json:"dept_id_list,omitempty"`
	ToAllUser  bool        `json:"to_all_user,omitempty"y`
	Msg        interface{} `json:"msg"` //不同消息类型使用不同格式，详细查看钉钉
}
type MsgText struct {
	MsgType string `json:"msgtype"`
	Text    Text   `json:"text"`
}
type MsgMarkDown struct {
	MsgType  string   `json:"msgtype"`
	MarkDown Markdown `json:"markdown"`
}
