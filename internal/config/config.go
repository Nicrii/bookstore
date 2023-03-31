package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

func GetConfig() Config {
	var c Config
	err := envconfig.Process("DB", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	return c
}
