package config

import "github.com/nats-io/nats.go"

type (
	Config struct {
		SubscriberConfig
		DBConfig
	}

	SubscriberConfig struct {
		Url        string
		ServerPort string
	}

	DBConfig struct {
		Url    string
		Driver string
	}
)

func GetConfig() *Config {
	return &Config{
		SubscriberConfig: getSubscriberConfig(),
		DBConfig:         getDbConfig(),
	}
}

func getSubscriberConfig() SubscriberConfig {
	return SubscriberConfig{
		Url:        nats.DefaultURL,
		ServerPort: ":8080",
	}
}

func getDbConfig() DBConfig {
	return DBConfig{
		Url:    "user=postgres password=11111 dbname=test sslmode=disable",
		Driver: "postgres",
	}
}
