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
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get a schedule by id",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(mctx, 10*time.Second)
			defer cancel()

			schedule, err := client.GetSchedule(ctx, &schedulerpb.GetScheduleRequest{Id: id})
			if err != nil {
				fmt.Printf("can not retrieve schedule: %v", err)
				return
			}

			scheduleJB, err := json.Marshal(schedule)
			if err != nil {
				fmt.Printf("error marshaling schedule to json: %v", err)
				return
			}

			fmt.Printf("%s", string(scheduleJB))
		},
	}
)

func init() {
	getCmd.Flags().StringVarP(&id, "id", "i", "", "the id of the schedule to retrieve")
}
