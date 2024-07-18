package webhook

import (
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	rTypes "github.com/sari3l/requests/types"
)

type Option struct {
	types.BaseOption `yaml:",inline"`
	Webhook          string            `yaml:"webhook"`
	Method           string            `yaml:"method"`
	Params           map[string]string `yaml:"params,omitempty"`
	Data             map[string]string `yaml:"data,omitempty"`
	Json             map[string]any    `yaml:"json,omitempty"`
}

type Notifier struct {
	*Option
}

func (opt *Option) ToNotifier() *Notifier {
	noticer := &Notifier{}
	noticer.Option = opt
	return noticer
}

func (n *Notifier) format(messages []string) (string, rTypes.Dict, rTypes.Dict, map[string]any) {
	formatMap := utils.GenerateMap(n.NotifyFormatter, messages)
	// webhook url 需要考虑要不要加一个urlencode
	utils.FormatAnyWithMap(&n.Webhook, &formatMap)
	utils.FormatAnyWithMap(&n.Params, &formatMap)
	utils.FormatAnyWithMap(&n.Data, &formatMap)
	utils.FormatAnyWithMap(&n.Json, &formatMap)
	return n.Webhook, n.Params, n.Data, n.Json
}

func (n *Notifier) Send(messages []string) error {
	url, params, form, json := n.format(messages)
	session := requests.Session{}
	_, prep := requests.PrepareRequest(n.Method, "GET", url, params, nil, nil, form, json, nil, nil, nil)
	resp := session.Send(prep)
	return utils.RespCheck("Custom", resp, func(request *requests.Response) bool {
		return resp.Ok
	})
}
