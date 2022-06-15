package pushover

import (
	"fmt"
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
)

// 文档 https://pushover.net/

const DefaultWebhook = "https://api.pushover.net/1/messages.json"

type Option struct {
	types.BaseOption `yaml:",inline"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	Token    string `yaml:"token" dict:"token"`
	User     string `yaml:"user" dict:"user"`
	Message  string `yaml:"message" dict:"message"`
	Device   string `yaml:"device,omitempty" dict:"device,omitempty"`
	Html     int    `yaml:"html,omitempty" dict:"html,omitempty"`
	Priority int    `yaml:"priority,omitempty" dict:"priority,omitempty"`
	Sound    string `yaml:"sound,omitempty" dict:"sound,omitempty"`
	Title    string `yaml:"title,omitempty" dict:"title,omitempty"`
	Url      string `yaml:"url,omitempty" dict:"url,omitempty"`
	UrlTitle string `yaml:"urlTitle,omitempty" dict:"urlTitle,omitempty"`
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
	data := utils.StructToDict(n.MessageParams)
	return DefaultWebhook, ext.Data(data)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp.Ok {
		return nil
	}
	return fmt.Errorf("[PushOver] [%v] %s", resp.StatusCode, resp.Content)
}
