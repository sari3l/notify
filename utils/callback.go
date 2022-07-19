package utils

import (
	"errors"
	"fmt"
	"github.com/sari3l/requests"
	"io"
	"strings"
	"time"
)

func InfoCallBack(resp *requests.Response, data any) error {
	if RunMode == 1 {
		fmt.Printf("[%s] %s %s", time.Now(), resp.Request.URL, resp.Content)
	} else if RunMode == 2 {
		var headersReq, headerResp string
		for k, v := range resp.Request.Header {
			headersReq += fmt.Sprintf("%s: %s\n", k, strings.Join(v, ""))
		}
		for k, v := range resp.Header {
			headerResp += fmt.Sprintf("%s: %s\n", k, strings.Join(v, ""))
		}
		buf := new(strings.Builder)
		if resp.Request.Body != nil {
			io.Copy(buf, resp.Request.Body)
		}

		fmt.Printf(
			"=============== %s ===============\n"+
				">> Request: %s\n"+
				"%s %s %s\n"+
				"%s\n"+
				"\n"+
				">> Response: \n"+
				"%s"+
				"%s %s\n"+
				"%s\n"+
				"%s",
			time.Now(), resp.Request.URL.String(), resp.Request.Method, resp.Request.URL.Path, resp.Request.Proto, headersReq, buf.String(), resp.Proto, resp.Status, headerResp, resp.Content)
	}
	if data == nil {
		return nil
	} else {
		return errors.New(data.(string))
	}
}
