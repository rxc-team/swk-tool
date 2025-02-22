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
	Short: "Print the version number of tool",
	Long:  `All software has versions. This is tool's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version: v1.0.0")
	},
}
