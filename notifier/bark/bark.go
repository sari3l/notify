package bark

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
	Sound     *string `yaml:"sound,omitempty" dict:"sound,omitempty"`
	IsArchive *int    `yaml:"isArchive,omitempty" dict:"isArchive,omitempty"`
	Icon      *string `yaml:"icon,omitempty" dict:"icon,omitempty"`
	Group     *string `yaml:"group,omitempty" dict:"group,omitempty"`
	Level     *string `yaml:"level,omitempty" dict:"level,omitempty"`
	Url       *string `yaml:"url,omitempty" dict:"url,omitempty"`
	Copy      *string `yaml:"copy,omitempty" dict:"copy,omitempty"`
	Badge     *int    `yaml:"badge,omitempty" dict:"badge,omitempty"`
	AutoCopy  *int    `yaml:"autoCopy,omitempty" dict:"autoCopy,omitempty"`
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
	url := utils.FormatAnyWithMap(n.Webhook, formatMap)
	data := utils.FormatAnyWithMap(n.MessageParams, formatMap)
	params := utils.StructToDict(data)
	return url.(string), ext.Params(params)
}

func (n *Notifier) Send(messages []string) error {
	resp := requests.Get(n.format(messages))
	return utils.RespCheck("Bark", resp, func(request *requests.Response) bool {
		return resp.Ok
	})
}
