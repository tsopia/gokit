package example

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/tsopia/gokit/conf"
	"os"
	"reflect"
	"strings"

	"testing"
)

type Config struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
	Name string `mapstructure:"name"`
	Path string `mapstructure:"path"`
}

func TestInitConfig(t *testing.T) {
	// 初始化 ViperConfigManager
	os.Setenv("PORT", "8080")
	os.Setenv("HOST", "DDD")
	os.Setenv("PATH", "pass")
	director, _ := os.Getwd()
	vcm := &conf.ViperConfigManager{
		ConfigPath: director,
		ConfigType: "yaml",
		ConfigName: "config",
	}
	configStruct := &Config{}
	err := vcm.InitConfig(configStruct)
	if err != nil {
		t.Errorf("InitConfig() failed, err: %v", err)
	}
	t.Logf("configStruct.port: %v", configStruct.Port)
	t.Logf("configStruct.Host: %v", configStruct.Host)
	t.Logf("configStruct.name: %v", configStruct.Name)
	t.Logf("configStruct.path: %v", configStruct.Path)
	t.Logf("configStruct.pass: %v", viper.GetString("pass"))

}

type AppConfig struct {
	DatabaseURL string `mapstructure:"DB_URL"`
	ServerPort  int    `mapstructure:"SERVER_PORT"`
	DebugMode   bool   `mapstructure:"DEBUG_MODE"`
}

func bindEnvForStruct(v *viper.Viper, prefix string, s interface{}) {
	val := reflect.ValueOf(s).Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("mapstructure")
		envKey := fmt.Sprintf("%s%s", prefix, tag)

		if envVal := v.Get(envKey); envVal != nil {
			v.BindEnv(envKey)
		}
	}
}
func TestDemo(t *testing.T) {
	os.Setenv("DATABASE_URL", "localhost:5432")
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("DEBUG_MODE", "true")
	os.Setenv("NAME", "APP")
	// 设置 Viper，使其自动读取环境变量
	// 将环境变量与配置文件中具有相同前缀的键关联起来
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	//viper.MustBindEnv(viper.GetString("APP_NAME"))
	viper.AutomaticEnv()
	t.Logf(viper.GetString("APP_NAME"))
	//viper.SetEnvPrefix("APP")
	// 创建配置结构体
	var config AppConfig
	// 打印配置信息
	fmt.Printf(os.Getenv("DB_URL"))
	fmt.Printf(os.Getenv("SERVER_PORT"))
	fmt.Printf(os.Getenv("DEBUG_MODE"))

	// 手动为结构体字段绑定环境变量
	// 将环境变量映射到结构体
	to := reflect.TypeOf(config)
	for i := 0; i < to.NumField(); i++ {
		field := to.Field(i)

		tag := field.Tag.Get("mapstructure")
		t.Logf(strings.ToUpper(tag))
		viper.BindEnv(strings.ToUpper(tag))
	}
	//viper.BindEnv("DB_URL")
	//viper.BindEnv("SERVER_PORT")
	//viper.BindEnv("DEBUG_MODE")
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Failed to unmarshal config: %s", err)
		return
	}
	// 打印配置信息
	fmt.Printf("Database URL: %s\n", config.DatabaseURL)
	fmt.Printf("Server Port: %d\n", config.ServerPort)
	fmt.Printf("Debug Mode: %t\n", config.DebugMode)
}
