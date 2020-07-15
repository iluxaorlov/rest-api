package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/iluxaorlov/rest-api/internal/app/api"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/api.toml", "Path to config file")
}

func main() {
	flag.Parse()

	config := api.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	a := api.New(config)
	if err := a.Start(); err != nil {
		log.Fatal(err)
	}
}
