package example

import (
	"github.com/tsopia/gokit/xhttp"
	"testing"
)

func TestNewHttpClient(t *testing.T) {
	c := xhttp.NewHttpClient()
	r, _ := c.EnableTrace().Get("/sdaa")
	t.Logf(r.Request.URL)

}
