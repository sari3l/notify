package discord

import (
	"fmt"
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
)

// 文档 https://discord.com/developers/docs/resources/webhook#execute-webhook

type Option struct {
	types.BaseOption `yaml:",inline"`
	Webhook          string `yaml:"webhook"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	Content   string `yaml:"content,omitempty" json:"content,omitempty"`
	Username  string `yaml:"username,omitempty" json:"username,omitempty"`
	AvatarUrl string `yaml:"avatarUrl,omitempty" json:"avatarUrl,omitempty"`
	Tts       bool   `yaml:"tts,omitempty" json:"tts,omitempty"`
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
	data := utils.StructToJson(n.MessageParams)
	return n.Webhook, ext.Json(data)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp.StatusCode == 204 {
		return nil
	}
	return fmt.Errorf("[Discord] [%v] %s", resp.StatusCode, resp.Content)
}
