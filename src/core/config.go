package core

import "github.com/nats-io/nats.go"

type Config struct {
	*SubscriberConfig
	*DBConfig
}

func NewConfig() *Config {
	return &Config{
		SubscriberConfig: newSubscriberConfig(),
		DBConfig:         newDbConfig(),
	}
}

type SubscriberConfig struct {
	Url        string
	ServerPort string
}

func newSubscriberConfig() *SubscriberConfig {
	return &SubscriberConfig{
		Url:        nats.DefaultURL,
		ServerPort: ":8080",
	}
}

type DBConfig struct {
	Url    string
	Driver string
}

func newDbConfig() *DBConfig {
	return &DBConfig{
		Url:    "user=postgres password=11111 dbname=test sslmode=disable",
		Driver: "postgres",
	}
}
