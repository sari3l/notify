package bark

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
	Sound     string `yaml:"sound,omitempty" dict:"sound,omitempty"`
	IsArchive *int   `yaml:"isArchive,omitempty" dict:"isArchive,omitempty"`
	Icon      string `yaml:"icon,omitempty" dict:"icon,omitempty"`
	Group     string `yaml:"group,omitempty" dict:"group,omitempty"`
	Level     string `yaml:"level,omitempty" dict:"level,omitempty"`
	Url       string `yaml:"url,omitempty" dict:"url,omitempty"`
	Copy      string `yaml:"copy,omitempty" dict:"copy,omitempty"`
	Badge     *int   `yaml:"badge,omitempty" dict:"badge,omitempty"`
	AutoCopy  *int   `yaml:"autoCopy,omitempty" dict:"autoCopy,omitempty"`
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
	params := utils.StructToDict(n.MessageParams)
	return n.Webhook, ext.Params(params)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Get(n.format(messages))
	if resp != nil && resp.Ok {
		return nil
	}
	return fmt.Errorf("[Bark] [%v] %s", resp.StatusCode, resp.Content)
}
