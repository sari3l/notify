package mailgun

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
	ApiKey           string `yaml:"apiKey"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	From    string `yaml:"from" dict:"from"`
	To      string `yaml:"to" dict:"to"`
	Subject string `yaml:"subject" dict:"subject"`
	Text    string `yaml:"text" dict:"text"`
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
	dict := utils.StructToDict(data)
	auth := rTypes.BasicAuth{Username: "api", Password: n.ApiKey}
	return n.Webhook, ext.Auth(auth), ext.Form(dict)
}

func (n *Notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	return utils.RespCheck("Mailgun", resp, func(request *requests.Response) bool {
		return resp.Ok
	})
}
