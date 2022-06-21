package feishu

import (
	"fmt"
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
	"time"
)

type Option struct {
	types.BaseOption `yaml:",inline"`
	Webhook          string `yaml:"webhook"`
	Secret           string `yaml:"secret,omitempty"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	MsgType   string         `yaml:"msgType" json:"msg_type"`
	Content   map[string]any `yaml:"content" json:"content"`
	Timestamp *string        `yaml:"timestamp,omitempty" json:"timestamp,omitempty"`
	Sign      *string        `json:"sign,omitempty"`
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
	if n.Secret != "" {
		timestamp := time.Now().UnixNano() / int64(time.Millisecond)
		sha256 := utils.HmacSha256(fmt.Sprintf("%d\n%s", timestamp, n.Secret), n.Secret)
		timestampStr := fmt.Sprintf("%d", timestamp)
		n.Timestamp = &timestampStr
		n.Sign = &sha256
	}
	formatMap := utils.GenerateMap(n.NotifyFormatter, messages)
	utils.FormatAnyWithMap(&n.MessageParams, &formatMap)
	json := utils.StructToJson(n.MessageParams)
	return n.Webhook, ext.Json(json)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp != nil && resp.Ok && resp.Json().Get("code").Int() == 0 {
		return nil
	}
	return fmt.Errorf("[FeiShu] [%v] %s", resp.StatusCode, resp.Content)
}
