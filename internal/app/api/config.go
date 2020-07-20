package api

type Config struct {
	Address     string `toml:"address"`
	LogLevel    string `toml:"log-level"`
	DatabaseUrl string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

func NewConfig() *Config {
	return &Config{
		Address:  ":8080",
		LogLevel: "debug",
	}
}
