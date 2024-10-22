package feishu

import (
	"fmt"
	"testing"
)

func TestNotifier_Send(t *testing.T) {
	option := Option{
		Webhook: "https://open.larksuite.com/open-apis/bot/v2/hook/xxxxxxxx",
	}
	option.MsgType = "interactive"
	option.NotifyFormatter = []string{}
	option.Card = map[string]interface{}{
		"elements": []interface{}{
			map[string]interface{}{
				"tag": "div",
				"text": map[string]interface{}{
					"content": "**西湖**，位于浙江省杭州市西湖区龙井路1号，杭州市区西部，景区总面积49平方千米，汇水面积为21.22平方千米，湖面面积为6.38平方千米。",
					"tag":     "lark_md",
				},
			},
			map[string]interface{}{
				"actions": []interface{}{
					map[string]interface{}{
						"tag": "button",
						"text": map[string]interface{}{
							"content": "更多景点介绍 :玫瑰:",
							"tag":     "lark_md",
						},
						"url":   "https://www.example.com",
						"type":  "default",
						"value": map[string]interface{}{},
					},
				},
				"tag": "action",
			},
		},
		"header": map[string]interface{}{
			"title": map[string]interface{}{
				"content": "今日旅游推荐",
				"tag":     "plain_text",
			},
		},
	}
	notifier := option.ToNotifier()
	resp := notifier.Send(nil)
	fmt.Println(resp)
}
