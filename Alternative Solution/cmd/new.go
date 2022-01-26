
package cmd

import (
	"context"
	"fmt"
	"strings"
	"todo/data"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "New",
	Short: "Creates a New Task With The Given Text",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please Provide a Task to Add")
			return
		}
		rdb := data.Client()
		data.Add(context.Background(), rdb, strings.Join(args, " "))
		fmt.Println("Task Added Successfully")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
