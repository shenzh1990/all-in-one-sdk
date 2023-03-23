# 企业微信机器人接口实现

## 核心函数解析
```go
    func (d *QyWxMcClient) QyWxMachineMessage(req *http.Request, callback Callback)  (string, error) 
```
入参
-  req ：为 httpRequest 请求，
- callback：回调函数，用来接受企业微信消息内容，并提供返回

出参


```go
type Callback func(message *QyWxResMessage) (interface{}, error)
```
**该函数需要自行实现**

入参
-  meesage： 企业微信接受到的消息，

出参
- interface: 返回需要回复实体对象
- error：错误信息

调用示例：
```go
var qywx = qywxsdk.NewQyWxMcClient(token, EncodingAESKey)

func main() {
r := gin.New()
r.Use(gin.Logger())
r.Use(gin.Recovery())
r.GET("/ping", func(c *gin.Context) {
c.JSON(200, gin.H{
"message": "pong",
})
})
//验签回调的url所需的get请求
r.GET("/qychat/msg", handleMsg)
//接收消息和回复所需的post请求
r.POST("/qychat/msg", handlePostMsg)
r.Run(":8080")
}
func handleMsg(c *gin.Context) {
signature := c.Query("msg_signature")
timestamp := c.Query("timestamp")
nonce := c.Query("nonce")
echostr := c.Query("echostr")
if signature == "" || timestamp == "" || nonce == "" || echostr == "" {
c.String(http.StatusBadRequest, "Invalid request")
return
}
//验签过程
check_msg_sign := qywx.GetSignature(timestamp, nonce, echostr)
if signature == check_msg_sign {
msg, err := qywx.DecryptMsg("", echostr)
if err != nil {
c.String(http.StatusBadRequest, err.Error())
} else {
c.String(http.StatusOK, msg)
}
} else {
c.String(http.StatusBadRequest, "Invalid signature")
}
}
func handlePostMsg(c *gin.Context) {
	//消息接收和恢复的回调函数，回调函数的实现在handler函数中
res, err := qywx.QyWxMachineMessage(c.Request, handler)
if err != nil {
c.String(http.StatusOK, err.Error())
return
}
fmt.Println(res)
c.String(http.StatusOK, res)
}
//接收和回复消息的主要接口，请参考实现
func handler(message *qywxsdk.QyWxResMessage) (interface{}, error) {
	//解密消息内容 ，注意实际的消息内容在Encrypt中，此内容解密后还是xml文件
msg, err := qywx.DecryptMsg("", message.Encrypt)
if nil != err {
return "", err
}
//解密后内容转成对象
wrm := qywxsdk.QyWxMessage{}

err = xml.Unmarshal([]byte(msg), &wrm)
if err != nil {
return "", errors.New("QyWxMessage can not unmarshal to object ! detail:" + err.Error())
}
//设置回复所需的内容，注意此对象需要转成xml
reply_obj := qywxsdk.NewQyWxMessage(wrm.FromUserName.Value, wrm.ToUserName.Value, "text", "收到你的消息啦", wrm.MsgId, wrm.AgentID)
reply_msg, err := xml.Marshal(&reply_obj)
if err != nil {
return "", errors.New("reply xml marshal error  ! detail:" + err.Error())
}
//加密回复内容
r_msg, err := qywx.EncryptMsg("", string(reply_msg))
fmt.Println(r_msg)
if nil != err {
return "", err
}
//下列操作为了把回复内容验签，企业微信能够验签和识别你的回复内容
timestamp := strconv.FormatInt(time.Now().Unix(), 10)
nonce := strconv.FormatInt(time.Now().AddDate(1, 1, 1).Unix(), 10)
signature := qywx.GetSignature(timestamp, nonce, r_msg)
repmsg := qywxsdk.NewQWxRepMessage(r_msg, signature, timestamp, nonce)
return repmsg, nil
}

```
