package material

import (
	"encoding/json"
	"fmt"
	"github.com/scrawld/easywechat/utils"
)

type MediaType string

const (
	MediaTypeImage MediaType = "image"
	MediaTypeVoice           = "voice"
	MediaTypeVideo           = "video"
	MediaTypeThumb           = "thumb"
)

const (
	mediaUploadUrl = "https://api.weixin.qq.com/cgi-bin/media/upload"
)

// MediaUpload
func (this *Material) MediaUpload(file string, mediaType MediaType) (r *MediaRes, err error) {
	uri := fmt.Sprintf("%s?access_token=%s&type=%s", mediaUploadUrl, this.AccessToken, mediaType)
	b, err := utils.PostFile(uri, "media", file)
	if err != nil {
		return
	}
	r = &MediaRes{}
	err = json.Unmarshal(b, r)
	if err != nil {
		return
	}
	if r.Errcode != 0 {
		err = fmt.Errorf("MediaUpload error: code=%d , msg=%s", r.Errcode, r.Errmsg)
		return
	}
	return
}
