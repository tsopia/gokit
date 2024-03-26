package xhttp

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
	"github.com/tsopia/gokit/log"
	"time"
)

type RestyLoggerAdapter struct {
	zerolog.Logger
}

func zerologClient() *RestyLoggerAdapter {
	logger := log.Logger
	return &RestyLoggerAdapter{*logger}
}

func (l *RestyLoggerAdapter) Errorf(format string, v ...interface{}) {
	l.Error().Msgf(format, v...)
}

func (l *RestyLoggerAdapter) Warnf(format string, v ...interface{}) {
	l.Warn().Msgf(format, v...)
}

func (l *RestyLoggerAdapter) Debugf(format string, v ...interface{}) {
	l.Debug().Msgf(format, v...)
}

func (l *RestyLoggerAdapter) Infof(format string, v ...interface{}) {
	l.Info().Msgf(format, v...)

}

// Client returns a new resty.Client with some default configuration.
func Client() *resty.Client {
	client := resty.New()
	//client.SetLogger(zerologClient())

	client.SetRetryCount(3) // 设置失败重试次数
	// 这里可以设置一些默认的配置，例如超时时间、重试次数等
	client.SetTimeout(30 * time.Second)             // 设置超时时间
	client.SetRetryWaitTime(100 * time.Millisecond) // 设置重试间隔时间
	client.SetRetryMaxWaitTime(2 * time.Second)

	client.OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
		if r.IsError() {
			log.Error(r.Request.Context()).Msgf("Request to %s failed with status %s\"", r.Request.URL, r.Status())
		}
		return nil
	})
	client.SetDebug(true)
	return client
}

// Get sends a GET request to the specified URL and deserializes the response into v.
func Get(ctx context.Context, url string, v interface{}) error {
	resp, err := Client().R().SetContext(ctx).SetResult(v).Get(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("GET %s returned status %s", url, resp.Status())
	}
	return nil
}

// Post sends a POST request to the specified URL with the provided body and deserializes the response into v.
func Post(ctx context.Context, url string, body interface{}, v interface{}) error {
	resp, err := Client().R().SetContext(ctx).SetBody(body).SetResult(v).Post(url)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("POST %s returned status %s", url, resp.Status())
	}
	return nil
}
