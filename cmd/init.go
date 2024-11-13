package cmd

import (
	"errors"
	"os"
	"path"

	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var auth bool

var initCmd = &cobra.Command{
	Use:   "init [addr] [username] [password] [namespace] [group]",
	Short: "init nacos",
	Long: `init nacos.
addr: (default is http://127.0.0.1:8848/nacos)
username: (default is nacos)
password: (default is nacos)
namespace: (default is public)
group: (default is DEFAULT_GROUP)`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 5 {
			return errors.New("too many arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		setArgs(args)
		detectConfigFile()
		if auth {
			loginResponse, err := nacosClient.Login(
				viper.GetString("nacos.addr"), viper.GetString("nacos.username"), viper.GetString("nacos.password"),
			)
			internal.CheckErr(err)
			viper.Set("nacos.token", loginResponse.AccessToken)
			internal.Info("%s login success", viper.GetString("nacos.username"))
		}
		err := viper.WriteConfig()
		internal.CheckErr(err)
	},
}

func detectConfigFile() {
	home, err := os.UserHomeDir()
	internal.CheckErr(err)

	if _, err := os.Stat(path.Join(home, ".nacos.toml")); os.IsNotExist(err) {
		_, err = os.Create(path.Join(home, ".nacos.toml"))
		internal.CheckErr(err)
	}
}

func setArgs(args []string) {
	if len(args) > 0 {
		viper.Set("nacos.addr", args[0])
	} else {
		viper.Set("nacos.addr", "http://127.0.0.1:8848/nacos")
	}
	internal.VerboseLog("addr: %s", viper.GetString("nacos.addr"))
	if len(args) > 1 {
		viper.Set("nacos.username", args[1])
	} else {
		viper.Set("nacos.username", "nacos")
	}
	internal.VerboseLog("username: %s", viper.GetString("nacos.username"))
	if len(args) > 2 {
		viper.Set("nacos.password", args[2])
	} else {
		viper.Set("nacos.password", "nacos")
	}
	internal.VerboseLog("password: %s", viper.GetString("nacos.password"))
	if len(args) > 3 {
		viper.Set("nacos.namespace", args[3])
	} else {
		viper.Set("nacos.namespace", "public")
	}
	internal.VerboseLog("namespace: %s", viper.GetString("nacos.namespace"))
	if len(args) > 4 {
		viper.Set("nacos.group", args[4])
	} else {
		viper.Set("nacos.group", "DEFAULT_GROUP")
	}
	internal.VerboseLog("group: %s", viper.GetString("nacos.group"))
	viper.Set("nacos.auth", auth)
	internal.VerboseLog("auth: %t", viper.GetBool("nacos.auth"))
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVar(&auth, "auth", false, "use username and password to authenticate")
}
