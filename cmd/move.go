package cmd

import (
	"endpoint-interview/directory"
	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:   "MOVE [source] [destination]",
	Short: "Move a directory",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		directory.MoveDir(args[0], args[1])
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)
}
