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
	Short: fmt.Sprintf("Print the version number of %s", ProductName),
	Long:  fmt.Sprintf(`All software has versions. This is %s's`, ProductName),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s v0.1 -- HEAD\n", ProductName)
	},
}