package main

import (
	"flag"
	"github.com/jessevdk/go-flags"
	"github.com/sari3l/notify/internal"
	"github.com/sari3l/notify/types"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input := flag.Args()
	if !terminal.IsTerminal(0) {
		b, err := ioutil.ReadAll(os.Stdin)
		if err == nil {
			input = append(input, string(b))
		}
	}
	data := strings.Join(input, " ")

	opt := new(types.Option)
	if _, err := flags.Parse(opt); err != nil {
		return
	}

	run := internal.InitRunner(opt)
	if run == nil {
		return
	}
	internal.Parse(run, opt)
	run.Run([]string{data})
}
