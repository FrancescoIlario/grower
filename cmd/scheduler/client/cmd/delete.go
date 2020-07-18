package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/FrancescoIlario/grower/pkg/schedulerpb"
	"github.com/spf13/cobra"
)

var (
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "delete a schedule by id",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(mctx, 10*time.Second)
			defer cancel()

			_, err := client.DeleteSchedule(ctx, &schedulerpb.DeleteScheduleRequest{Id: id})
			if err != nil {
				fmt.Printf("can not delete schedule: %v", err)
				return
			}

			fmt.Printf("schedule deleted %s", id)
		},
	}
)

func init() {
	deleteCmd.Flags().StringVarP(&id, "id", "i", "", "the id to delete")
}
