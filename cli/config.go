package cli

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Nexus switch configurations",
	Long:  `List and describe Nexus switch configurations`,
}

func init() {
	RootCmd.AddCommand(configCmd)
}
