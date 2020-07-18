package conf

import (
	"github.com/kelseyhightower/envconfig"
)

// EnvPrefix prefix for environmental variables parsed by application
const EnvPrefix = "SCHEDULER"

// Configuration Scheduler service configuration
type Configuration struct {
	Address       string `default:"localhost:24101"`
	ValveCmdrHost string `default:"localhost:24100"`
}

// GetConfigurationFromEnv ...
func GetConfigurationFromEnv() (s Configuration, err error) {
	err = envconfig.Process(EnvPrefix, &s)
	return
}
