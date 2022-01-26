package cmd

import (
	"context"
	"fmt"
	"strings"
	"todo/data"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Marks the given task as completed",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a task to complete")
			return
		}
		rdb := data.Client()
		data.Done(context.Background(), rdb, strings.Join(args, " "))
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
