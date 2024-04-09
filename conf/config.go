package conf

import (
	"github.com/spf13/viper"
	"github.com/tsopia/gokit/model"
	"github.com/tsopia/gokit/xerrors"
	"log"
	"os"
	"reflect"
	"strings"
)

var DefaultConf = &model.ManagerConf{}

func init() {
	director, _ := os.Getwd()
	vcm := &ViperConfigManager{
		ConfigPath: director,
		ConfigType: "yaml",
		ConfigName: "config",
	}
	_ = vcm.InitConfig(DefaultConf)

}

// ViperConfigManager 是基于viper实现的配置管理器，负责初始化和管理配置文件
type ViperConfigManager struct {
	ConfigPath string // 配置文件路径
	ConfigName string // 配置文件名
	ConfigType string // 配置文件类型（如：json, yaml等）
}

// InitConfig 初始化viper配置并映射到指定的结构体上。如果未指定结构体，则仅初始化viper。
// 参数conf: 可选，配置数据将映射到此结构体上。如果为空，则不执行映射。
// 返回值: 返回可能遇到的错误，如果初始化成功则返回nil。
func (vc *ViperConfigManager) InitConfig(conf interface{}) error {
	var err *xerrors.Error
	defer func() {
		if err != nil {
			log.Printf("Failed to initialize config: %v\n", err)
		}

	}()
	// 初始化viper配置路径、文件名和类型，并设置环境变量替换规则
	viper.AddConfigPath(vc.ConfigPath)
	viper.SetConfigName(vc.ConfigName)
	viper.SetConfigType(vc.ConfigType)
	// 设置环境变量键的替换器，将点(.)替换为下划线(_)。
	// 这使得Viper可以正确读取以点分隔的环境变量名，例如"APP_NAME"将被识别为"app.name"。
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 自动加载环境变量。
	// 该功能使得Viper可以自动将环境变量作为配置项，无需显式设置。
	viper.AutomaticEnv()

	// 尝试读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		err = xerrors.Wrap(err, xerrors.ErrFileNotFound, "Failed to read config file")
	}
	// 如果提供了结构体，将配置映射到该结构体上
	if conf != nil {
		err = bindEnv(err, conf)
	}
	return err
}

func bindEnv(err error, conf interface{}) *xerrors.Error {
	to := reflect.TypeOf(conf).Elem()
	for i := 0; i < to.NumField(); i++ {
		field := to.Field(i)
		// 将环境变量绑定到结构体字段
		err = viper.BindEnv(strings.ToUpper(field.Name))
		if err != nil {
			return xerrors.Wrap(err, xerrors.ErrEnvBinding, "Failed to bind env variable")
		}
	}

	// 将配置数据解析到结构体中
	if err = viper.Unmarshal(conf); err != nil {
		return xerrors.Wrap(err, xerrors.ErrUnmarshal, "Failed to unmarshal config to struct")
	}
	return nil
}
