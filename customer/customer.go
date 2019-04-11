package customer

import (
	"encoding/json"
	"fmt"
	"github.com/scrawld/easywechat/context"
	"github.com/scrawld/easywechat/material"
	"github.com/scrawld/easywechat/utils"
)

const (
	sendUrl = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
)

type Customer struct {
	ContactToken string
}

func New(contactToken string) *Customer {
	o := &Customer{
		ContactToken: contactToken}
	return o
}

// SendText
func (this *Customer) SendText(token, openid, content string) (err error) {
	var data = map[string]interface{}{
		"touser":  openid,
		"msgtype": "text",
		"text":    map[string]string{"content": content}}
	return this.Send(token, data)
}

// SendImage
func (this *Customer) SendImage(token, openid, file string) (err error) {
	media, err := material.New(token).MediaUpload(file, material.MediaTypeImage)
	if err != nil {
		err = fmt.Errorf("Media upload error: %s", err)
		return
	}
	var data = map[string]interface{}{
		"touser":  openid,
		"msgtype": "image",
		"image":   map[string]string{"media_id": media.MediaId}}
	return this.Send(token, data)
}

// Send
func (this *Customer) Send(token string, data map[string]interface{}) (err error) {
	uri := fmt.Sprintf("%s?access_token=%s", sendUrl, token)
	b, err := utils.PostJson(uri, data)
	if err != nil {
		err = fmt.Errorf("Customer request error: %s", err)
		return
	}
	var r = &context.WxError{}
	err = json.Unmarshal(b, r)
	if err != nil {
		err = fmt.Errorf("Customer json unmarshal error: %s", err)
		return
	}
	if r.Errcode != 0 {
		err = fmt.Errorf("Customer send error: code=%d , msg=%s", r.Errcode, r.Errmsg)
		return
	}
	return
}
