package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/FrancescoIlario/grower/pkg/schedulerpb"
	"github.com/spf13/cobra"
)

var (
	openHours, openMinutes, closeHours, closeMinutes uint8

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "add a schedule",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(mctx, 10*time.Second)
			defer cancel()

			sched, err := client.CreateSchedule(ctx, &schedulerpb.CreateScheduleRequest{
				OpenTime: &schedulerpb.TimePoint{
					Hours:   int32(openHours),
					Minutes: int32(openMinutes),
				},
				CloseTime: &schedulerpb.TimePoint{
					Hours:   int32(closeHours),
					Minutes: int32(closeMinutes),
				},
			})
			if err != nil {
				fmt.Printf("can not create schedule: %v", err)
				return
			}

			fmt.Printf("schedule created %s", sched.GetId())
		},
	}
)

func init() {
	addCmd.Flags().Uint8VarP(&openHours, "open-hours", "oh", 0, "the hours at which the valve must open")
	addCmd.Flags().Uint8VarP(&openMinutes, "open-minutes", "om", 0, "the minutes at which the valve must open")
	addCmd.Flags().Uint8VarP(&closeHours, "close-hours", "oh", 0, "the hours at which the valve must close")
	addCmd.Flags().Uint8VarP(&closeMinutes, "close-minutes", "om", 0, "the minutes at which the valve must close")
}
