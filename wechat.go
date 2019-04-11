package easywechat

import (
	"github.com/scrawld/easywechat/checked"
	"github.com/scrawld/easywechat/customer"
	"github.com/scrawld/easywechat/material"
)

type Wechat struct {
}

func New() *Wechat {
	o := &Wechat{}
	return o
}

// GetChecked 内容安全
func (this *Wechat) GetChecked(token string) *checked.Checked {
	return checked.New(token)
}

// GetCustomer 客服消息
func (this *Wechat) GetCustomer(contactToken string) *customer.Customer {
	return customer.New(contactToken)
}

// GetMaterial 素材管理
func (this *Wechat) GetMaterial(token string) *material.Material {
	return material.New(token)
}
