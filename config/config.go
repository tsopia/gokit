package config

import (
	"github.com/spf13/viper"
)

var config *viper.Viper

func init() {
	config = viper.New()
	// 可以在这里设置Viper的默认配置
	// 例如：config.SetDefault("key", "value")
}

// LoadConfig 用于加载配置文件
func LoadConfig(filePath string) error {
	config.SetConfigFile(filePath)
	return config.ReadInConfig()
}

// GetString 获取配置文件中的字符串配置
func GetString(key string) string {
	return config.GetString(key)
}

// GetInt 获取配置文件中的整数配置
func GetInt(key string) int {
	return config.GetInt(key)
}

// GetBool 获取配置文件中的布尔配置
func GetBool(key string) bool {
	return config.GetBool(key)
}

// 添加其他你需要的配置获取方法...
