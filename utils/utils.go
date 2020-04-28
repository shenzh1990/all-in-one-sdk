package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetHmacCode(value string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(value))
	return hex.EncodeToString(h.Sum(nil))
}

func GetIp() (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "https://ip.cn", nil)

	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	htmlstr := string(data)
	start := strings.Index(htmlstr, "<code>") + 6
	end := strings.Index(htmlstr, "</code>")
	ip := string([]byte(htmlstr)[start:end])
	return ip, nil
}
