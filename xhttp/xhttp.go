package xhttp

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tsopia/gokit/log"
	"github.com/tsopia/gokit/xlog"
	"time"
)

// Client returns a new resty.Client with some default configuration.
func Client() *resty.Client {
	client := resty.New()
	client.SetLogger(xlog.New())
	client.SetLogger(log.New())

	client.SetRetryCount(3) // 设置失败重试次数
	// 这里可以设置一些默认的配置，例如超时时间、重试次数等
	client.SetTimeout(30 * time.Second)             // 设置超时时间
	client.SetRetryWaitTime(100 * time.Millisecond) // 设置重试间隔时间
	client.SetRetryMaxWaitTime(2 * time.Second)
	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		if r.IsError() {
			r.Error()
			log.Error(context.Background()).Err(r.Error().(error)).Msgf("Request to %s failed with status %s\", r.Request.URL, r.Status()")
		}
		return nil
	})
	//client.SetDebug(true)
	//todo 设置日志记录器
	//logger := log.New(os.Stdout, "", 0)
	//client.SetLogger(logger)
	return client
}

// Get sends a GET request to the specified URL and deserializes the response into v.
func Get(url string, v interface{}) error {
	resp, err := Client().R().SetResult(v).Get(url)

	if err != nil {

		return err
	}
	if resp.IsError() {
		return fmt.Errorf("GET %s returned status %s", url, resp.Status())
	}
	return nil
}

// Post sends a POST request to the specified URL with the provided body and deserializes the response into v.
func Post(url string, body interface{}, v interface{}) error {
	resp, err := Client().R().SetBody(body).SetResult(v).Post(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("POST %s returned status %s", url, resp.Status())
	}
	return nil
}
