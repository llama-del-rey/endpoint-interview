package cmd

import (
	"endpoint-interview/directory"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "DELETE [path]",
	Short: "Delete a directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		directory.DeleteDir(args[0])
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
