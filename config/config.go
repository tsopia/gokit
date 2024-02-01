package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var config *viper.Viper

func init() {
	config = viper.New()
	// 可以在这里设置Viper的默认配置
	// 例如：config.SetDefault("key", "value")
}
func InitializeConfig() *viper.Viper {

	// 设置配置文件路径

	config := "config.yaml"

	// 生产环境可以通过设置环境变量来改变配置文件路径
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	// 初始化 viper
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s \n", err))
	}

	return v
}
func WriteConfig() error {
	// 监听配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		// 重载配置
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})
	// 将配置赋值给全局变量
	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}
	return config.WriteConfig()
}
