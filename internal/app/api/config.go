package api

import "github.com/iluxaorlov/rest-api/internal/app/store"

type Config struct {
	Address  string `toml:"address"`
	LogLevel string `toml:"log-level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		Address:  ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
