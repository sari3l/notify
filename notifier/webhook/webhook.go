package webhook

import (
	"fmt"
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
)

type Option struct {
	types.BaseOption `yaml:",inline"`
	Webhook          string            `yaml:"webhook"`
	Method           string            `yaml:"method"`
	Params           map[string]string `yaml:"params,omitempty"`
	Data             map[string]string `yaml:"data,omitempty"`
	Json             map[string]any    `yaml:"json,omitempty"`
}

type notifier struct {
	*Option
}

func (opt *Option) ToNotifier() *notifier {
	noticer := &notifier{}
	noticer.Option = opt
	return noticer
}

func (n *notifier) format(messages []string) (string, ext.Dict, ext.Dict, map[string]any) {
	formatMap := utils.GenerateMap(n.NotifyFormatter, messages)
	// webhook url 需要考虑要不要加一个urlencode
	utils.FormatAnyWithMap(&n.Webhook, formatMap)
	utils.FormatAnyWithMap(&n.Params, formatMap)
	utils.FormatAnyWithMap(&n.Data, formatMap)
	utils.FormatAnyWithMap(&n.Json, formatMap)
	return n.Webhook, n.Params, n.Data, n.Json
}

func (n *notifier) Send(messages []string) error {
	url, params, data, json := n.format(messages)
	session := requests.Session(5, "", false, false)
	_, prep := requests.PrepareRequest(n.Method, url, params, nil, nil, data, json, nil, nil, nil, nil)
	resp := session.Send(prep)
	if resp != nil && resp.Ok {
		return nil
	}
	if resp == nil {
		return fmt.Errorf("[Custom] connection refused\n")
	}
	return fmt.Errorf("[Custom] [%v] %s", resp.StatusCode, resp.Content)
}
