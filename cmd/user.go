package cmd

import (
	"errors"

	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/spf13/cobra"
)

var (
	username string
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "user management",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var userInitCmd = &cobra.Command{
	Use:   "init <password>",
	Short: "init admin user password",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("password arg is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		password := args[0]
		r, err := nacosClient.InitAdmin(password)
		internal.CheckErr(err)
		internal.Info("Username:%s\nPassword:%s", r.Username, r.Password)
	},
}

var userPassCmd = &cobra.Command{
	Use:   "pass <password>",
	Short: "change user password",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("password arg is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		password := args[0]
		r, err := nacosClient.UpdateUser(username, password)
		internal.CheckErr(err)
		internal.Info("%s", r)
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(userInitCmd)
	userCmd.AddCommand(userPassCmd)
	userPassCmd.Flags().StringVarP(&username, "username", "u", "nacos", "username")
}
