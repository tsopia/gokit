package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// ConfigManager 是一个管理配置的接口
type ConfigManager interface {
	Init(viperConfig interface{}) *viper.Viper
}

// ViperConfigManager 是 ConfigManager 接口的具体实现
type ViperConfigManager struct {
	configPath string
	configName string
	configType string
}

// Init 初始化 Viper，如果提供了结构体，将配置映射到该结构体；如果没有提供结构体，直接返回 Viper 对象
func (vc *ViperConfigManager) Init(viperConfig interface{}) (*viper.Viper, error) {
	// 初始化 Viper
	viper.AddConfigPath(vc.configPath)
	viper.SetConfigName(vc.configType)
	viper.SetConfigType(vc.configType)
	viper.AutomaticEnv()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		// 如果读取配置文件失败，记录错误并返回 Viper 对象
		return nil, fmt.Errorf("Error reading config file: %w", err)
		//return viper.GetViper()
	}

	// 如果提供了结构体，尝试将其映射到配置
	if viperConfig != nil {
		if err := viper.Unmarshal(viperConfig); err != nil {
			// 如果映射配置到结构体失败，记录错误并返回 Viper 对象
			return nil, fmt.Errorf("Error unmarshaling config to struct: %w", err)

		}
	}

	// 返回 Viper 对象
	return viper.GetViper(), nil
}
