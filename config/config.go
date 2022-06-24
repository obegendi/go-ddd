package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Server   string
	Database string
}

func (c *Config) Read(env string) {
	if _, err := toml.DecodeFile("config."+env+".toml", &c); err != nil {
		log.Println("t")
		log.Fatal(err)
	}
}
