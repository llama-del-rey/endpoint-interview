package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dirmanager",
	Short: "Directory management CLI",
	Long:  `A CLI to create, move, delete, and list directories.`,
}

var fromFileCmd = &cobra.Command{
	Use:   "FROMFILE [file path]",
	Short: "Run commands from a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ProcessCommandsFromFile(args[0])
	},
}

func init() {
	rootCmd.AddCommand(fromFileCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
