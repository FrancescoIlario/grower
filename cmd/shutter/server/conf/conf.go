package conf

import (
	"github.com/kelseyhightower/envconfig"
)

// EnvPrefix prefix for environmental variables parsed by application
const EnvPrefix = "SHUTTER"

// Configuration service configuration
type Configuration struct {
	Address   string `default:"localhost:24102"`
	OutputPin uint   `default:"27"`
}

// GetConfigurationFromEnv ...
func GetConfigurationFromEnv() (s Configuration, err error) {
	err = envconfig.Process(EnvPrefix, &s)
	return
}
