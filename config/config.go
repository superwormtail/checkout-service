package config

import (
	"fmt"
	"log"
	"os"
)

// Env Global Configuration
var Env *tConfig
var strError string

type tConfig struct {
	App appConfig
}

type appConfig struct {
	AppName           string
	HTTPPort          string
	HTTPServerTimeOut string
}

func getEnv(ENV string) string {
	val := os.Getenv(ENV)
	if val == "" {
		if strError != "" {
			strError = fmt.Sprintf("%s,", strError)
		}
		strError = fmt.Sprintf("%s%s", strError, ENV)
	}
	return val
}

// Initialization :
func Initialization() error {

	conf := tConfig{
		appConfig{
			AppName:           getEnv("AppName"),
			HTTPPort:          getEnv("HTTPPort"),
			HTTPServerTimeOut: getEnv("HTTPServerTimeOut"),
		},
	}

	Env = &conf

	// validate
	if strError != "" {
		log.Fatalf("Exit: environment variable must be set: %s.", strError)
		os.Exit(1)
	}

	return nil
}
