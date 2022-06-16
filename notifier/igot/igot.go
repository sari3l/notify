package igot

import (
	"fmt"
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
)

// 文档 http://hellyw.com/#/

type Option struct {
	types.BaseOption `yaml:",inline"`
	Webhook          string `yaml:"webhook"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	Content           string `yaml:"content" json:"content"`
	Title             string `yaml:"title,omitempty" json:"title,omitempty"`
	Url               string `yaml:"url,omitempty" dict:"url,omitempty"`
	Copy              string `yaml:"copy,omitempty" dict:"copy,omitempty"`
	Urgent            *int   `yaml:"urgent,omitempty" dict:"urgent,omitempty"`
	AutomaticallyCopy *int   `yaml:"automaticallyCopy,omitempty" json:"automaticallyCopy,omitempty"`
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
	utils.FormatAnyWithMap(&n.Webhook, &formatMap)
	json := utils.StructToJson(n.MessageParams)
	return n.Webhook, ext.Json(json)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp != nil && resp.Ok && resp.Json().Get("ret").Int() == 0 {
		return nil
	}
	return fmt.Errorf("[iGot] [%v] %s", resp.StatusCode, resp.Content)
}
