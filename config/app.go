package config

type App struct {
	Env  string `mapstructure:"env" json:"env" yaml:"env"`    // 开发环境 dev prod
	Port string `mapstructure:"port" json:"port" yaml:"port"` // 运行端口
}
