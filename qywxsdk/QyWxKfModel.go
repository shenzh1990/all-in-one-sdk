package qywxsdk

type SyncMsg struct {
	Cursor      string `json:"cursor"`
	Token       string `json:"token"`
	Limit       int    `json:"limit"`
	VoiceFormat int    `json:"voice_format"`
	OpenKfid    string `json:"open_kfid"`
}
type QyWxKfMessage struct {
	QyWxMessage
	Event    CDATA `xml:"event"`
	Token    CDATA `xml:"Token"`
	OpenKfId CDATA `xml:"OpenKfId"`
}
