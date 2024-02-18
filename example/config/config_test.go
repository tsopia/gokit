package conf

import (
	"github.com/tsopia/gokit/conf"
	"os"

	"testing"
)

// func TestViperConfigManager(t *testing.T) {
//
// }
// 编写InitConfig方法的测试用例
func TestInitConfig(t *testing.T) {
	// 初始化 ViperConfigManager
	os.Setenv("APP_NAME", "DEMO")
	os.Setenv("host", "DDD")
	//os.Setenv("host", "CCC")
	os.Setenv("PASS", "pass")
	director, _ := os.Getwd()
	vcm := &conf.ViperConfigManager{
		ConfigPath: director,
		ConfigType: "yaml",
		ConfigName: "config",
	}

	// 初始化配置, 读取配置文件
	//c, err := vcm.InitConfig(nil)
	//if err != nil {
	//	t.Errorf("InitConfig() failed, err: %v", err)
	//}
	//port := c.GetString("port")
	//t.Logf("port: %v", port)
	// 初始化配置, 读取配置文件, 并将配置映射到结构体
	configStruct := &Config{}
	cc, err := vcm.InitConfig(configStruct)
	if err != nil {
		t.Errorf("InitConfig() failed, err: %v", err)
	}
	t.Logf("configStruct.port: %v", configStruct.Port)
	t.Logf("configStruct.Host: %v", configStruct.Host)
	t.Logf("configStruct.name: %v", configStruct.Name)
	t.Logf("configStruct.path: %v", configStruct.Path)
	t.Logf("configStruct.DEMO_DB_HOST: %v", cc.GetString("HOST"))

}

type Config struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
	Name string `mapstructure:"name"`
	Path string `mapstructure:"path"`
}
