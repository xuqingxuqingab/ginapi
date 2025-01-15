package config

// 配置文件映射结构体，对应config.yaml下的app节点，命名请与配置文件中保持一致，方便理解
type App struct {
	Env string `mapstructure:"env" json:"env" yaml:"env"`
    Port string `mapstructure:"port" json:"port" yaml:"port"`
    AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
}