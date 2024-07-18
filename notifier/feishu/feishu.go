package feishu

import (
	"fmt"
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
	rTypes "github.com/sari3l/requests/types"
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

type Notifier struct {
	*Option
}

func (opt *Option) ToNotifier() *Notifier {
	noticer := &Notifier{}
	noticer.Option = opt
	return noticer
}

func (n *Notifier) format(messages []string) (string, rTypes.Ext) {
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

func (n *Notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	return utils.RespCheck("FeiShu", resp, func(request *requests.Response) bool {
		return resp.Ok && resp.Json().Get("code").Int() == 0
	})
}
