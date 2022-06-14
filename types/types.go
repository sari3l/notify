package types

const DefaultConfigFilePath = ".config/notify-config.yaml"
const VersionInfo = "v0.0.1 from https://github.com/sari3l/notify\n"

type Option struct {
	Show       bool     `short:"s" long:"show" description:"显示配置信息" required:"false"`
	Config     string   `short:"c" long:"config" description:"指定配置文件, 如: $HOME/.config/notify-config.yaml" required:"false"`
	Id         []int    `short:"i" long:"id" description:"指定通知id" required:"false"`
	Level      []int    `short:"l" long:"level" description:"指定通告等级" default:"0" required:"false"`
	AboveLevel *int     `short:"a" long:"aboveLevel" description:"指定最低通告等级" required:"false"`
	Noticer    []string `short:"n" long:"noticer" description:"指定通知模块" required:"false"`
	Version    bool     `short:"v" long:"version" description:"版本信息" required:"false"`
}

type Client struct {
	IdMap    map[int]*Notifier
	LevelMap map[int][]*Notifier
	NameMap  map[string][]*Notifier
}

type Notifier interface {
	Send(data []string) error
}

type BaseOption struct {
	NotifyLevel       int      `yaml:"notifyLevel,omitempty"`
	NotifyFormatter   []string `yaml:"notifyFormatter,omitempty"`
	NotifyDescription string   `yaml:"notifyDescription,omitempty"`
}
