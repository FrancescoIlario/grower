package main

import (
	"os"

	"github.com/FrancescoIlario/grower/cmd/scheduler/client/cmd"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableLevelTruncation: true,
		DisableTimestamp:       false,
		ForceColors:            true,
		FullTimestamp:          true,
		DisableColors:          false,
	})

	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	cmd.Execute()
}
