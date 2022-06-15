package notifier

import (
	"github.com/sari3l/notify/notifier/bark"
	"github.com/sari3l/notify/notifier/chanify"
	"github.com/sari3l/notify/notifier/dingtalk"
	"github.com/sari3l/notify/notifier/discord"
	"github.com/sari3l/notify/notifier/googlechat"
	"github.com/sari3l/notify/notifier/mailgun"
	"github.com/sari3l/notify/notifier/pushbullet"
	"github.com/sari3l/notify/notifier/pushdeer"
	"github.com/sari3l/notify/notifier/pushover"
	"github.com/sari3l/notify/notifier/rocketchat"
	"github.com/sari3l/notify/notifier/telegram"
	"github.com/sari3l/notify/notifier/webhook"
)

type NotifiesPackage struct {
	Bark       []*bark.Option       `yaml:"bark,omitempty"`
	Chanify    []*chanify.Option    `yaml:"chanify,omitempty"`
	Dingtalk   []*dingtalk.Option   `yaml:"dingtalk,omitempty"`
	Discord    []*discord.Option    `yaml:"discord,omitempty"`
	GoogleChat []*googlechat.Option `yaml:"googlechat,omitempty"`
	Mailgun    []*mailgun.Option    `yaml:"mailgun,omitempty"`
	PushBullet []*pushbullet.Option `yaml:"pushbullet,omitempty"`
	PushDeer   []*pushdeer.Option   `yaml:"pushdeer,omitempty"`
	PushOver   []*pushover.Option   `yaml:"pushover,omitempty"`
	Rocketchat []*rocketchat.Option `yaml:"rocketchat,omitempty"`
	//Sendgraid  []*string            `yaml:"sendgraid,omitempty"`
	Telegram []*telegram.Option `yaml:"telegram,omitempty"`
	Webhook  []*webhook.Option  `yaml:"webhook,omitempty"`
}
