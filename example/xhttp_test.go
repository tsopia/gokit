package example

import (
	"context"
	"fmt"
	"github.com/tsopia/gokit/log"
	"github.com/tsopia/gokit/xhttp"
	"testing"
)

func TestNewHttpClient(t *testing.T) {

	var result MyStruct
	ctx := log.With().
		Str("request_id", "123").
		Str("user_id", "abc").
		Logger().WithContext(context.Background())
	err := xhttp.Get(ctx, "https://chat.openai.co1m/c/c60d0aec-3688-4683-a739-cf4428a7a537", &result)
	//err := xhttp.Get(ctx, "https://api.oioweb.cn/api/weather/GetWeather", &result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Field1:", result.Field1)
	fmt.Println("Field2:", result.Field2)

}

type MyStruct struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}
