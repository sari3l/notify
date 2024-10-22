package gitter

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
	Token            string `yaml:"token"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	Text string `yaml:"text" json:"text"`
}

type Notifier struct {
	*Option
}

func (opt *Option) ToNotifier() *Notifier {
	noticer := &Notifier{}
	noticer.Option = opt
	return noticer
}

func (n *Notifier) format(messages []string) (string, rTypes.Ext, rTypes.Ext) {
	formatMap := utils.GenerateMap(n.NotifyFormatter, messages)
	data := utils.FormatAnyWithMap(n.MessageParams, formatMap)
	auth := rTypes.BearerAuth{Token: n.Token}
	json := utils.StructToJson(data)
	return n.Webhook, ext.Auth(auth), ext.Json(json)
}

func (n *Notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	return utils.RespCheck("Gitter", resp, func(request *requests.Response) bool {
		return resp.Json().Get("error").Str == ""
	})
}
