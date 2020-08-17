package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/FrancescoIlario/grower/pkg/valvepb/grpc"
	"github.com/FrancescoIlario/grower/pkg/valvepb/shared"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Retrieve the status of the valve",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(mctx, 10*time.Second)
		defer cancel()

		rs, err := client.GetStatus(ctx, &grpc.GetStatusRequest{})
		if err != nil {
			logrus.Fatalf("Could not open the valve: %v", err)
		}

		fmt.Println("Valve opened")
		st := rs.Status
		switch st {
		case shared.ValveStatus_VALVE_INVALID:
			fmt.Printf("The valve is in an invalid status")
		case shared.ValveStatus_VALVE_OPEN:
			fmt.Printf("The valve is open")
		case shared.ValveStatus_VALVE_OPENING:
			fmt.Printf("The valve is opening")
		case shared.ValveStatus_VALVE_CLOSE:
			fmt.Printf("The valve is close")
		case shared.ValveStatus_VALVE_CLOSING:
			fmt.Printf("The valve is closing")
		case shared.ValveStatus_VALVE_UNSPECIFIED:
			logrus.Fatal("The server respondend with UNSPECIFIED STATUS")
		}
	},
}
