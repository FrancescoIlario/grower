package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/FrancescoIlario/grower/pkg/valvepb/grpc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens the valve",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(mctx, 10*time.Second)
		defer cancel()

		if _, err := client.OpenValve(ctx, &grpc.OpenValveRequest{}); err != nil {
			logrus.Fatalf("Could not open the valve: %v", err)
		}

		fmt.Println("Valve opened")
	},
}
