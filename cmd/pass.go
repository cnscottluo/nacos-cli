package cmd

import (
	"errors"

	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/spf13/cobra"
)

var username string
var password string

var passCmd = &cobra.Command{
	Use:   "pass",
	Short: "user password management",
	Long:  "user password management",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "change password",
	Long:  `change user password.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if username == "" {
			return errors.New("username is required")
		}
		if password == "" {
			return errors.New("password is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		r, err := nacosClient.UpdatePassword(username, password)
		internal.CheckErr(err)
		internal.Info("%s", r)
	},
}

func init() {
	rootCmd.AddCommand(passCmd)
	passCmd.AddCommand(changeCmd)
	changeCmd.Flags().StringVarP(&username, "username", "u", "nacos", "username")
	changeCmd.Flags().StringVarP(&password, "password", "p", "", "password")
}
