package main

import (
	"net"
	"os"

	"github.com/FrancescoIlario/grower/cmd/scheduler/server/conf"
	"github.com/FrancescoIlario/grower/internal/scheduler"
	"github.com/FrancescoIlario/grower/pkg/schedulerpb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	config, err := conf.GetConfigurationFromEnv()
	if err != nil {
		logrus.Fatalf("unable to parse configuration: %v", err)
	}

	startServer(config)
}

func startServer(config conf.Configuration) {
	logrus.Debug("Starting server")

	logrus.Debugf("acquiring conf.Address %v", config.Address)
	ls, err := net.Listen("tcp", config.Address)
	if err != nil {
		logrus.Fatalf("failed to listen at %v: %v", config.Address, err)
	}
	logrus.Debugf("acquired conf.Address %v", config.Address)

	grpcServer, schedSvr := grpc.NewServer(), scheduler.NewServer()
	schedulerpb.RegisterScheduleServiceServer(grpcServer, schedSvr)

	logrus.Debugf("starting server at %v", config.Address)
	if err := grpcServer.Serve(ls); err != nil {
		logrus.Fatalf("Error serving: %v", err)
	}
}

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
