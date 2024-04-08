package example

import (
	"context"
	"fmt"
	"github.com/tsopia/gokit/conf"
	"github.com/tsopia/gokit/xhttp"
	"github.com/tsopia/gokit/xlog"
	"os"
	"testing"
)

func TestNewHttpClient(t *testing.T) {
	//director, _ := os.Getwd()
	//time.Sleep(3 * time.Second)
	//t.Logf(director)
	os.Setenv("GO_RESET_DEBUG", "true")
	vcm := &conf.ViperConfigManager{
		//ConfigPath: director,
		//ConfigType: "yaml",
		//ConfigName: "config",
	}
	////初始化viper
	_ = vcm.InitConfig(nil)
	//t.Logf("config: %+v", viper.Get("GO_RESET_DEBUG"))
	//
	//t.Logf("config: %+v", os.Getenv("GO_RESET_DEBUG"))
	var result MyStruct
	ctx := xlog.With().
		Str("request_id", "123").
		Str("user_id", "abc").
		Logger().WithContext(context.Background())
	//err := xhttp.Get(ctx, "https://chat.openai.com/c/c60d0aec-3688-4683-a739-cf4428a7a537", &result)
	err := xhttp.Get(ctx, "https://api.oioweb.cn/api/weather/GetWeather9/1", &result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	t.Log("==============================================================================\n")

	_ = xhttp.Get(ctx, "https://api.oioweb.cn/api/weather/GetWeather", &result)

	fmt.Println("Field1:", result.Field1)
	fmt.Println("Field2:", result.Field2)
	fmt.Println("Field2:", os.Getenv("GoResetDebug"))

}

type MyStruct struct {
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}
