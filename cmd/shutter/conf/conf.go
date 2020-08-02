package conf

import (
	"github.com/kelseyhightower/envconfig"
)

// EnvPrefix prefix for environmental variables parsed by application
const EnvPrefix = "SHUTTER"

// Configuration service configuration
type Configuration struct {
	InputPin  uint `default:"17"`
	OutputPin uint `default:"27"`
}

// GetConfigurationFromEnv ...
func GetConfigurationFromEnv() (s Configuration, err error) {
	err = envconfig.Process(EnvPrefix, &s)
	return
}
