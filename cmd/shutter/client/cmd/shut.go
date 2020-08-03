package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/FrancescoIlario/grower/pkg/shutterpb"
	"github.com/spf13/cobra"
)

var (
	openHours, openMinutes, closeHours, closeMinutes uint8

	shutCmd = &cobra.Command{
		Use:   "shut",
		Short: "shut a schedule",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(mctx, 10*time.Second)
			defer cancel()

			if _, err := client.Shut(ctx, &shutterpb.ShutRequest{}); err != nil {
				fmt.Printf("can not create schedule: %v", err)
				return
			}

			fmt.Printf("shut completed")
		},
	}
)
