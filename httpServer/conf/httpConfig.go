package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"myEntry/httpServer/common"
)

// Config global config
type Config struct {
	// common
	Http C
	// component config
	TcpPool common.TcpConfig
}

// C AppConfig app config
type C struct {
	Name    string
	Addr    string
	Host    string
	PicPath string
}

var (
	// Conf  app global config
	Conf = &Config{}
)

func Init(configPath string) *Config {
	viper.SetConfigType("yml")
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&Conf); err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&Conf); err != nil {
			panic(err)
		}
	})
	return Conf
}
