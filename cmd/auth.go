package cmd

import (
	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/setting"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Auth user",
	Long:  `Auth user by username and password`,
	PreRun: func(cmd *cobra.Command, args []string) {
		setting.CreateIfNotExistConfigFile()
		bindViper(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		loginResponse, err := nacosClient.Login(
			viper.GetString(setting.NacosAddrKey), viper.GetString(setting.NacosUsernameKey),
			viper.GetString(setting.NacosPasswordKey),
		)
		internal.CheckErr(err)
		err = setting.SaveSetting(loginResponse.AccessToken)
		internal.CheckErr(err)
	},
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear user auth",
	Long:  `Clear user auth`,
	Run: func(cmd *cobra.Command, args []string) {
		setting.DeleteConfigFile()
	},
}

func bindViper(cmd *cobra.Command) {
	_ = viper.BindPFlag(setting.NacosAddrKey, cmd.Flags().Lookup("addr"))
	_ = viper.BindPFlag(setting.NacosUsernameKey, cmd.Flags().Lookup("username"))
	_ = viper.BindPFlag(setting.NacosPasswordKey, cmd.Flags().Lookup("password"))
	_ = viper.BindPFlag(setting.NacosNamespaceKey, cmd.Flags().Lookup("namespace"))
	_ = viper.BindPFlag(setting.NacosGroupKey, cmd.Flags().Lookup("group"))
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.Flags().StringP("addr", "a", setting.DefaultAddr, "nacos server address")
	authCmd.Flags().StringP("username", "u", setting.DefaultUsername, "nacos username")
	authCmd.Flags().StringP("password", "p", setting.DefaultPassword, "nacos password")
	authCmd.Flags().StringP("namespace", "n", setting.DefaultNamespace, "nacos namespace")
	authCmd.Flags().StringP("group", "g", setting.DefaultGroup, "nacos group")
	authCmd.AddCommand(clearCmd)
}
