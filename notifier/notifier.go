package notifier

import (
	"github.com/sari3l/notify/notifier/bark"
	"github.com/sari3l/notify/notifier/chanify"
	"github.com/sari3l/notify/notifier/dingtalk"
	"github.com/sari3l/notify/notifier/discord"
	"github.com/sari3l/notify/notifier/feishu"
	"github.com/sari3l/notify/notifier/ftqq"
	"github.com/sari3l/notify/notifier/gitter"
	"github.com/sari3l/notify/notifier/googlechat"
	"github.com/sari3l/notify/notifier/igot"
	"github.com/sari3l/notify/notifier/mailgun"
	"github.com/sari3l/notify/notifier/pushbullet"
	"github.com/sari3l/notify/notifier/pushdeer"
	"github.com/sari3l/notify/notifier/pushover"
	"github.com/sari3l/notify/notifier/qpush"
	"github.com/sari3l/notify/notifier/rocketchat"
	"github.com/sari3l/notify/notifier/slack"
	"github.com/sari3l/notify/notifier/telegram"
	"github.com/sari3l/notify/notifier/webhook"
	"github.com/sari3l/notify/notifier/xz"
	"github.com/sari3l/notify/notifier/zulip"
)

type NotifiesPackage struct {
	Bark       []*bark.Option       `yaml:"bark,omitempty"`
	Chanify    []*chanify.Option    `yaml:"chanify,omitempty"`
	Dingtalk   []*dingtalk.Option   `yaml:"dingtalk,omitempty"`
	Discord    []*discord.Option    `yaml:"discord,omitempty"`
	FeiShu     []*feishu.Option     `yaml:"feishu,omitempty"`
	FTQQ       []*ftqq.Option       `yaml:"ftqq,omitempty"`
	Gitter     []*gitter.Option     `yaml:"gitter,omitempty"`
	GoogleChat []*googlechat.Option `yaml:"googlechat,omitempty"`
	IGot       []*igot.Option       `yaml:"igot,omitempty"`
	Mailgun    []*mailgun.Option    `yaml:"mailgun,omitempty"`
	PushBullet []*pushbullet.Option `yaml:"pushbullet,omitempty"`
	PushDeer   []*pushdeer.Option   `yaml:"pushdeer,omitempty"`
	PushOver   []*pushover.Option   `yaml:"pushover,omitempty"`
	QPush      []*qpush.Option      `yaml:"qpush,omitempty"`
	RocketChat []*rocketchat.Option `yaml:"rocketchat,omitempty"`
	Slack      []*slack.Option      `yaml:"slack,omitempty"`
	Telegram   []*telegram.Option   `yaml:"telegram,omitempty"`
	XZ         []*xz.Option         `yaml:"xz,omitempty"`
	Webhook    []*webhook.Option    `yaml:"webhook,omitempty"`
	ZuLip      []*zulip.Option      `yaml:"zulip,omitempty"`

	// 计划中
	// Sendgraid  []*sendgraid.Option  `yaml:"sendgraid,omitempty"`
	// Mattermost []*mattermost.Option `yaml:"mattermost,omitempty" //https://api.mattermost.com/
	// NowPush []*nowpush.Option `yaml:"nowpush,omitempty"` //https://www.web.nowpush.app/
}
