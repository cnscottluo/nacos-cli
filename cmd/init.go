package cmd

import (
	"fmt"
	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/spf13/viper"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var auth bool

var initCmd = &cobra.Command{
	Use:   "init <addr> [username] [password] [namespace] [group]",
	Short: "init nacos config",
	Long:  `init nacos config.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("addr is required")
		}
		if len(args) > 5 {
			return fmt.Errorf("too many arguments")
		}
		if auth && len(args) < 3 {
			return fmt.Errorf("username and password is required when auth is true")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		setArgs(args)
		detectConfigFile()
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
		viper.Set("nacos.namespace", args[3])
	} else {
		viper.Set("nacos.namespace", "public")
	}
	if len(args) > 4 {
		viper.Set("nacos.group", args[4])
	} else {
		viper.Set("nacos.group", "DEFAULT_GROUP")
	}
	viper.Set("nacos.auth", auth)
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVar(&auth, "auth", false, "use username and password to authenticate")
}
