package config

import (
	"github.com/alonelegion/notes_system/api_service/app/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	isDebug *bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type" env-default="port"`
		BindIP string `yaml:"bind_ip" env-default="localhost"`
		Port   string `yaml:"port" env-default="8080"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(&instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
