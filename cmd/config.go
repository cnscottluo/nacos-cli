package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/editor"
	"github.com/cnscottluo/nacos-cli/internal/nacos"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var save bool
var dataId string
var configType string

// config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config management",
	Long:  `config management.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// list config in namespace
var configListCmd = &cobra.Command{
	Use:   "list <namespaceId>",
	Short: "list config in namespace",
	Long:  `list config in namespace.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("namespaceId is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		namespaceId := args[0]
		result, err := nacosClient.GetConfigs(namespaceId)
		internal.CheckErr(err)
		internal.TableShow([]string{"DataId", "Group", "Type"}, internal.GenData(result, func(resp nacos.ConfigResp) []string {
			return []string{
				resp.DataId,
				resp.Group,
				resp.Type,
			}
		}))
	},
}

// get config
var configGetCmd = &cobra.Command{
	Use:   "get <dataId>",
	Short: "get config",
	Long:  `get config.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("dataId is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		dataId := args[0]
		content, err := nacosClient.GetConfig(dataId)
		internal.CheckErr(err)
		internal.ShowConfig(dataId, content)
		if save {
			internal.SaveConfig(dataId, content)
		}
	},
}

func initGet() {
	configGetCmd.Flags().BoolVarP(&save, "save", "s", false, "save config to current directory")
	configCmd.AddCommand(configGetCmd)
}

// delete config
var configDeleteCmd = &cobra.Command{
	Use:   "delete <dataId>",
	Short: "delete config",
	Long:  `delete config.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("dataId is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		dataId := args[0]
		_, err := nacosClient.DeleteConfig(dataId)
		internal.CheckErr(err)
		internal.Info("delete config %s success", dataId)
	},
}

// publish config
var configPublishCmd = &cobra.Command{
	Use:   "publish <file-path>",
	Short: "publish config",
	Long:  `publish config.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("file-path is required")
		}
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		filePath := args[0]
		if len(dataId) == 0 {
			dataId = nacos.DetermineDataId(filePath)
		}
		if len(configType) == 0 {
			configType = nacos.DetermineConfigType(filePath)
		} else {
			if !nacos.IsValidConfigType(configType) {
				return errors.New("invalid config type")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		content, err := internal.ReadFile(filePath)
		internal.CheckErr(err)
		_, err = nacosClient.PublishConfig(dataId, content, configType)
		internal.CheckErr(err)
		internal.Info("publish config %s success", dataId)
	},
}

// init publish command
func initPublish() {
	configPublishCmd.Flags().StringVarP(&dataId, "data-id", "d", "", "config file path")
	configPublishCmd.Flags().StringVarP(&configType, "type", "t", "", "config type (text,json,xml,yaml,html,properties)")
	configCmd.AddCommand(configPublishCmd)
}

// edit config
var configEditCmd = &cobra.Command{
	Use:   "edit <dataId>",
	Short: "edit config",
	Long:  `edit config.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("dataId is required")
		}
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		dataId := args[0]
		if len(configType) == 0 {
			configType = nacos.DetermineConfigType(dataId)
		} else {
			if !nacos.IsValidConfigType(configType) {
				return errors.New("invalid config type")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		dataId := args[0]
		content, err := nacosClient.GetConfig(dataId)
		internal.CheckErr(err)
		originMD5 := internal.GenStringMD5(content)
		internal.VerboseLog("origin config md5: %s", originMD5)

		// use default editor
		editor := editor.NewDefaultEditor([]string{})

		buf := &bytes.Buffer{}
		buf.Write([]byte(content))

		editedContent, filePath, err := editor.LaunchTempFile(fmt.Sprintf("%s-edit-", filepath.Base(os.Args[0])), dataId, buf)
		internal.CheckErr(err)

		editedMD5 := internal.GenBytesMD5(editedContent)
		if originMD5 == editedMD5 {
			internal.VerboseLog("no change")
			return
		}
		defer func(f string) {
			if e := os.Remove(f); e != nil {
				internal.VerboseLog("delete temp filePath %s error: %s", f, e)
			}
		}(filePath)
		_, err = nacosClient.PublishConfig(dataId, string(editedContent), configType)
		internal.CheckErr(err)
		internal.Info("edit config %s success", dataId)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	//configCmd.PersistentFlags().StringP("namespaceId", "n", "", "namespaceId")
	//configCmd.PersistentFlags().StringP("group", "g", "DEFAULT_GROUP", "group")
	//_ = viper.BindPFlag("nacos.namespace", configCmd.Flags().Lookup("namespaceId"))
	//_ = viper.BindPFlag("nacos.group", configCmd.Flags().Lookup("group"))

	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configDeleteCmd)
	configCmd.AddCommand(configEditCmd)
	initGet()
	initPublish()
}
