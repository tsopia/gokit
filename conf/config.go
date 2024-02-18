package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

// ConfigManager 是一个管理配置的接口
type ConfigManager interface {
	InitConfig(viperConfig interface{}) *viper.Viper
}

// ViperConfigManager 是 ConfigManager 接口的具体实现
type ViperConfigManager struct {
	ConfigPath string
	ConfigName string
	ConfigType string
}

// InitConfig 初始化 Viper，如果提供了结构体，将配置映射到该结构体；如果没有提供结构体，直接返回 Viper 对象
func (vc *ViperConfigManager) InitConfig(viperConfig interface{}) (*viper.Viper, error) {
	// 初始化 Viper
	viper.AddConfigPath(vc.ConfigPath)
	viper.SetConfigName(vc.ConfigName)
	//viper.SetConfigType(vc.ConfigType)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		// 如果读取配置文件失败，记录错误并返回 Viper 对象
		return nil, fmt.Errorf("Error reading conf file: %w", err)
		//return viper.GetViper()
	}

	// 设置环境变量前缀
	//viper.SetEnvPrefix("DEMO")
	//viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// 如果提供了结构体，尝试将其映射到配置

	if viperConfig != nil {
		if err := viper.Unmarshal(viperConfig); err != nil {
			// 如果映射配置到结构体失败，记录错误并返回 Viper 对象
			return nil, fmt.Errorf("Error unmarshaling conf to struct: %w", err)

		}
	}
	fmt.Println(viper.AllSettings())

	return viper.GetViper(), nil
}
