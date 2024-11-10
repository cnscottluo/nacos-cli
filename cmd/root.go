package cmd

import (
	"fmt"
	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/nacos"
	"github.com/cnscottluo/nacos-cli/internal/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
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

func handleError(err error) {
	if err != nil {
		// 记录错误到日志文件
		log.Printf("Error: %v\n", err)
		// 打印到标准错误流
		fmt.Println("Custom Error:", err)
		// 终止程序
		os.Exit(1)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		//handleError(err)
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
		internal.Log("Using config file: %s", viper.ConfigFileUsed())
	}

	err := viper.Unmarshal(Config)
	internal.CheckErr(err)
	nacosClient = nacos.NewClient(Config)
}
