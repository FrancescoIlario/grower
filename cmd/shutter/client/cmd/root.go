package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/FrancescoIlario/grower/pkg/shutterpb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	mctx    = context.Background()
	address string
	client  shutterpb.ShutterServiceClient
	id      string

	rootCmd = &cobra.Command{
		Use:   "shuttercli",
		Short: "shuttercli is the CLI to manually control the shutter through the gRPC server",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&address, "host", "a", "localhost:24102", "address where server is published")
	rootCmd.AddCommand(shutCmd)
}

// Execute ...
func Execute() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client = shutterpb.NewShutterServiceClient(conn)
	mctx = context.Background()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
