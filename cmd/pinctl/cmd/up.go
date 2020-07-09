package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stianeikeland/go-rpio/v4"
)

var highCmd = &cobra.Command{
	Use:   "high",
	Short: "set high on the provided pin",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Highering pin %d\n", pin)
		p := rpio.Pin(pin)
		p.Output()
		p.High()
		fmt.Printf("Pin %d set to high (1)\n", pin)
	},
}

func init() {
	highCmd.PersistentFlags().Uint8VarP(&pin, "pin", "p", 0, "sets the pin to govern (default 0)")
}
