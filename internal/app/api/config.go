package api

type Config struct {
	Address  string `toml:"address"`
	LogLevel string `toml:"log-level"`
}

func NewConfig() *Config {
	return &Config{
		Address: ":8080",
		LogLevel: "debug",
	}
}
