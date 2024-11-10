package cmd

import (
	"fmt"
	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/nacos"
	"github.com/cnscottluo/nacos-cli/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var nacosClient *nacos.Client
var config = new(types.Config)

var rootCmd = &cobra.Command{
	Use:   "nacos-cli",
	Short: "A CLI tool for Nacos",
	Long:  "A CLI tool for Nacos",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nacos.toml)")
	rootCmd.PersistentFlags().BoolVar(&internal.Verbose, "verbose", false, "verbose output")

	rootCmd.PersistentFlags().String("addr", "http://127.0.0.1:8848/nacos", "nacos server address")
	rootCmd.PersistentFlags().StringP("username", "u", "nacos", "nacos server username")
	rootCmd.PersistentFlags().StringP("password", "p", "nacos", "nacos server password")
	rootCmd.PersistentFlags().StringP("namespace", "n", "public", "nacos server namespace")
	rootCmd.PersistentFlags().StringP("group", "g", "DEFAULT_GROUP", "nacos server group")

	_ = viper.BindPFlag("nacos.addr", rootCmd.PersistentFlags().Lookup("addr"))
	_ = viper.BindPFlag("nacos.username", rootCmd.PersistentFlags().Lookup("username"))
	_ = viper.BindPFlag("nacos.password", rootCmd.PersistentFlags().Lookup("password"))
	_ = viper.BindPFlag("nacos.namespace", rootCmd.PersistentFlags().Lookup("namespace"))
	_ = viper.BindPFlag("nacos.group", rootCmd.PersistentFlags().Lookup("group"))

}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".nacos")
	}

	if err := viper.ReadInConfig(); err == nil {
		internal.Log("Using config file: %s", viper.ConfigFileUsed())
	}

	err := viper.Unmarshal(config)
	internal.CheckErr(err)
	nacosClient = nacos.NewClient(config)

	for key, value := range viper.AllSettings() {
		fmt.Printf("%s: %v\n", key, value)
	}
}
