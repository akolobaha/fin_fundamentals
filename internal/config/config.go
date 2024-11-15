package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	SourceUrl      string `env:"SOURCE_URL"`
	RabbitUser     string `env:"RABBIT_USER"`
	RabbitPassword string `env:"RABBIT_PASSWORD"`
	RabbitHost     string `env:"RABBIT_HOST"`
	RabbitPort     string `env:"RABBIT_PORT"`
	RabbitQueue    string `env:"RABBIT_QUEUE"`
	RabbitDsn      string
}

func Parse(s string) (*Config, error) {
	c := &Config{}
	if err := cleanenv.ReadConfig(s, c); err != nil {
		return nil, err
	}

	c.RabbitDsn = InitRabbitDSN(c)

	return c, nil
}

func InitRabbitDSN(c *Config) string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%s/", c.RabbitUser, c.RabbitPassword, c.RabbitHost, c.RabbitPort,
	)
}
