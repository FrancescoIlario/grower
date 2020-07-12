package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/FrancescoIlario/grower/pkg/valvepb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	mctx    = context.Background()
	address string
	client  valvepb.ValveServiceClient

	rootCmd = &cobra.Command{
		Use:   "vcmdrcli",
		Short: "Vcmdrcli is the CLI to manually control the valve through the gRPC server",
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&address, "host", "a", "localhost:24100", "address where server is published")
	rootCmd.AddCommand(statusCmd, openCmd, closeCmd)
}

// Execute ...
func Execute() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client = valvepb.NewValveServiceClient(conn)
	mctx = context.Background()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
