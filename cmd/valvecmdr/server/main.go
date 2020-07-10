package main

import (
	"net"
	"os"

	"github.com/FrancescoIlario/grower/cmd/valvecmdr/server/conf"
	"github.com/FrancescoIlario/grower/internal/valve"
	"github.com/FrancescoIlario/grower/pkg/valvepb"
	"github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
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
	logrus.Debug("Starting server")

	c, err := conf.GetConfigurationFromEnv()
	if err != nil {
		logrus.Fatalf("unable to parse configuration: %v", err)
	}

	logrus.Debugf("acquiring conf.Address %v", c.Address)
	ls, err := net.Listen("tcp", c.Address)
	if err != nil {
		logrus.Fatalf("failed to listen at %v: %v", c.Address, err)
	}
	logrus.Debugf("acquired conf.Address %v", c.Address)

	rpio.Open()
	defer rpio.Close()

	pp, np := rpio.Pin(c.PositivePin), rpio.Pin(c.NegativePin)
	cmder := valve.NewCommander(pp, np, c.PulseLength)
	logrus.Debugf("valve cmder initialized with ppin %d, npin %d, and duration %v", pp, np, c.PulseLength)

	grpcServer := grpc.NewServer()
	valveServer := valve.NewGrpcServer(cmder)
	valvepb.RegisterValveServiceServer(grpcServer, valveServer)

	logrus.Debugf("starting server at %v", c.Address)
	if err := grpcServer.Serve(ls); err != nil {
		logrus.Fatalf("Error serving: %v", err)
	}
}
