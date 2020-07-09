package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/stianeikeland/go-rpio/v4"
)

var pin uint8 = 0

var rootCmd = &cobra.Command{
	Use:   "pinctl",
	Short: "pinctl is the CLI to manually control the rpi pins",
}

func init() {
	rootCmd.AddCommand(lowCmd, highCmd)
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
