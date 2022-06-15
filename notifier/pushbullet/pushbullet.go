package pushbullet

import (
	"fmt"
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
)

// 文档 https://docs.pushbullet.com/

const DefaultWebhook = "https://api.pushbullet.com/v2/pushes"

type Option struct {
	types.BaseOption `yaml:",inline"`
	Token            string `yaml:"token"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	Type  string `yaml:"type" json:"type"`
	Title string `yaml:"title,omitempty" json:"title,omitempty"`
	Body  string `yaml:"body,omitempty" json:"body,omitempty"`
}

type notifier struct {
	*Option
}

func (opt *Option) ToNotifier() *notifier {
	noticer := &notifier{}
	noticer.Option = opt
	return noticer
}

func (n *notifier) format(messages []string) (string, ext.Ext, ext.Ext) {
	formatMap := utils.GenerateMap(n.NotifyFormatter, messages)
	utils.FormatAnyWithMap(&n.MessageParams, formatMap)
	headers := ext.Dict{"Access-Token": n.Token}
	json := utils.StructToJson(n.MessageParams)
	return DefaultWebhook, ext.Headers(headers), ext.Json(json)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp.Ok {
		return nil
	}
	return fmt.Errorf("[PushBullet] [%v] %s", resp.StatusCode, resp.Content)
}
