package cmd

import (
	"fmt"

	"github.com/FrancescoIlario/grower/internal/valve"
	"github.com/spf13/cobra"
	"github.com/stianeikeland/go-rpio/v4"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Retrieve the status of the valve",
	Run: func(cmd *cobra.Command, args []string) {
		pp, np := rpio.Pin(positiveRelayPin), rpio.Pin(negativeRelayPin)
		cmder := valve.NewCommander(pp, np)

		st := cmder.Status()
		switch st {
		case valve.StatusInvalid:
			fmt.Printf("The valve is in an invalid status")
		case valve.StatusOpen:
			fmt.Printf("The valve is open")
		case valve.StatusOpening:
			fmt.Printf("The valve is opening")
		case valve.StatusClose:
			fmt.Printf("The valve is close")
		case valve.StatusClosing:
			fmt.Printf("The valve is closing")
		}
	},
}
