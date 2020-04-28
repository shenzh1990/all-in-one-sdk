package dingding

//text、markdown、actionCard、feedCard
const (
	Message_Text       = "text"
	Message_MarkDown   = "markdown"
	Message_ActionCard = "actionCard"
	Message_FeedCard   = "feedCard"
)

type McInMessage struct {
	Msgtype           string   `json:"msgtype"`
	Text              Text     `json:"text"`
	MsgID             string   `json:"msgId"`
	CreateAt          int64    `json:"createAt"`
	ConversationType  string   `json:"conversationType"`
	ConversationID    string   `json:"conversationId"`
	ConversationTitle string   `json:"conversationTitle"`
	SenderID          string   `json:"senderId"`
	SenderNick        string   `json:"senderNick"`
	SenderCorpID      string   `json:"senderCorpId"`
	SenderStaffID     string   `json:"senderStaffId"`
	ChatbotUserID     string   `json:"chatbotUserId"`
	AtUsers           []AtUser `json:"atUsers"`
}
type AtUser struct {
	DingtalkID string `json:"dingtalkId"`
	StaffID    string `json:"staffId"`
}

type TestMessage struct {
	MsgType string `json:"msgtype"`
	Text    Text   `json:"text"`
	At      At     `json:"at"`
}
type Text struct {
	Content string `json:"content"`
}

type MarkDownMessage struct {
	MsgType  string   `json:"msgtype"`
	Markdown Markdown `json:"markdown"`
	At       At       `json:"at"`
}
type Markdown struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
type At struct {
	AtMobiles     []string `json:"atMobiles"`
	AtDingtalkIds []string `json:"atDingtalkIds"`
	IsAtAll       bool     `json:"isAtAll"`
}

type ActionCardMessage struct {
	MsgType    string     `json:"msgtype"`
	ActionCard ActionCard `json:"actionCard"`
}
type ActionCard struct {
	Title       string `json:"title"`
	Text        string `json:"text"`
	SingleTitle string `json:"singleTitle"`
	SingleURL   string `json:"singleURL"`
}

type FeedCardMessage struct {
	MsgType  string   `json:"msgtype"`
	FeedCard FeedCard `json:"feedCard"`
}
type FeedCard struct {
	Links []Link `json:"links"`
}
type Link struct {
	Title      string `json:"title"`
	MessageURL string `json:"messageURL"`
	PicURL     string `json:"picURL"`
}
