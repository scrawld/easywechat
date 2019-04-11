package context

type WxError struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
