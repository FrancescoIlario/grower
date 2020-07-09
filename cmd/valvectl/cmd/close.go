package cmd

import (
	"fmt"

	"github.com/FrancescoIlario/grower/internal/valve"
	"github.com/spf13/cobra"
	"github.com/stianeikeland/go-rpio/v4"
)

var closeCmd = &cobra.Command{
	Use:   "close",
	Short: "Closes the valve",
	Run: func(cmd *cobra.Command, args []string) {
		pp, np := rpio.Pin(positiveRelayPin), rpio.Pin(negativeRelayPin)
		cmder := valve.NewCommander(pp, np)

		cmder.Close()
		fmt.Println("Valve closed")
	},
}
