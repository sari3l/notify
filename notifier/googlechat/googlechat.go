package googlechat

import (
	"fmt"
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
	Text string `yaml:"text" json:"text"`
}

type notifier struct {
	*Option
}

func (opt *Option) ToNotifier() *notifier {
	noticer := &notifier{}
	noticer.Option = opt
	return noticer
}

func (n *notifier) format(messages []string) (string, rTypes.Ext) {
	formatMap := utils.GenerateMap(n.NotifyFormatter, messages)
	utils.FormatAnyWithMap(&n.MessageParams, &formatMap)
	data := utils.StructToJson(n.MessageParams)
	return n.Webhook, ext.Json(data)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp != nil && resp.Ok {
		return utils.InfoCallBack(resp, nil)
	}
	return utils.InfoCallBack(resp, fmt.Errorf("[GoogleChat] [%v] %s", resp.StatusCode, resp.Content))
}
