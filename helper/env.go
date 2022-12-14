package helper

import (
	"github.com/caarlos0/env"
	_ "github.com/joho/godotenv/autoload"
)

type config struct {
	APP_PASSWORD string `env:"APP_PASSWORD"`
	APP_EMAIL    string `env:"APP_EMAIL"`
}

func LoadConfig() *config {
	cfg := new(config)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
