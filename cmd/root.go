package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/nacos"
	"github.com/cnscottluo/nacos-cli/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"github.com/spf13/viper"
)

var cfgFile string
var nacosClient *nacos.Client
var (
	namespace string
	group     string
)

var config = new(types.Config)

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
	e := doc.GenMarkdownTree(rootCmd, "./docs")
	if e != nil {
		log.Fatal(e)
	}
	println("执行 execute")
	err := rootCmd.Execute()
	if err != nil {
		internal.Error("%s", err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "setting", "", "setting file (default is $HOME/.nacos.toml)")
	rootCmd.PersistentFlags().BoolVar(&internal.Verbose, "verbose", false, "verbose output")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		internal.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".nacos")
	}

	if err := viper.ReadInConfig(); err == nil {
		internal.VerboseLog("Using setting file: %s", viper.ConfigFileUsed())
	}

	err := viper.Unmarshal(config)
	internal.CheckErr(err)
	nacosClient = nacos.NewClient(config)

	for key, value := range viper.AllSettings() {
		internal.VerboseLog("%s: %+v", key, value)
	}
}

func checkAddr() error {
	if len(config.Nacos.Addr) == 0 {
		return errors.New("nacos addr is required, place execute init command to initialize setting")
	}
	return nil
}
