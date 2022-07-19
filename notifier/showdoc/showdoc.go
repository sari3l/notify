package showdoc

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
	Title   string `yaml:"title" json:"title"`
	Content string `yaml:"content" json:"content"`
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
	utils.FormatAnyWithMap(&n.MessageParams, &formatMap)
	json := utils.StructToJson(n.MessageParams)
	return n.Webhook, ext.Json(json)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp != nil && resp.Json().Get("error_code").Int() == 0 {
		return utils.InfoCallBack(resp, nil)
	}
	return utils.InfoCallBack(resp, fmt.Errorf("[ShowDoc] [%v] %s", resp.StatusCode, resp.Content))
}