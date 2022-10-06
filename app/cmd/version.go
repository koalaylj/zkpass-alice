package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "alice version",
	Long:  `Print the version zkpass alice`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("zkpass test client alice v0.1")
	},
}
