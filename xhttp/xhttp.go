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
