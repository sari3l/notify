package internal

import (
	"fmt"
	"github.com/sari3l/notify/notifier"
	"github.com/sari3l/notify/types"
	"github.com/sari3l/notify/utils"
	"log"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
	"sync"
)

type Runner struct {
	Noticers map[int]*types.Notifier
	Client   *types.Client
}

func InitRunner(opt *types.Option) *Runner {
	runner := &Runner{
		Noticers: make(map[int]*types.Notifier, 0),
	}
	client := &types.Client{
		IdMap:    make(map[int]*types.Notifier, 0),
		LevelMap: make(map[int][]*types.Notifier, 0),
		NameMap:  make(map[string][]*types.Notifier, 0),
	}
	pro := new(notifier.NotifiesPackage)

	if opt.Config == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Print(err)
			return nil
		}
		opt.Config = path.Join(home, types.DefaultConfigFilePath)
	}

	if err := utils.ReadFromYaml(opt.Config, pro); err != nil {
		log.Fatal(err)
		return nil
	}

	index := 0
	ref := reflect.ValueOf(*pro)
	for i := 0; i < ref.NumField(); i++ {
		field := ref.Field(i)
		name := strings.ToLower(ref.Type().Field(i).Name)
		for j := 0; j < field.Len(); j++ {
			obj := field.Index(j)
			val := obj.MethodByName("ToNotifier").Call(nil)
			n := val[0].Interface().(types.Notifier)
			baseOpt := obj.Elem().FieldByName("BaseOption").Interface().(types.BaseOption)

			client.IdMap[index] = &n
			index++
			client.NameMap[name] = append(client.NameMap[name], &n)
			client.LevelMap[baseOpt.NotifyLevel] = append(client.LevelMap[baseOpt.NotifyLevel], &n)
		}
	}
	runner.Client = client
	return runner
}

func (runner *Runner) registerNoticer(noticers ...*types.Notifier) {
	for _, noticer := range noticers {
		obj := *noticer
		addr, _ := strconv.Atoi(fmt.Sprintf("%d", &obj))
		runner.Noticers[addr] = noticer
	}
}

func (runner *Runner) Run(data *[]string) {
	var wg sync.WaitGroup
	wg.Add(len(runner.Noticers))
	for _, v := range runner.Noticers {
		noticer := *v
		go func() {
			err := noticer.Send(*data)
			if err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
