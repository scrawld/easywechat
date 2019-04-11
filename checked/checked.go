package checked

import (
	"encoding/json"
	"fmt"
	"github.com/scrawld/easywechat/context"
	"github.com/scrawld/easywechat/utils"
	"strings"
)

const (
	msgCheckUrl = "https://api.weixin.qq.com/wxa/msg_sec_check"
)

type Checked struct {
	AccessToken string
}

func New(token string) *Checked {
	o := &Checked{
		AccessToken: token}
	return o
}

// MsgChecked
func (this *Checked) MsgChecked(contents ...string) (err error) {
	data := map[string]string{"content": strings.Join(contents, "")}
	uri := fmt.Sprintf("%s?access_token=%s", msgCheckUrl, this.AccessToken)
	b, err := utils.PostJson(uri, data)
	if err != nil {
		err = fmt.Errorf("MsgChecked request error: %s", err)
		return
	}
	var r = &context.WxError{}
	err = json.Unmarshal(b, r)
	if err != nil {
		err = fmt.Errorf("MsgChecked json unmarshal error: %s", err)
		return
	}

	if r.Errcode != 0 {
		err = fmt.Errorf("MsgCheck error: code=%d , msg=%s", r.Errcode, r.Errmsg)
		return
	}

	return nil
}
