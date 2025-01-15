package bootstrap

import (
	"fmt"
	"ginapi/app/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)
func InitializeConfig() *viper.Viper {
	// 初始化配置文件
	println("初始化配置文件")	
	config := "config/config.yaml"

	// 加载配置文件
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")

	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("read config failed: %s", err))
	}
	fmt.Println("config:", v.Get("app"))


	// 监听配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)

		// 发生变化后重新映射结构体
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})

	// 将配置赋值给全局变量
    if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}
	return v
}