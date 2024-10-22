package feishu

import (
	"fmt"
	"testing"
)

func TestNotifier_Card(t *testing.T) {
	option := Option{
		Webhook: "https://open.larksuite.com/open-apis/bot/v2/hook/xxxxx",
	}
	option.MsgType = "interactive"
	option.NotifyFormatter = []string{"{{subject}}", "{{rule}}", "{{timestamp}}"}
	option.Card = map[string]interface{}{
		"config": map[string]interface{}{
			"wide_screen_mode": true,
		},
		"elements": []interface{}{
			map[string]interface{}{
				"tag":     "markdown",
				"content": "邮件通知！\n邮件主题：<font color=\"blue\">{{subject}}</font>\n匹配规则：<font color=\"red\">{{rule}}</font>\n匹配时间：{{timestamp}}\n",
			},
		},
		"header": map[string]interface{}{
			"template": "red",
			"title": map[string]interface{}{
				"content": "邮件通知",
				"tag":     "plain_text",
			},
		},
	}
	notifier := option.ToNotifier()
	var err error
	err = notifier.Send([]string{"test subject", "test rule", "test timestamp"})
	err = notifier.Send([]string{"test subject2", "test rule2", "test timestamp2"})
	fmt.Println(err)
}
