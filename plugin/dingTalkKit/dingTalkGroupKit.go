// Package dingTalkKit 钉钉开发群工具包
package dingTalkKit

import (
	"encoding/json"
	"fmt"
	"github.com/xingcxb/goKit/core/cryptoKit"
	"github.com/xingcxb/goKit/core/httpKit"
	"github.com/xingcxb/goKit/core/strKit"
	"strconv"
	"time"
)

type DingTalkBot struct {
	WebHookUrl string // 钉钉机器人webhook地址
	Secret     string // 钉钉机器人secret
}

// NewDingTalkBot 创建一个钉钉机器人
func NewDingTalkBot(webHookUrl, secret string) *DingTalkBot {
	return &DingTalkBot{WebHookUrl: webHookUrl, Secret: secret}
}

// SendTextMessage 发送文本消息
/*
 * @param text 发送内容
 * @param atIdType 1:@手机, 2:@userId, 3:@所有人, other:不@
 * @param at 传入的手机号或者用户id
 */
func (d *DingTalkBot) SendTextMessage(text string, atIdType int, at ...string) error {
	msg := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]string{
			"content": text,
		},
	}
	return d.core(atIdType, msg, at)
}

// SendLinkMessage 发送link消息
/*
 * @param title 标题
 * @param text 发送内容
 * @param messageUrl 跳转地址
 * @param picUrl 图片地址
 * @param atIdType 1:@手机, 2:@userId, 3:@所有人, other:不@
 * @param at 传入的手机号或者用户id
 */
func (d *DingTalkBot) SendLinkMessage(title, text, messageUrl, picUrl string, atIdType int, at ...string) error {
	msg := map[string]interface{}{
		"msgtype": "link",
		"link": map[string]string{
			"title":      title,
			"text":       text,
			"messageUrl": messageUrl,
			"picUrl":     picUrl,
		},
	}
	return d.core(atIdType, msg, at)
}

// SendMarkdownMessage 发送markdown消息
/*
 * @param title 标题(用于左侧显示消息列表中显示缩略消息，在具体的内容中不显示)
 * @param text 发送内容
 * @param atIdType 1:@手机, 2:@userId, 3:@所有人, other:不@
 * @param at 传入的手机号或者用户id
 */
func (d *DingTalkBot) SendMarkdownMessage(title, text string, atIdType int, at ...string) error {
	msg := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"title": title,
			"text":  text,
		},
	}
	return d.core(atIdType, msg, at)
}

// SendActionCardMessage 发送actionCard消息
/*
 * @param title 标题
 * @param text 发送内容
 * @param btnOrientation 按钮排序方向 0：按钮竖直排列,1：按钮横向排列
 * @param singleTitle 单标题
 * @param singleURL 单链接
 */
func (d *DingTalkBot) SendActionCardMessage(title, text string, btnOrientation int, singleTitle, singleURL string) error {
	msg := map[string]interface{}{
		"msgtype": "actionCard",
		"actionCard": map[string]interface{}{
			"title":          title,
			"text":           text,
			"singleTitle":    singleTitle,
			"singleURL":      singleURL,
			"btnOrientation": btnOrientation,
		},
	}
	return d.core(0, msg, []string{})
}

// core 核心
/*
 * @param atIdType 1:@手机, 2:@userId, 3:@所有人, other:不@
 * @param msg 消息体
 * @param at 传入的手机号或者用户id
 */
func (d *DingTalkBot) core(atIdType int, msg map[string]interface{}, at []string) error {
	switch atIdType {
	case 1:
		// @手机号
		msg["at"] = map[string]interface{}{
			"atMobiles": at,
			"isAtAll":   false,
		}
	case 2:
		// @用户id
		msg["at"] = map[string]interface{}{
			"atUserIds": at,
			"isAtAll":   false,
		}
	case 3:
		// @所有人
		msg["at"] = map[string]interface{}{
			"isAtAll": true,
		}
	}
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	headers := make(map[string]string, 0)
	headers["Content-Type"] = "application/json"
	response, err := httpKit.HttpPostFull(d.getURL(), headers, nil, b, -1)
	if err != nil {
		return err
	} else {
		fmt.Println(response)
		return nil
	}
}

func (d *DingTalkBot) getURL() string {
	timestamp := time.Now().UnixMilli()
	stringToSign := strKit.Splicing(strconv.FormatInt(timestamp, 10), "\n", d.Secret)
	sign := cryptoKit.Hmac256(stringToSign, d.Secret)
	url := strKit.Splicing(d.WebHookUrl, "&timestamp=", strconv.FormatInt(timestamp, 10), "&sign=", sign)
	return url
}
