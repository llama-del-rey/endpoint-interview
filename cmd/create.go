package cmd

import (
	"endpoint-interview/directory"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "CREATE [path]",
	Short: "Create a directory",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		directory.CreateDir(args[0])
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
