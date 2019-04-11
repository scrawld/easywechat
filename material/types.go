package material

import (
	"github.com/scrawld/easywechat/context"
)

// media
type (
	// res
	MediaRes struct {
		context.WxError
		Type      string `json:"type"`
		MediaId   string `json:"media_id"`
		CreatedAt int    `json:"created_at"`
	}
)
