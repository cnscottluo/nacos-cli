package cmd

import (
	"errors"
	"strconv"

	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/nacos"

	"github.com/spf13/cobra"
)

var nsCmd = &cobra.Command{
	Use:   "ns",
	Short: "namespaceId management",
	Long:  "namespaceId management",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var nsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list namespaces",
	Long:  "list namespaces",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := nacosClient.GetNamespaces()
		internal.CheckErr(err)
		internal.ShowTable(
			[]string{"Namespace", "Name", "Desc", "Quota", "Count", "Type"}, internal.GenData(
				result, func(resp nacos.NamespaceResp) []string {
					return []string{
						resp.Namespace,
						resp.NamespaceShowName,
						resp.NamespaceDesc,
						strconv.Itoa(resp.Quota),
						strconv.Itoa(resp.ConfigCount),
						func() string {
							switch resp.Type {
							case 0:
								return "Global"
							case 1:
								return "Private"
							case 2:
								return "Custom"
							default:
								return "Unknown"
							}
						}(),
					}
				},
			),
		)
	},
}

var nsGetCmd = &cobra.Command{
	Use:   "get [namespaceId]",
	Short: "get namespaceId",
	Long:  `get namespaceId.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("too many arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		namespaceId := ""
		if len(args) == 1 {
			namespaceId = args[0]
		}
		result, err := nacosClient.GetNamespace(namespaceId)
		internal.CheckErr(err)
		internal.ShowTable(
			[]string{"Namespace", "Name", "Desc", "Quota", "Count", "Type"}, internal.GenData(
				&[]nacos.NamespaceResp{*result}, func(resp nacos.NamespaceResp) []string {
					return []string{
						resp.Namespace,
						resp.NamespaceShowName,
						resp.NamespaceDesc,
						strconv.Itoa(resp.Quota),
						strconv.Itoa(resp.ConfigCount),
						func() string {
							switch resp.Type {
							case 0:
								return "Global"
							case 1:
								return "Private"
							case 2:
								return "Custom"
							default:
								return "Unknown"
							}
						}(),
					}
				},
			),
		)
	},
}

var nsCreateCmd = &cobra.Command{
	Use:   "create <namespaceId> <namespaceName> [namespaceDesc]",
	Short: "create namespaceId",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 || len(args) > 3 {
			return errors.New("namespaceId and namespaceName args is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		namespaceId := args[0]
		namespaceName := args[1]
		namespaceDesc := ""
		if len(args) == 3 {
			namespaceDesc = args[2]
		}
		_, err := nacosClient.CreateNamespace(namespaceId, namespaceName, namespaceDesc)
		internal.CheckErr(err)
		internal.Info("create namespaceId %s success", namespaceId)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update <namespaceId> <namespaceName> [namespaceDesc]",
	Short: "update namespaceId",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 || len(args) > 3 {
			return errors.New("namespaceId and namespaceName args is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		namespaceId := args[0]
		namespaceName := args[1]
		namespaceDesc := ""
		if len(args) == 3 {
			namespaceDesc = args[2]
		}
		_, err := nacosClient.UpdateNamespace(namespaceId, namespaceName, namespaceDesc)
		internal.CheckErr(err)
		internal.Info("update namespaceId %s success", namespaceId)
	},
}

var nsDeleteCmd = &cobra.Command{
	Use:   "delete <namespaceId>",
	Short: "delete namespaceId",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("namespaceId arg is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		namespaceId := args[0]
		_, err := nacosClient.DeleteNamespace(namespaceId)
		internal.CheckErr(err)
		internal.Info("delete namespaceId %s success", namespaceId)
	},
}

func init() {
	rootCmd.AddCommand(nsCmd)
	nsCmd.AddCommand(nsListCmd)
	nsCmd.AddCommand(nsGetCmd)
	nsCmd.AddCommand(nsCreateCmd)
	nsCmd.AddCommand(updateCmd)
	nsCmd.AddCommand(nsDeleteCmd)
}
