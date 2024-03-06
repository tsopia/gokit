package xhttp

import (
	"github.com/go-resty/resty/v2"
	"time"
)

func NewHttpClient() *resty.Request {
	client := resty.New()
	client.SetTimeout(30 * time.Second)
	client.BaseURL = "http://localhost:8080"
	return client.R()
}

var client = resty.New()

func Get(url string) (*resty.Response, error) {
	return client.R().Get(url)
}
func R() *resty.Request {
	return client.R()
}
