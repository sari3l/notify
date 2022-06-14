package utils

import (
	"bufio"
	"os"
)

func OutputString(content string) {
	f := bufio.NewWriter(os.Stdout)
	defer func(f *bufio.Writer) {
		err := f.Flush()
		if err != nil {
			return
		}
	}(f)
	_, err := f.Write([]byte(content))
	if err != nil {
		return
	}
}
