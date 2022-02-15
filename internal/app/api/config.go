package api

import "StandartWebServer/store"

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	LoggerLevel string `toml:"logger_level"`
	Store       *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ": 8080",
		LoggerLevel: "debug",
		Store:       store.NewConfig(),
	}
}
