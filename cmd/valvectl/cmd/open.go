package cmd

import (
	"fmt"

	"github.com/FrancescoIlario/grower/internal/valve/proc"
	"github.com/spf13/cobra"
	"github.com/stianeikeland/go-rpio/v4"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens the valve",
	Run: func(cmd *cobra.Command, args []string) {
		pp, np := rpio.Pin(positiveRelayPin), rpio.Pin(negativeRelayPin)
		cmder := proc.NewCommander(pp, np, pulseLength)

		cmder.Open()
		fmt.Println("Valve opened")
	},
}
