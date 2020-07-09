package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stianeikeland/go-rpio/v4"
)

var gpiochip string
var positiveRelayPin uint8
var negativeRelayPin uint8

var rootCmd = &cobra.Command{
	Use:   "valvectl",
	Short: "Valvectl is the CLI to manually control the valve",
}

func init() {
	rootCmd.PersistentFlags().Uint8VarP(&positiveRelayPin, "ppin", "p", 10, "pin for the relay that control the valve's positive")
	rootCmd.PersistentFlags().Uint8VarP(&negativeRelayPin, "npin", "n", 4, "pin for the relay that control the valve's negative")
	rootCmd.AddCommand(statusCmd, openCmd, closeCmd)
}

// Execute ...
func Execute() {
	rpio.Open()
	defer rpio.Close()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
