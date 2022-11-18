package initialize

import (
	"easyweb/global"
	"github.com/spf13/viper"
	"os"
)

// InitViper 初始化配置
func InitViper() *viper.Viper {
	config := "config.yaml"
	//如果配置的环境变量则从环境变量里读取
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		panic("配置文件读取失败")
	}

	//配置文件解析到结构体中
	if err := v.Unmarshal(&global.Cfg); err != nil {
		panic("初始化配置失败")
	}
	return v
}
