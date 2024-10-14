package cmd

import (
	"grr/gui"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grr",
	Short: "grr is 'go rss reader'",
	Run: func(cmd *cobra.Command, args []string) {
		gui.Start()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
