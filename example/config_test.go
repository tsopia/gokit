package example

import (
	"github.com/spf13/viper"
	"github.com/tsopia/gokit/conf"
	"os"
	"testing"
)

func TestInitConfig(t *testing.T) {
	// 初始化 ViperConfigManager
	//os.Setenv("PORT", "8080")
	//os.Setenv("HOST", "DDD")
	//os.Setenv("PATH", "pass")
	director, _ := os.Getwd()
	vcm := &conf.ViperConfigManager{
		ConfigPath: director,
		ConfigType: "yaml",
		ConfigName: "config",
	}

	_ = vcm.InitConfig()
	configStruct := conf.DefaultConf
	t.Log(configStruct)
	t.Log(configStruct.Port)
	t.Log(*configStruct.Mysql.Host)
	t.Log(*configStruct.Mysql.Host)
	t.Log(viper.GetString("mysql.host"))
	t.Log(viper.GetString("Port"))
	t.Log(viper.GetString("serverName"))

	// 现在可以通过managerConf.Mysql.Host访问host的值
}

type AppConfig struct {
	DatabaseURL string `mapstructure:"DB_URL"`
	ServerPort  int    `mapstructure:"SERVER_PORT"`
	DebugMode   bool   `mapstructure:"DEBUG_MODE"`
}

//func bindEnvForStruct(v *viper.Viper, prefix string, s interface{}) {
//	val := reflect.ValueOf(s).Elem()
//	typ := val.Type()
//
//	for i := 0; i < typ.NumField(); i++ {
//		field := typ.Field(i)
//		tag := field.Tag.Get("mapstructure")
//		envKey := fmt.Sprintf("%s%s", prefix, tag)
//
//		if envVal := v.Get(envKey); envVal != nil {
//			v.BindEnv(envKey)
//		}
//	}
//}
