package config

import "github.com/nats-io/nats.go"

type (
	Config struct {
		ServerConfig
		SubscriberConfig
		DBConfig
	}

	ServerConfig struct {
		Port string
	}

	SubscriberConfig struct {
		Url     string
		Subject string
	}

	DBConfig struct {
		Url    string
		Driver string
	}
)

func GetConfig() *Config {
	return &Config{
		ServerConfig:     getServerConfig(),
		SubscriberConfig: getSubscriberConfig(),
		DBConfig:         getDbConfig(),
	}
}

func getServerConfig() ServerConfig {
	return ServerConfig{
		Port: ":8080",
	}
}

func getSubscriberConfig() SubscriberConfig {
	return SubscriberConfig{
		Url:     nats.DefaultURL,
		Subject: "foo",
	}
}

func getDbConfig() DBConfig {
	return DBConfig{
		Url:    "user=postgres password=11111 dbname=test sslmode=disable",
		Driver: "postgres",
	}
}
