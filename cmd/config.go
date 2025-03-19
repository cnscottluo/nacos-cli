package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/editor"
	"github.com/cnscottluo/nacos-cli/internal/nacos"
	"github.com/spf13/cobra"
)

var save bool

var (
	namespaceId string
	group       string
	dataId      string
	configType  string
)

// config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config management",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

// list config in namespaceId
var configListCmd = &cobra.Command{
	Use:   "list [namespaceId]",
	Short: "list config in namespaceId",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("too many arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			namespaceId = args[0]
		}
		result, err := nacosClient.GetConfigs(namespaceId)
		internal.CheckErr(err)
		internal.ShowTable(
			[]string{"DataId", "Group", "Type"}, internal.GenData(
				result, func(resp nacos.ConfigResp) []string {
					return []string{
						resp.DataId,
						resp.Group,
						resp.Type,
					}
				},
			),
		)
	},
}

// get config
var configGetCmd = &cobra.Command{
	Use:   "get <dataId>",
	Short: "get config",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("too many arguments")
		}
		if len(args) < 1 {
			return errors.New("dataId arg is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		dataId := args[0]
		content, err := nacosClient.GetConfig(namespaceId, group, dataId)
		internal.CheckErr(err)
		internal.ShowConfig(dataId, content)
		if save {
			err := internal.SaveConfig(
				nacosClient.GetNamespaceId(namespaceId), nacosClient.GetGroup(group), dataId, content,
			)
			internal.CheckErr(err)
		}
	},
}

// delete config
var configDeleteCmd = &cobra.Command{
	Use:   "delete <dataId>",
	Short: "delete config",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("too many arguments")
		}
		if len(args) < 1 {
			return errors.New("dataId arg is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		dataId := args[0]
		_, err := nacosClient.DeleteConfig(namespaceId, group, dataId)
		internal.CheckErr(err)
		internal.Info("delete config %s success(maybe error result)", dataId)
	},
}

// apply config
var configApplyCmd = &cobra.Command{
	Use:   "apply <file>",
	Short: "apply config",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("too many arguments")
		}
		if len(args) < 1 {
			return errors.New("file arg is required")
		}
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		file := args[0]
		if dataId == "" {
			dataId = nacos.DetermineDataId(file)
		}
		if configType == "" {
			configType = nacos.DetermineConfigType(file)
		} else {
			if !nacos.IsValidConfigType(configType) {
				return errors.New("invalid config type")
			}
			configType = strings.ToLower(configType)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		content, err := internal.ReadFile(file)
		internal.CheckErr(err)
		_, err = nacosClient.ApplyConfig(namespaceId, group, dataId, content, configType)
		internal.CheckErr(err)
		internal.Info("apply config %s success", dataId)
	},
}

// edit config
var configEditCmd = &cobra.Command{
	Use:   "edit <dataId>",
	Short: "edit config",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("too many arguments")
		}
		if len(args) < 1 {
			return errors.New("dataId arg is required")
		}
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		dataId := args[0]
		if configType == "" {
			configType = nacos.DetermineConfigType(dataId)
		} else {
			if !nacos.IsValidConfigType(configType) {
				return errors.New("invalid config type")
			}
			configType = strings.ToLower(configType)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		dataId := args[0]
		content, err := nacosClient.GetConfig(namespaceId, group, dataId)
		internal.CheckErr(err)
		originMD5 := internal.GenStringMD5(content)
		internal.VerboseLog("origin config md5: %s", originMD5)

		// use default editor
		editor := editor.NewDefaultEditor([]string{})

		buf := &bytes.Buffer{}
		buf.Write([]byte(content))

		editedContent, filePath, err := editor.LaunchTempFile(
			fmt.Sprintf("%s-edit-", filepath.Base(os.Args[0])), dataId, buf,
		)
		internal.CheckErr(err)
		defer func(f string) {
			if e := os.Remove(f); e != nil {
				internal.VerboseLog("delete temp file %s error: %s", f, e)
			}
		}(filePath)

		editedMD5 := internal.GenBytesMD5(editedContent)
		if originMD5 == editedMD5 {
			internal.VerboseLog("no change")
		} else {
			_, err = nacosClient.ApplyConfig(namespaceId, group, dataId, string(editedContent), configType)
			internal.CheckErr(err)
			internal.Info("edit config %s success", dataId)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configDeleteCmd)
	configCmd.AddCommand(configApplyCmd)
	configCmd.AddCommand(configEditCmd)

	configGetCmd.Flags().BoolVarP(&save, "save", "s", false, "save config to current directory")
	configGetCmd.Flags().StringVarP(&namespaceId, "namespaceId", "n", "", "namespaceId")
	configGetCmd.Flags().StringVarP(&group, "group", "g", "", "group")

	configDeleteCmd.Flags().StringVarP(&namespaceId, "namespaceId", "n", "", "namespaceId")
	configDeleteCmd.Flags().StringVarP(&group, "group", "g", "", "group")

	configApplyCmd.Flags().StringVarP(&namespaceId, "namespaceId", "n", "", "namespaceId")
	configApplyCmd.Flags().StringVarP(&group, "group", "g", "", "group")
	configApplyCmd.Flags().StringVarP(&dataId, "data-id", "d", "", "config data id")
	configApplyCmd.Flags().StringVarP(
		&configType, "type", "t", "", "config type (text,json,xml,yaml,html,properties,toml)",
	)

	configEditCmd.Flags().StringVarP(&namespaceId, "namespaceId", "n", "", "namespaceId")
	configEditCmd.Flags().StringVarP(&group, "group", "g", "", "group")
	configEditCmd.Flags().StringVarP(
		&configType, "type", "t", "", "config type (text,json,xml,yaml,html,properties,toml)",
	)
}
