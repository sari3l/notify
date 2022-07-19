package mailgun

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
	ApiKey           string `yaml:"apiKey"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	From    string `yaml:"from" dict:"from"`
	To      string `yaml:"to" dict:"to"`
	Subject string `yaml:"subject" dict:"subject"`
	Text    string `yaml:"text" dict:"text"`
}

type notifier struct {
	*Option
}

func (opt *Option) ToNotifier() *notifier {
	noticer := &notifier{}
	noticer.Option = opt
	return noticer
}

func (n *notifier) format(messages []string) (string, ext.Ext, ext.Ext) {
	formatMap := utils.GenerateMap(n.NotifyFormatter, messages)
	utils.FormatAnyWithMap(&n.MessageParams, &formatMap)
	data := utils.StructToDict(n.MessageParams)
	auth := ext.BasicAuth{Username: "api", Password: n.ApiKey}
	return n.Webhook, ext.Auth(auth), ext.Data(data)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp != nil && resp.Ok {
		return utils.InfoCallBack(resp, nil)
	}
	return utils.InfoCallBack(resp, fmt.Errorf("[Mailgun] [%v] %s", resp.StatusCode, resp.Content))
}
