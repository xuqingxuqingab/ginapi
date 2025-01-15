package config

// 配置文件主体映射结构体---config.yaml的第一层
type Configuration struct {
	App App `mapstructure:"app" json:"app" yaml:"app"`
	Db  Db  `mapstructure:"db" json:"db" yaml:"db"`
}