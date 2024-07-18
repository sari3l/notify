package utils

import (
	"fmt"
	"github.com/sari3l/requests"
)

func RespCheck(name string, resp *requests.Response, checkFunc func(response *requests.Response) bool) (err error) {
	if resp == nil {
		return fmt.Errorf("[%s] request error", name)
	}
	ok := checkFunc(resp)
	if !ok {
		return fmt.Errorf("[%s] response error [%v] %s ", name, resp.StatusCode, resp.Text())
	}
	return nil
}
