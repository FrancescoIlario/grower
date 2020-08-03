package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/FrancescoIlario/grower/cmd/shutter/server/conf"
	"github.com/FrancescoIlario/grower/internal/shutter"
	"github.com/FrancescoIlario/grower/pkg/shutterpb"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	c, err := conf.GetConfigurationFromEnv()
	if err != nil {
		return err
	}

	if err := startServer(c); err != nil {
		return err
	}
	return nil
}

func startServer(config conf.Configuration) error {
	logrus.Debug("Starting server")

	logrus.Debugf("acquiring conf.Address %v", config.Address)
	ls, err := net.Listen("tcp", config.Address)
	if err != nil {
		return fmt.Errorf("failed to listen at %v: %w", config.Address, err)
	}
	logrus.Debugf("acquired conf.Address %v", config.Address)

	shutterSvr := shutter.NewGrpcServer(config.OutputPin)
	grpcServer := grpc.NewServer()
	shutterpb.RegisterShutterServiceServer(grpcServer, shutterSvr)

	logrus.Debugf("starting server at %v", config.Address)
	if err := grpcServer.Serve(ls); err != nil {
		return fmt.Errorf("error serving: %w", err)
	}
	return nil
}
