package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	SourceUrl      string `env:"SOURCE_URL"`
	RabbitUser     string `env:"RABBIT_USER"`
	RabbitPassword string `env:"RABBIT_PASSWORD"`
	RabbitHost     string `env:"RABBIT_HOST"`
	RabbitPort     string `env:"RABBIT_PORT"`
}

func Parse(s string) (*Config, error) {
	c := &Config{}
	if err := cleanenv.ReadConfig(s, c); err != nil {
		return nil, err
	}

	return c, nil
}
