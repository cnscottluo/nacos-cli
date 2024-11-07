package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config management",
	Long:  `config management.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var getCmd = &cobra.Command{
	Use:   "get []",
	Short: "get config",
	Long:  `get config.`,
}

func init() {
	rootCmd.AddCommand(configCmd)

}
