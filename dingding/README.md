# 钉钉机器人接口实现

## 核心函数解析
```go
    func (d *DingMcClient) DingMachineMessage(req *http.Request, callback Callback) (string, error)
```
###入参
-  req ：为 httpRequest 请求，
- callback：回调函数，用来接受钉钉消息内容，并提供返回

###出参
- string: 返回需要回复json 内容
- error：错误信息

```go
type Callback func(message *McInMessage) (interface{}, error)
```
**该函数需要自行实现**
###入参
-  meesage： 钉钉机器人接受到的消息，

###出参
- interface: 返回需要回复实体对象
- error：错误信息

调用示例：
```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenzh1990/all-in-one-sdk/dingding"
	"net/http"
	"strings"
)
var dmcc =dingding.NewDingMcClient("","")
func main(){
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/readmes",getMsg)
	r.Run(":8080")
}
func getMsg(c *gin.Context){
	res,err:=dmcc.DingMachineMessage(c.Request,handler)
	if err!=nil{
		c.String(http.StatusOK,err.Error())
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	c.String(http.StatusOK,res)
}

func handler(message *dingding.McInMessage)( interface{},error){
   if strings.Contains(message.Text.Content,"mark") {
		tmsg :=dingding.MarkDownMessage{
			MsgType: dingding.Message_MarkDown,
			Markdown: struct {
				Title string `json:"title"`
				Text  string `json:"text"`
			}{
				Title: "markDown测试11111",
				Text:  "#### 杭州天气 @150XXXXXXXX \n> 9度，西北风1级，空气良89，相对温度73%\n> ![screenshot](https://img.alicdn.com/tfs/TB1NwmBEL9TBuNjy1zbXXXpepXa-2400-1218.png)\n> ###### 10点20分发布 [天气](https://www.dingalk.com) \n",
			},
		}
		return tmsg,nil
	}else if strings.Contains( message.Text.Content,"action") {
		tmsg :=dingding.ActionCardMessage{
			MsgType: dingding.Message_ActionCard,
			ActionCard: struct {
				Title       string `json:"title"`
				Text        string `json:"text"`
				SingleTitle string `json:"singleTitle"`
				SingleURL   string `json:"singleURL"`
			}{
				Title:       "测试 actioncard",
				Text:        " 你哦好百度",
				SingleTitle: "百度",
				SingleURL:   "https://baidu.com",
			},
		}

		return tmsg,nil

	}else if strings.Contains(message.Text.Content, "feed") {

		links:=make([]dingding.Link,0)
       link:=dingding.Link{
			Title:      "测试一",
			MessageURL: "https://www.baidu.com",
			PicURL:     "http://qinniu.qinyule.com/image/blur-business-coffee-commerce-273222.jpg",
		}
		links=append(links, link)
		link=dingding.Link{
			Title:      "测试二",
			MessageURL: "https://www.qinyule.com",
			PicURL:     "http://qinniu.qinyule.com/image/blur-business-coffee-commerce-273222.jpg",
		}
		links=append(links, link)
		tmsg :=dingding.FeedCardMessage{
			MsgType:  dingding.Message_FeedCard,
			FeedCard: dingding.FeedCard{
				Links: links,
			},
		}
		return tmsg,nil
	}else {
		tmsg :=dingding.TestMessage{
			MsgType: dingding.Message_Text,
			Text: struct {
				Content string `json:"content"`
			}{
				Content: "text 测试,你说："+message.Text.Content,
			},
			At: struct {
				AtMobiles     []string `json:"atMobiles"`
				AtDingtalkIds []string `json:"atDingtalkIds"`
				IsAtAll       bool     `json:"isAtAll"`
			}{},
		}
		return tmsg,nil
	}
}

```
