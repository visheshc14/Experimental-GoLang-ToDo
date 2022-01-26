
package cmd

import (
	"context"
	"fmt"
	"todo/data"

	"github.com/spf13/cobra"
)

var Verbose bool
var Source string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the tasks",
	Run: func(cmd *cobra.Command, args []string) {

		rdb := data.Client()
		done, err := cmd.Flags().GetBool("done")
		if err != nil {
			fmt.Println("Could not list", err)
		}
		data.List(context.Background(), rdb, done)
	},
}

func init() {
	listCmd.PersistentFlags().BoolVarP(&Verbose, "done", "d", false, "list of done tasks")
	rootCmd.AddCommand(listCmd)
}
