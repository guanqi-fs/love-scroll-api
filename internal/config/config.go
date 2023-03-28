package config

import (
	"github.com/spf13/viper"
	"sync"
)

type Config struct {
	Server struct {
		Address string
	}

	Database struct {
		Driver   string
		DSN      string
	}

	JWT struct {
		Secret string
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")

		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}

		err = viper.Unmarshal(instance)
		if err != nil {
			panic(err)
		}
	})

	return instance
}
