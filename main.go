package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/obegendi/go-ddd/api"
	cfg "github.com/obegendi/go-ddd/config"
)

const (
	defaultEnv = "local"
)

func main() {
	flag.Parse()

	env := getEnvironment("APP_ENVIRONMENT", defaultEnv)
	fmt.Println(env)
	config := cfg.Config{}
	config.Read(env)
	api.Init(&config)
}

func getEnvironment(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
