package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DBDriver string `env:"DB_DRIVER" env-default:"postgres"`
	DBSource string `env:"DB_SOURCE"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		cfg := Config{}
		if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
			log.Fatal("cannot read the config", err)
		}
		instance = &cfg
	})

	if instance == nil {
		log.Fatal("failed to initialize config")
	}
	return instance
}
