package valve_test

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"

	"github.com/FrancescoIlario/grower/internal/mocks"
	vgrpc "github.com/FrancescoIlario/grower/internal/valve/grpc"
	"github.com/FrancescoIlario/grower/internal/valve/proc"
	valvepb "github.com/FrancescoIlario/grower/pkg/valvepb/grpc"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var (
	rabbitmqConnstr string
	conn            *grpc.ClientConn
	lis             *bufconn.Listener
	cmder           mocks.Commander
)

func init() {
	rabbitmqConnstr = os.Getenv("RABBITMQ_CONNSTR")
	cmder = mocks.NewValveCmder(100 * time.Millisecond)
}

func arrange(ctx context.Context, t *testing.T) {
	publisher, err := cmdPublisher()
	if err != nil {
		t.Fatalf("error creating command publisher: %v", err)
	}

	s := grpc.NewServer()
	vlvsrv, err := vgrpc.NewGrpcServer(publisher)
	if err != nil {
		t.Fatalf("Failed to create grpc server: %v", err)
	}
	valvepb.RegisterValveServiceServer(s, vlvsrv)

	lis = bufconn.Listen(bufSize)
	conn, err = grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	go func() {
		if err := startConsumers(ctx, rabbitmqConnstr, cmder); err != nil {
			log.Fatalf("Consumers exited with error: %v", err)
		}
	}()
}

func cmdPublisher() (message.Publisher, error) {
	logger := watermill.NewStdLogger(false, false)
	commandsAMQPConfig := amqp.NewDurableQueueConfig(rabbitmqConnstr)
	commandsPublisher, err := amqp.NewPublisher(commandsAMQPConfig, logger)
	if err != nil {
		return nil, fmt.Errorf("unable to create watermill commands publisher: %w", err)
	}

	return commandsPublisher, nil
}

func startConsumers(ctx context.Context, connStr string, cmder proc.Commander) error {
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
	commandsAMQPConfig := amqp.NewDurableQueueConfig(connStr)
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
	durablePubSub := amqp.NewDurablePubSubConfig(connStr, nil)
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

func Test_OpenValve_Integration(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	arrange(ctx, t)

	client := valvepb.NewValveServiceClient(conn)
	_, err := client.OpenValve(ctx, &valvepb.OpenValveRequest{})
	if err != nil {
		t.Fatalf("error invoking endpoint: %v", err)
	}
	if exp, obt := 1, cmder.OpenInvokation(); obt != exp {
		t.Errorf("expected %v open invokation, obtained %v", exp, obt)
	}
	if exp, obt := 0, cmder.CloseInvokation(); obt != exp {
		t.Errorf("expected %v close invokation, obtained %v", exp, obt)
	}
}

func Test_CloseValve_Integration(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	arrange(ctx, t)

	client := valvepb.NewValveServiceClient(conn)
	_, err := client.CloseValve(ctx, &valvepb.CloseValveRequest{})
	if err != nil {
		t.Fatalf("error invoking endpoint: %v", err)
	}
	if exp, obt := 1, cmder.CloseInvokation(); obt != exp {
		t.Errorf("expected %v close invokation, obtained %v", exp, obt)
	}
	if exp, obt := 0, cmder.OpenInvokation(); obt != exp {
		t.Errorf("expected %v open invokation, obtained %v", exp, obt)
	}
}
