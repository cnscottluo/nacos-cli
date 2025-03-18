package cmd

import (
	"errors"

	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User management",
	Long:  `User management`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var initCmd = &cobra.Command{
	Use:   "init <password>",
	Short: "Init admin user password",
	Long:  `Init admin user password.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("password is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		r, err := nacosClient.InitAdmin(args[0])
		internal.CheckErr(err)
		internal.Info("%s", r)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(initCmd)
}
