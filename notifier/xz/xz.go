package xz

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
	Title   string  `yaml:"title" json:"title"`
	Content *string `yaml:"content,omitempty" json:"content,omitempty"`
	Type    *string `yaml:"type,omitempty" json:"type,omitempty"` // 立即 null | 每月 m ｜ 每周 w | 每天 d ｜ 每小时 h | 每分钟 mm
	Time    *string `yaml:"time,omitempty" json:"time,omitempty"` // 月周天 "24:60" | 小时 "60"
	Date    *string `yaml:"date,omitempty" json:"date,omitempty"` // 月 1~31 ｜ 周 0-6
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
	json := utils.StructToJson(n.MessageParams)
	return n.Webhook, ext.Json(json)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp != nil && resp.Ok && resp.Json().Get("code").Int() == 200 {
		return utils.InfoCallBack(resp, nil)
	}
	return utils.InfoCallBack(resp, fmt.Errorf("[XZ] [%v] %s", resp.StatusCode, resp.Content))
}
