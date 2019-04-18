# WeChat SDK for Go

使用Golang开发的微信SDK，简单、易用。

## 快速开始

以下是一个处理客服消息发送的例子：

```
// 指定mp后台配置的Token(令牌)
custom := easywechat.New().GetCustomer("token")

// 回复text 传入access_tokn、openid、content
err := custom.SendText("access_token", "openid", "content")

// 回复图片 传入access_tokn、openid、本地图片路径
err := custom.SendImage("access_token", "openid", "./data/1.jpg")
```

## 内容安全

```
// 传入access_token
wx := easywechat.New().GetChecked("access_token")

// 传入需要检验的内容
err := wx.MsgChecked("content", ...)
```
