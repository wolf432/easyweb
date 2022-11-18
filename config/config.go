package config

type ServerConfig struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
	App   App   `mapstructure:"app" json:"app" yaml:"app"`
	Jwt   Jwt   `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
