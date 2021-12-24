package core

import "github.com/nats-io/nats.go"

type Config struct {
	*SubscriberConfig
	*DBConfig
}

func GetConfig() *Config {
	return &Config{
		SubscriberConfig: getSubscriberConfig(),
		DBConfig:         getDbConfig(),
	}
}

type SubscriberConfig struct {
	Url        string
	ServerPort string
}

func getSubscriberConfig() *SubscriberConfig {
	return &SubscriberConfig{
		Url:        nats.DefaultURL,
		ServerPort: ":8080",
	}
}

type DBConfig struct {
	Url    string
	Driver string
}

func getDbConfig() *DBConfig {
	return &DBConfig{
		Url:    "user=postgres password=11111 dbname=test sslmode=disable",
		Driver: "postgres",
	}
}
