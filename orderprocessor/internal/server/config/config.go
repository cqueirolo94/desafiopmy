package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Subject       string `default:"LOMA"`
	HttpPort      string `split_words:"true" default:"8080"`
	NatsUrl       string `split_words:"true" default:"nats://localhost:4222"`
	MaxGoroutines int    `split_words:"true" default:"1040000"`
}

func NewConfig() *Config {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		panic(err)
	}

	return &c
}
