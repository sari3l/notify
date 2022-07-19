package chanify

import (
	"fmt"
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
)

type Option struct {
	types.BaseOption `yaml:",inline"`
	Webhook          string `yaml:"webhook"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	Text              *string              `yaml:"text,omitempty" json:"text,omitempty"`
	Title             *string              `yaml:"title,omitempty" json:"title,omitempty"`
	Copy              *string              `yaml:"copy,omitempty" json:"copy,omitempty"`
	AutoCopy          *int                 `yaml:"autoCopy,omitempty" json:"autocopy,omitempty"`
	Sound             *int                 `yaml:"sound,omitempty" json:"sound,omitempty"`
	Priority          *int                 `yaml:"priority,omitempty" json:"priority,omitempty"`
	InterruptionLevel *string              `yaml:"interruptionLevel,omitempty" json:"interruptionlevel,omitempty"`
	Actions           *map[string][]string `yaml:"actions,omitempty" json:"actions,omitempty"`
	Timeline          *map[string]any      `yaml:"timeline,omitempty" json:"timeline,omitempty"`
	Link              *string              `yaml:"link,omitempty" json:"link,omitempty"`
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
	utils.FormatAnyWithMap(&n.MessageParams, &formatMap)
	json := utils.StructToJson(n.MessageParams)
	return n.Webhook, ext.Json(json)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp != nil && resp.Ok && resp.Json().Get("request-uid").Str != "" {
		return utils.InfoCallBack(resp, nil)
	}
	return utils.InfoCallBack(resp, fmt.Errorf("[Chanify] [%v] %s", resp.StatusCode, resp.Content))
}
