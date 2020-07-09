package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stianeikeland/go-rpio/v4"
)

var lowCmd = &cobra.Command{
	Use:   "low",
	Short: "set 0 on pin",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Lowering pin %d\n", pin)
		p := rpio.Pin(pin)
		p.Output()
		p.Low()
		fmt.Printf("Pin %d set to low (0)\n", pin)
	},
}

func init() {
	lowCmd.PersistentFlags().Uint8VarP(&pin, "pin", "p", 0, "sets the pin to govern (default 0)")
}
