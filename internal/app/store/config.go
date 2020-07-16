package store

type Config struct {
	Url string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
