package common

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

// Config for this API.
type Config struct {
	BackendDBName   string `required:"true" envconfig:"backend_db_name"`
	BackendPassword string `required:"true" envconfig:"backend_password"`
	BackendURL      string `required:"true" envconfig:"backend_url"`
	BackendUsername string `required:"true" envconfig:"backend_username"`
	Mode            string `required:"true" envconfig:"mode"`
}

// Package private variables for config synchronization.
var config *Config
var configSync sync.Once

// GetConfig get this API's configuration.
func GetConfig() *Config {
	configSync.Do(func() {
		// Extract config from environment.
		config = &Config{}
		if err := envconfig.Process("api", config); err != nil {
			log.Fatalf("Error extracting environment config: %s", err.Error())
			panic(err)
		}
	})

	return config
}
