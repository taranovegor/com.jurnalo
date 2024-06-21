package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var config Config

type Config struct {
	Http     Http
	Database Database
}

func init() {
	var c []byte
	var err error
	for _, configFile := range []string{
		"config.yaml",
		"/etc/jurnalo/config.yaml",
	} {
		c, err = os.ReadFile(configFile)
		if err == nil {
			break
		}
	}

	err = yaml.Unmarshal(c, &config)
	if err != nil {
		log.Fatalf("Failure in config file parsing: %v", err)
	}
}

func Get() Config {
	return config
}
