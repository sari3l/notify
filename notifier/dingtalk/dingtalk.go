package dingtalk

import (
	"fmt"
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
	rTypes "github.com/sari3l/requests/types"
	"time"
)

const DefaultServer = "https://oapi.dingtalk.com/robot/send"
const DefaultMessageType = "text"

type AtOption struct {
	AtMobiles *[]string `yaml:"atMobiles,omitempty" json:"atMobiles,omitempty"`
	AtUserIds *[]string `yaml:"atUserIds,omitempty" json:"atUserIds,omitempty"`
	IsAtAll   *bool     `yaml:"isAtAll,omitempty" json:"isAtAll,omitempty"`
}

type BtnOption struct {
	Title     *string `yaml:"title,omitempty" json:"title,omitempty"`
	ActionURL *string `yaml:"actionURL,omitempty" json:"actionURL,omitempty"`
}

type LinkOption struct {
	Title      *string `yaml:"title,omitempty" json:"title,omitempty"`
	MessageURL *string `yaml:"messageURL,omitempty" json:"messageURL,omitempty"`
	PicURL     *string `yaml:"picURL,omitempty" json:"picURL,omitempty"`
}

type MessageParams struct {
	Text           *string           `yaml:"text,omitempty" json:"text,omitempty"`                     // Markdown | Link | ActionCard
	Title          *string           `yaml:"title,omitempty" json:"title,omitempty"`                   // Markdown | Link | ActionCard
	SingleTitle    *string           `yaml:"singleTitle,omitempty" json:"singleTitle,omitempty"`       // Markdown | Link | ActionCard
	SingleUrl      *string           `yaml:"singleUrl,omitempty" json:"singleUr,omitemptyl"`           // ActionCard
	BtnOrientation *string           `yaml:"btnOrientation,omitempty" json:"btnOrientation,omitempty"` // ActionCard
	Content        *string           `yaml:"content,omitempty" json:"content,omitempty"`               // Text
	PicUrl         *string           `yaml:"picUrl,omitempty" json:"picUrl,omitempty"`                 // Link
	MessageUrl     *string           `yaml:"messageUrl,omitempty" json:"messageUrl,omitempty"`         // Link
	Btns           *[]map[string]any `yaml:"btns,omitempty" json:"btns,omitempty"`                     // ActionCard
	Links          *[]map[string]any `yaml:"links,omitempty" json:"links,omitempty"`                   // FeedCard
}

type Option struct {
	types.BaseOption `yaml:",inline"`
	Token            string `yaml:"token"`
	MessageType      string `yaml:"messageType,omitempty"`
	Secret           string `yaml:"secret,omitempty"`
	AtOption         `yaml:",inline"`
	MessageParams    `yaml:",inline"`
}

type notifier struct {
	*Option
}

func (opt *Option) ToNotifier() *notifier {
	noticer := &notifier{}
	noticer.Option = opt
	return noticer
}

func (n *notifier) format(messages []string) (string, rTypes.Ext, rTypes.Ext) {
	if n.MessageType == "" {
		n.MessageType = DefaultMessageType
	}
	params := rTypes.Dict{
		"access_token": n.Token,
	}
	// 安全-加签处理
	if n.Secret != "" {
		timestamp := time.Now().UnixNano() / int64(time.Millisecond)
		sha256 := utils.HmacSha256(fmt.Sprintf("%d\n%s", timestamp, n.Secret), n.Secret)
		params["timestamp"] = fmt.Sprintf("%d", timestamp)
		params["sign"] = sha256
	}
	// 信息-类型处理
	formatMap := utils.GenerateMap(n.NotifyFormatter, messages)
	utils.FormatAnyWithMap(&n.MessageParams, &formatMap)

	data := map[string]interface{}{
		"msgtype": n.MessageType,
		"at":      utils.StructToJson(n.AtOption),
	}
	data[n.MessageType] = utils.StructToJson(n.MessageParams)
	return DefaultServer, ext.Params(params), ext.Json(data)
}

func (n *notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	if resp != nil && resp.Ok && resp.Json().Get("errcode").Int() == 0 {
		return nil
	}
	return fmt.Errorf("[Dingtalk] [%v] %s", resp.StatusCode, resp.Raw)
}
