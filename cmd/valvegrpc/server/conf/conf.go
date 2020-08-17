package conf

import (
	"github.com/kelseyhightower/envconfig"
)

// EnvPrefix prefix for environmental variables parsed by application
const EnvPrefix = "VALVEGRPC"

// Configuration ValveCmdr service configuration
type Configuration struct {
	Address              string `default:"localhost:24100"`
	AmqpConnectionString string `default:"amqp://guest:guest@rabbitmq:5672/"`
}

// GetConfigurationFromEnv ...
func GetConfigurationFromEnv() (s Configuration, err error) {
	err = envconfig.Process(EnvPrefix, &s)
	return
}
