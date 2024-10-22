package chanify

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

type Notifier struct {
	*Option
}

func (opt *Option) ToNotifier() *Notifier {
	noticer := &Notifier{}
	noticer.Option = opt
	return noticer
}

func (n *Notifier) format(messages []string) (string, rTypes.Ext) {
	formatMap := utils.GenerateMap(n.NotifyFormatter, messages)
	data := utils.FormatAnyWithMap(n.MessageParams, formatMap)
	json := utils.StructToJson(data)
	return n.Webhook, ext.Json(json)
}

func (n *Notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	return utils.RespCheck("Chanify", resp, func(request *requests.Response) bool {
		return resp.Ok && resp.Json().Get("request-uid").Str != ""
	})
}
