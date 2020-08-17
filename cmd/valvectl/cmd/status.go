package cmd

import (
	"fmt"

	"github.com/FrancescoIlario/grower/internal/valve/proc"
	"github.com/spf13/cobra"
	"github.com/stianeikeland/go-rpio/v4"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Retrieve the status of the valve",
	Run: func(cmd *cobra.Command, args []string) {
		pp, np := rpio.Pin(positiveRelayPin), rpio.Pin(negativeRelayPin)
		cmder := proc.NewCommander(pp, np, pulseLength)

		st := cmder.Status()
		switch st {
		case proc.StatusInvalid:
			fmt.Printf("The valve is in an invalid status")
		case proc.StatusOpen:
			fmt.Printf("The valve is open")
		case proc.StatusOpening:
			fmt.Printf("The valve is opening")
		case proc.StatusClose:
			fmt.Printf("The valve is close")
		case proc.StatusClosing:
			fmt.Printf("The valve is closing")
		}
	},
}
