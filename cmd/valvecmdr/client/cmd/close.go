package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/FrancescoIlario/grower/pkg/valvepb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Closes the valve",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(mctx, 10*time.Second)
		defer cancel()

		if _, err := client.CloseValve(ctx, &valvepb.CloseValveRequest{}); err != nil {
			logrus.Fatalf("Could not close the valve: %v", err)
		}

		fmt.Println("Valve closed")
	},
}
