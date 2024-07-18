package pushdeer

import (
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
	rTypes "github.com/sari3l/requests/types"
)

const DefaultWebhook = "https://api2.pushdeer.com/message/push"

type Option struct {
	types.BaseOption `yaml:",inline"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	PushKey     string  `yaml:"pushKey" dict:"pushkey"`
	Type        *string `yaml:"type,omitempty" dict:"type,omitempty"`
	Text        *string `yaml:"text,omitempty" dict:"text,omitempty"`
	Description *string `yaml:"desc,omitempty" dict:"desc,omitempty"`
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
	utils.FormatAnyWithMap(&n.MessageParams, &formatMap)
	data := utils.StructToDict(n.MessageParams)
	return DefaultWebhook, ext.Form(data)
}

func (n *Notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	return utils.RespCheck("PushDeer", resp, func(request *requests.Response) bool {
		return resp.Ok
	})
}
