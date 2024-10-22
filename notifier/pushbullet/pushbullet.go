package pushbullet

import (
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"github.com/sari3l/requests"
	"github.com/sari3l/requests/ext"
	rTypes "github.com/sari3l/requests/types"
)

const DefaultWebhook = "https://api.pushbullet.com/v2/pushes"

type Option struct {
	types.BaseOption `yaml:",inline"`
	Token            string `yaml:"token"`
	MessageParams    `yaml:",inline"`
}

type MessageParams struct {
	Type                    string    `yaml:"type" json:"type"`
	Iden                    *string   `yaml:"iden,omitempty" json:"iden,omitempty"`
	Active                  *bool     `yaml:"active,omitempty" json:"active,omitempty"`
	Created                 *float32  `yaml:"created,omitempty" json:"created,omitempty"`
	Modified                *float32  `yaml:"modified,omitempty" json:"modified,omitempty"`
	Dismissed               *bool     `yaml:"dismissed,omitempty" json:"dismissed,omitempty"`
	Guid                    *string   `yaml:"guid,omitempty" json:"guid,omitempty"`
	Direction               *string   `yaml:"direction,omitempty" json:"direction,omitempty"`
	SenderIden              *string   `yaml:"senderIden,omitempty" json:"sender_iden,omitempty"`
	SenderEmail             *string   `yaml:"senderEmail,omitempty" json:"sender_email,omitempty"`
	SenderEmailNormalized   *string   `yaml:"senderEmailNormalized,omitempty" json:"sender_email_normalized,omitempty"`
	SenderName              *string   `yaml:"senderName,omitempty" json:"sender_name,omitempty"`
	ReceiverIden            *string   `yaml:"receiverIden,omitempty" json:"receiver_iden,omitempty"`
	ReceiverEmail           *string   `yaml:"receiverEmail,omitempty" json:"receiver_email,omitempty"`
	ReceiverEmailNormalized *string   `yaml:"receiverEmailNormalized,omitempty" json:"receiver_email_normalized,omitempty"`
	TargetDeviceIden        *string   `yaml:"targetDeviceIden,omitempty" json:"target_device_iden,omitempty"`
	SourceDeviceIden        *string   `yaml:"sourceDeviceIden,omitempty" json:"source_device_iden,omitempty"`
	ClientIden              *string   `yaml:"clientIden,omitempty" json:"client_iden,omitempty"`
	ChannelIden             *string   `yaml:"channelIden,omitempty" json:"channel_iden,omitempty"`
	AwakeAppGuids           *[]string `yaml:"awakeAppGuids,omitempty" json:"awake_app_guids,omitempty"`
	Title                   *string   `yaml:"title,omitempty" json:"title,omitempty"`
	Body                    *string   `yaml:"body,omitempty" json:"body,omitempty"`
	Url                     *string   `yaml:"url,omitempty" json:"url,omitempty"`
	FileName                *string   `yaml:"fileName,omitempty" json:"file_name,omitempty"`
	FileType                *string   `yaml:"fileType,omitempty" json:"file_type,omitempty"`
	FileUrl                 *string   `yaml:"fileUrl,omitempty" json:"file_url,omitempty"`
	ImageUrl                *string   `yaml:"imageUrl,omitempty" json:"image_url,omitempty"`
	ImageWidth              *int      `yaml:"imageWidth,omitempty" json:"image_width,omitempty"`
	ImageHeight             *int      `yaml:"imageHeight,omitempty" json:"image_height,omitempty"`
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
	headers := rTypes.Dict{"Access-Token": n.Token}
	json := utils.StructToJson(data)
	return DefaultWebhook, ext.Headers(headers), ext.Json(json)
}

func (n *Notifier) Send(messages []string) error {
	resp := requests.Post(n.format(messages))
	return utils.RespCheck("PushBullet", resp, func(request *requests.Response) bool {
		return resp.Ok
	})
}
