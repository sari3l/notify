package telegram

import (
	"fmt"
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
)

// 文档 https://core.telegram.org/bots/api

type Option struct {
	types.BaseOption `yaml:",inline"`
	Webhook          string `yaml:"webhook"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	ChatId string `yaml:"chatId" dict:"chat_id"`
	Text   string `yaml:"text,omitempty" dict:"text"`
}

type notifier struct {
	*Option
}

func (opt *Option) ToNotifier() *notifier {
	noticer := &notifier{}
	noticer.Option = opt
	return noticer
}

func (n *notifier) format(messages []string) (string, ext.Ext) {
	formatMap := utils.GenerateMap(n.NotifyFormatter, messages)
	utils.FormatAnyWithMap(&n.MessageParams, formatMap)
	params := utils.StructToDict(n.MessageParams)
	return n.Webhook, ext.Params(params)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Get(n.format(messages))
	if resp.Ok && resp.Json().Get("ok").Bool() == true {
		return nil
	}
	return fmt.Errorf("[Telegram] [%v] %s", resp.StatusCode, resp.Content)
}
