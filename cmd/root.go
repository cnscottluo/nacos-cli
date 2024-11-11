package cmd

import (
	"os"

	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/nacos"
	"github.com/cnscottluo/nacos-cli/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var nacosClient *nacos.Client
var (
	namespaceId string
	groupName   string
)

var rootCmd = &cobra.Command{
	Use:           "nacos-cli",
	Short:         "A CLI tool for Nacos",
	Long:          "A CLI tool for Nacos",
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		internal.Error("%s", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nacos.toml)")
	rootCmd.PersistentFlags().BoolVar(&internal.Verbose, "verbose", false, "verbose output")
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
		internal.VerboseLog("Using config file: %s", viper.ConfigFileUsed())
	}

	var config = new(types.Config)
	err := viper.Unmarshal(config)
	internal.CheckErr(err)
	nacosClient = nacos.NewClient(config)

	for key, value := range viper.AllSettings() {
		internal.VerboseLog("%s: %+v", key, value)
	}
}
