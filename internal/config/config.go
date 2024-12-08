package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
	"time"
)

type Config struct {
	SourceUrl      string      `env:"SOURCE_URL"`
	RabbitUser     string      `env:"RABBIT_USER"`
	RabbitPassword string      `env:"RABBIT_PASSWORD"`
	RabbitHost     string      `env:"RABBIT_HOST"`
	RabbitPort     string      `env:"RABBIT_PORT"`
	RabbitQueue    string      `env:"RABBIT_QUEUE"`
	LogLevel       string      `env:"LOG_LEVEL"`
	TickInterval   time.Ticker `env:"TICK_INTERVAL"`
}

func Parse(s string) (*Config, error) {
	c := &Config{}
	if err := cleanenv.ReadConfig(s, c); err != nil {
		return nil, err
	}
	setLogLevel(c.LogLevel)
	return c, nil
}

func setLogLevel(level string) {
	switch level {
	case "debug":
		slog.SetLogLoggerLevel(-4)
	case "info":
		slog.SetLogLoggerLevel(0)
	case "warn":
		slog.SetLogLoggerLevel(4)
	case "error":
		slog.SetLogLoggerLevel(8)
	default:
		slog.SetLogLoggerLevel(4)
	}
}

func (c *Config) GetRabbitDSN() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%s/", c.RabbitUser, c.RabbitPassword, c.RabbitHost, c.RabbitPort,
	)
}
