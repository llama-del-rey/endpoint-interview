package cmd

import (
	"endpoint-interview/directory"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "LIST",
	Short: "List directories",
	Run: func(cmd *cobra.Command, args []string) {
		directory.ListDirs(directory.Root(), "")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
