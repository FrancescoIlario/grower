package main

import (
	"fmt"
	"net"
	"os"

	"github.com/FrancescoIlario/grower/cmd/valvegrpc/server/conf"
	vgrpc "github.com/FrancescoIlario/grower/internal/valve/grpc"
	valvepb "github.com/FrancescoIlario/grower/pkg/valvepb/grpc"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"

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

	commandsPublisher, err := cmdPublisher(c)
	if err != nil {
		logrus.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	valveServer, err := vgrpc.NewGrpcServer(commandsPublisher)
	if err != nil {
		logrus.Fatalf("Error initializing server: %v", err)
	}
	valvepb.RegisterValveServiceServer(grpcServer, valveServer)

	logrus.Debugf("starting server at %v", c.Address)
	if err := grpcServer.Serve(ls); err != nil {
		logrus.Fatalf("Error serving: %v", err)
	}
}

func cmdPublisher(c conf.Configuration) (message.Publisher, error) {
	logger := watermill.NewStdLogger(false, false)
	commandsAMQPConfig := amqp.NewDurableQueueConfig(c.AmqpConnectionString)
	commandsPublisher, err := amqp.NewPublisher(commandsAMQPConfig, logger)
	if err != nil {
		return nil, fmt.Errorf("unable to create watermill commands publisher: %w", err)
	}

	return commandsPublisher, nil
}
