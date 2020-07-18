package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/FrancescoIlario/grower/pkg/schedulerpb"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list all schedules",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(mctx, 10*time.Second)
			defer cancel()

			schedule, err := client.ListSchedules(ctx, &schedulerpb.ListSchedulesRequest{})
			if err != nil {
				fmt.Printf("can not retrieve schedules: %v", err)
				return
			}

			schedulesJB, err := json.Marshal(schedule)
			if err != nil {
				fmt.Printf("error marshaling schedules to json: %v", err)
				return
			}

			fmt.Printf("%s", string(schedulesJB))
		},
	}
)
