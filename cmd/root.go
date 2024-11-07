package cmd

import (
	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/nacos"
	"github.com/cnscottluo/nacos-cli/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
var nacosClient *nacos.Client
var Config = new(types.Config)

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
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "", "config file (default is $HOME/.nacos.toml)")
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
		internal.Log("Using config file: %s", viper.ConfigFileUsed())
	}

	err := viper.Unmarshal(Config)
	internal.CheckErr(err)
	nacosClient = nacos.NewClient(Config)
}
