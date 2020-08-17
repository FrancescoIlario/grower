package main

import (
	"context"
	"fmt"
	"os"

	"github.com/FrancescoIlario/grower/cmd/valvecmdsproc/conf"
	"github.com/FrancescoIlario/grower/internal/valve/proc"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
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
		logrus.Fatal(err)
	}
}

func run() error {
	logrus.Debug("Starting server")

	c, err := conf.GetConfigurationFromEnv()
	if err != nil {
		return fmt.Errorf("unable to parse configuration: %w", err)
	}

	rpio.Open()
	defer rpio.Close()

	pp, np := rpio.Pin(c.PositivePin), rpio.Pin(c.NegativePin)
	cmder := proc.NewCommander(pp, np, c.PulseLength)
	logrus.Debugf("valve cmder initialized with ppin %d, npin %d, and duration %v", pp, np, c.PulseLength)

	ctx := context.Background()
	if err := startConsumers(ctx, c, cmder); err != nil {
		return err
	}
	return nil
}

func startConsumers(ctx context.Context, c conf.Configuration, cmder proc.Commander) error {
	logger := watermill.NewStdLogger(false, false)
	cqrsMarshaler := cqrs.ProtobufMarshaler{}

	// CQRS is built on messages router. Detailed documentation: https://watermill.io/docs/messages-router/
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return err
	}

	// You can use any Pub/Sub implementation from here: https://watermill.io/docs/pub-sub-implementations/
	// Detailed RabbitMQ implementation: https://watermill.io/docs/pub-sub-implementations/#rabbitmq-amqp
	// Commands will be send to queue, because they need to be consumed once.
	commandsAMQPConfig := amqp.NewDurableQueueConfig(c.AmqpConnectionString)
	commandsPublisher, err := amqp.NewPublisher(commandsAMQPConfig, logger)
	if err != nil {
		return err
	}
	commandsSubscriber, err := amqp.NewSubscriber(commandsAMQPConfig, logger)
	if err != nil {
		return err
	}

	// Events will be published to PubSub configured Rabbit, because they may be consumed by multiple consumers.
	// (in that case BookingsFinancialReport and OrderBeerOnRoomBooked).
	durablePubSub := amqp.NewDurablePubSubConfig(c.AmqpConnectionString, nil)
	eventsPublisher, err := amqp.NewPublisher(durablePubSub, logger)
	if err != nil {
		return err
	}

	// Simple middleware which will recover panics from event or command handlers.
	// More about router middlewares you can find in the documentation:
	// https://watermill.io/docs/messages-router/#middleware
	//
	// List of available middlewares you can find in message/router/middleware.
	router.AddMiddleware(middleware.Recoverer)

	// cqrs.Facade is facade for Command and Event buses and processors.
	// You can use facade, or create buses and processors manually (you can inspire with cqrs.NewFacade)
	if _, err := cqrs.NewFacade(cqrs.FacadeConfig{
		GenerateCommandsTopic: func(commandName string) string {
			// we are using queue RabbitMQ config, so we need to have topic per command type
			return commandName
		},
		CommandHandlers: func(cb *cqrs.CommandBus, eb *cqrs.EventBus) []cqrs.CommandHandler {
			return []cqrs.CommandHandler{
				proc.NewOpenHandler(eb, cmder),
				proc.NewCloseHandler(eb, cmder),
			}
		},
		CommandsPublisher: commandsPublisher,
		CommandsSubscriberConstructor: func(handlerName string) (message.Subscriber, error) {
			// we can reuse subscriber, because all commands have separated topics
			return commandsSubscriber, nil
		},
		GenerateEventsTopic: func(eventName string) string {
			// because we are using PubSub RabbitMQ config, we can use one topic for all events
			return "events"

			// we can also use topic per event type
			// return eventName
		},
		EventsPublisher:       eventsPublisher,
		Router:                router,
		CommandEventMarshaler: cqrsMarshaler,
		Logger:                logger,
	}); err != nil {
		return err
	}

	// processors are based on router, so they will work when router will start
	if err := router.Run(context.Background()); err != nil {
		return err
	}

	return nil
}
