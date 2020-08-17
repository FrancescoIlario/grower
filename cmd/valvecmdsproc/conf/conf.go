package conf

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

// EnvPrefix prefix for environmental variables parsed by application
const EnvPrefix = "VCMDSVR"

// Configuration ValveCmdr service configuration
type Configuration struct {
	PositivePin          uint8         `default:"10" split_words:"true"`
	NegativePin          uint8         `default:"4" split_words:"true"`
	PulseLength          time.Duration `default:"20ms" split_words:"true"`
	AmqpConnectionString string        `default:"amqp://guest:guest@rabbitmq:5672/"`
}

// GetConfigurationFromEnv ...
func GetConfigurationFromEnv() (s Configuration, err error) {
	err = envconfig.Process(EnvPrefix, &s)
	return
}
