package telegram

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
	ChatId                   string          `yaml:"chatId" json:"chat_id"`
	Text                     string          `yaml:"text" json:"text"`
	ParseMode                *string         `yaml:"parseMode,omitempty" json:"parse_mode,omitempty"`
	Entities                 *map[string]any `yaml:"entities,omitempty" json:"entities,omitempty,omitempty"`
	DisableWebPagePreview    *bool           `yaml:"disableWebPagePreview,omitempty" json:"disable_web_page_preview,omitempty"`
	DisableNotification      *bool           `yaml:"disableNotification,omitempty" json:"disable_notification,omitempty"`
	ProtectContent           *bool           `yaml:"protectContent,omitempty" json:"protect_content,omitempty"`
	ReplyToMessageId         *bool           `yaml:"replyToMessageId,omitempty" json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool           `yaml:"allowSendingWithoutReply,omitempty" json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              *map[string]any `yaml:"replyMarkup,omitempty" json:"reply_markup,omitempty"`
}

type Entity struct {
	Type     string  `yaml:"type" json:"type"`
	Offset   int     `yaml:"offset" json:"offset"`
	Length   int     `yaml:"length" json:"length"`
	Url      *string `yaml:"url,omitempty" json:"url,omitempty"`
	User     *User   `yaml:"user,omitempty" json:"user,omitempty"`
	Language *string `yaml:"language,omitempty" json:"language,omitempty"`
}

type User struct {
	Id                      string  `yaml:"id" json:"id"`
	IsBot                   bool    `yaml:"isBot" json:"is_bot"`
	FirstName               string  `yaml:"firstName" json:"first_name"`
	LastName                *string `yaml:"lastName,omitempty" json:"last_name,omitempty"`
	Username                *string `yaml:"username,omitempty" json:"username,omitempty"`
	LanguageCode            *string `yaml:"languageCode,omitempty" json:"language_code,omitempty"`
	CanJoinGroups           *bool   `yaml:"canJoinGroups,omitempty" json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages *bool   `yaml:"canReadAllGroup,omitempty" json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   *bool   `yaml:"supportsInlineQueries,omitempty" json:"supports_inline_queries,omitempty"`
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
	resp := requests.Get(n.format(messages))
	if resp != nil && resp.Ok && resp.Json().Get("ok").Bool() == true {
		return nil
	}
	return fmt.Errorf("[Telegram] [%v] %s", resp.StatusCode, resp.Raw)
}
