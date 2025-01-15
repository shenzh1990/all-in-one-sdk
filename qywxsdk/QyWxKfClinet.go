package qywxsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type WeWorkClient struct {
	Qywxclient QyWxClient
	Url        string
}

func NewWeWorkClient(Url string) *WeWorkClient {
	we := &WeWorkClient{
		Url: Url,
	}
	return we
}

func (we *WeWorkClient) SyncMsg(syncMsg SyncMsg) (string, error) {
	access_token := we.Qywxclient.GetAccessToken()
	url := we.Url + "/cgi-bin/kf/sync_msg?access_token=" + access_token
	// 将syncMsg编码为JSON格式的字节
	jsonBytes, err := json.Marshal(syncMsg)
	if err != nil {
		return "", fmt.Errorf("Error encoding JSON: %v", err)
	}
	// 使用字节切片创建一个具有读取功能的对象
	reader := bytes.NewReader(jsonBytes)
	session := &http.Client{}

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return "", fmt.Errorf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := session.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body: %v", err)
	}
	return string(body), nil
}
func (we *WeWorkClient) SendMsg(message string) (string, error) {
	access_token := we.Qywxclient.GetAccessToken()
	url := we.Url + "/cgi-bin/kf/send_msg?access_token=" + access_token
	// 将syncMsg编码为JSON格式的字节
	jsonBytes, err := json.Marshal(message)
	if err != nil {
		return "", fmt.Errorf("Error encoding JSON: %v", err)
	}
	// 使用字节切片创建一个具有读取功能的对象
	reader := bytes.NewReader(jsonBytes)
	session := &http.Client{}

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return "", fmt.Errorf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := session.Do(req)
	if err != nil {
		return "", fmt.Errorf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading response body: %v", err)
	}

	return string(body), nil
}
