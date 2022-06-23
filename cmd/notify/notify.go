package main

import (
	"bufio"
	"github.com/jessevdk/go-flags"
	"github.com/sari3l/notify/internal"
	"github.com/sari3l/notify/types"
	"os"
)

func main() {
	data := new([]string)

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			*data = append(*data, scanner.Text())
		}
	}

	opt := new(types.Option)
	args, err := flags.Parse(opt)
	if err != nil {
		return
	}

	if len(args) != 0 {
		*data = append(*data, args...)
	}

	run := internal.InitRunner(opt)
	if run == nil {
		return
	}
	internal.Parse(run, opt)
	run.Run(data)
}
