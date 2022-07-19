package internal

import (
	"fmt"
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"reflect"
	"strings"
)

func Parse(runner *Runner, opt *types.Option) {
	// -v --version
	if opt.Version {
		utils.OutputString(utils.VersionInfo)
	}
	// -m model
	if opt.Mode != nil {
		utils.RunMode = *opt.Mode
	}
	// -s --show
	if opt.Show {
		var content string
		content += fmt.Sprintf("%-4s\t%-5s\t%-15s\t%-20s\t%s\n", "ID", "Level", "Type", "Formatter", "Description")
		for i := 0; i < len(runner.Client.IdMap); i++ {
			option := reflect.ValueOf(*runner.Client.IdMap[i]).Elem().FieldByName("Option")
			level := option.Elem().FieldByName("NotifyLevel").Int()
			formatterField := option.Elem().FieldByName("NotifyFormatter")
			formatterSlice := make([]string, 0)
			for i := 0; i < formatterField.Len(); i++ {
				formatterSlice = append(formatterSlice, formatterField.Index(i).String())
			}
			description := option.Elem().FieldByName("BaseOption").FieldByName("NotifyDescription").String()
			path := strings.Split(reflect.TypeOf(option.Interface()).Elem().PkgPath(), "/")
			content += fmt.Sprintf("[%2d]\t[%3d]\t%-15s\t%-20v\t%+v \n", i, level, path[len(path)-1], strings.Join(formatterSlice, ","), description)
		}
		utils.OutputString(content)
	}
	// -i --id
	if len(opt.Id) != 0 {
		for _, id := range opt.Id {
			if runner.Client.IdMap[id] != nil {
				runner.registerNoticer(runner.Client.IdMap[id])
			}
		}
	}
	// -l --level
	if len(opt.Level) != 0 {
		for _, level := range opt.Level {
			runner.registerNoticer(runner.Client.LevelMap[level]...)
		}
	}
	// -al --aboveLevel
	if opt.AboveLevel != nil {
		for level, noticers := range runner.Client.LevelMap {
			if level >= *opt.AboveLevel {
				runner.registerNoticer(noticers...)
			}
		}
	}
	// -n --noticer
	if len(opt.Noticer) != 0 {
		for _, name := range opt.Noticer {
			runner.registerNoticer(runner.Client.NameMap[strings.ToLower(name)]...)
		}
	}
}
