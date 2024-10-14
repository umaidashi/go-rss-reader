package cmd

import (
	"fmt"
	"grr/utils"

	"github.com/spf13/cobra"
)

var coffeeCmd = &cobra.Command{
	Use:   "coffee",
	Short: "Print a coffee cup",
	Run: func(cmd *cobra.Command, args []string) {
		aa := utils.GetAAFromFp("coffee.txt")
		fmt.Println(aa)
	},
}

func init() {
	rootCmd.AddCommand(coffeeCmd)
}
