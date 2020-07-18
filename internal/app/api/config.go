package api

type Config struct {
	Address     string `toml:"address"`
	LogLevel    string `toml:"log-level"`
	DatabaseUrl string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		Address:     ":8080",
		LogLevel:    "debug",
	}
}
