package ftqq

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
	Title       string  `yaml:"title" dict:"title"`
	Description *string `yaml:"desp,omitempty" dict:"desp,omitempty"`
	Channel     *string `yaml:"channel,omitempty" dict:"channel,omitempty"`
	OpenId      *string `yaml:"openId,omitempty" dict:"openId,omitempty"`
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
	dict := utils.StructToDict(data)
	return n.Webhook, ext.Form(dict)
}

func (n *Notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	return utils.RespCheck("FTQQ", resp, func(request *requests.Response) bool {
		return resp.Ok
	})
}
