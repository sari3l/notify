package zulip

import (
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
	rTypes "github.com/sari3l/requests/types"
)

type Option struct {
	types.BaseOption `yaml:",inline"`
	Webhook          string `yaml:"webhook"`
	BotEmail         string `yaml:"botEmail"`
	BotKey           string `yaml:"botKey"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	Type    string  `yaml:"type" dict:"type"`
	To      string  `yaml:"to" dict:"to"`
	Content string  `yaml:"content" dict:"content"`
	Topic   *string `yaml:"topic,omitempty" dict:"topic,omitempty"`
	QueueId *string `yaml:"queueId,omitempty" dict:"queue_id,omitempty"`
	LocalId *string `yaml:"localId,omitempty" dict:"local_id,omitempty"`
}

type Notifier struct {
	*Option
}

func (opt *Option) ToNotifier() *Notifier {
	noticer := &Notifier{}
	noticer.Option = opt
	return noticer
}

func (n *Notifier) format(messages []string) (string, rTypes.Ext, rTypes.Ext) {
	formatMap := utils.GenerateMap(n.NotifyFormatter, messages)
	utils.FormatAnyWithMap(&n.MessageParams, &formatMap)
	data := utils.StructToDict(n.MessageParams)
	auth := rTypes.BasicAuth{Username: n.BotEmail, Password: n.BotKey}
	return n.Webhook, ext.Auth(auth), ext.Form(data)
}

func (n *Notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	return utils.RespCheck("ZuLip", resp, func(request *requests.Response) bool {
		return resp.Ok
	})
}
