package global

import (
	"ginapi/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Dbs         map[string]*gorm.DB
}

var App = new(Application)
