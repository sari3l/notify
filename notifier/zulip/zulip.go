package zulip

import (
	"fmt"
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
)

// 文档 https://zulip.com/api/send-message

type Option struct {
	types.BaseOption `yaml:",inline"`
	Webhook          string `yaml:"webhook"`
	BotEmail         string `yaml:"botEmail"`
	BotKey           string `yaml:"botKey"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	Type    string `yaml:"type" dict:"type"`
	To      string `yaml:"to" dict:"to"`
	Content string `yaml:"content" dict:"content"`
	Topic   string `yaml:"topic,omitempty" dict:"topic,omitempty"`
	QueueId string `yaml:"queueId,omitempty" dict:"queue_id,omitempty"`
	LocalId string `yaml:"localId,omitempty" dict:"local_id,omitempty"`
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
	utils.FormatAnyWithMap(&n.MessageParams, &formatMap)
	data := utils.StructToDict(n.MessageParams)
	auth := ext.BasicAuth{Username: n.BotEmail, Password: n.BotKey}
	return n.Webhook, ext.Auth(auth), ext.Data(data)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp != nil && resp.Ok {
		return nil
	}
	return fmt.Errorf("[ZuLip] [%v] %s", resp.StatusCode, resp.Content)
}
