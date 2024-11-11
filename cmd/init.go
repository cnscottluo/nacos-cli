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
	Use:   "init <addr> [username] [password] [namespaceId] [groupName]",
	Short: "init nacos",
	Long:  `init nacos.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("addr arg is required")
		}
		if len(args) > 5 {
			return errors.New("too many arguments")
		}
		if auth && len(args) < 3 {
			return errors.New("username and password args is required when auth is true")
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
	viper.Set("nacos.addr", args[0])
	if len(args) > 1 {
		viper.Set("nacos.username", args[1])
	}
	if len(args) > 2 {
		viper.Set("nacos.password", args[2])
	}
	if len(args) > 3 {
		viper.Set("nacos.namespaceId", args[3])
	} else {
		viper.Set("nacos.namespaceId", "public")
	}
	if len(args) > 4 {
		viper.Set("nacos.groupName", args[4])
	} else {
		viper.Set("nacos.groupName", "DEFAULT_GROUP")
	}
	viper.Set("nacos.auth", auth)
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVar(&auth, "auth", false, "use username and password to authenticate")
}
