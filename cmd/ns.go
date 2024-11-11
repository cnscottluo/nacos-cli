package cmd

import (
	"fmt"
	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/nacos"
	"strconv"

	"github.com/spf13/cobra"
)

var all bool

var nsCmd = &cobra.Command{
	Use:   "ns",
	Short: "namespace management",
	Long:  "namespace management",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var nsGetCmd = &cobra.Command{
	Use:   "get [namespaceId]",
	Short: "get namespace",
	Args: func(cmd *cobra.Command, args []string) error {
		if all && len(args) > 0 {
			return fmt.Errorf("cannot use both --all and namespaceId")
		}
		if !all && len(args) == 0 {
			return fmt.Errorf("requires namespaceId")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if all {
			internal.VerboseLog("get all namespaces")
			result, err := nacosClient.GetNamespaces()
			internal.CheckErr(err)
			internal.TableShow([]string{"Namespace", "Name", "Desc", "Quota", "Count", "Type"}, internal.GenData(result, func(resp nacos.NamespaceResp) []string {
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
			}))
		} else {
			internal.VerboseLog("get namespace %s", args[0])
			result, err := nacosClient.GetNamespace(args[0])
			internal.CheckErr(err)
			internal.TableShow([]string{"Namespace", "Name", "Desc", "Quota", "Count", "Type"}, internal.GenData(&[]nacos.NamespaceResp{*result}, func(resp nacos.NamespaceResp) []string {
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
			}))
		}
	},
}

var nsCreateCmd = &cobra.Command{
	Use:   "create <namespaceId> <namespaceName> [namespaceDesc]",
	Short: "create namespace",
	Args:  cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		namespaceId := args[0]
		namespaceName := args[1]
		namespaceDesc := ""
		if len(args) == 3 {
			namespaceDesc = args[2]
		}
		internal.VerboseLog("create namespace %s %s %s", namespaceId, namespaceName, namespaceDesc)
		_, err := nacosClient.CreateNamespace(namespaceId, namespaceName, namespaceDesc)
		internal.CheckErr(err)
		internal.Info("create namespace %s success", namespaceId)
	},
}

var updateCmd = &cobra.Command{
	Use:   "update <namespaceId> <namespaceName> [namespaceDesc]",
	Short: "update namespace",
	Args:  cobra.RangeArgs(2, 3),
	Run: func(cmd *cobra.Command, args []string) {
		namespaceId := args[0]
		namespaceName := args[1]
		namespaceDesc := ""
		if len(args) == 3 {
			namespaceDesc = args[2]
		}
		internal.VerboseLog("update namespace %s %s %s", namespaceId, namespaceName, namespaceDesc)
		_, err := nacosClient.UpdateNamespace(namespaceId, namespaceName, namespaceDesc)
		internal.CheckErr(err)
		internal.Info("update namespace %s success", namespaceId)
	},
}

var nsDeleteCmd = &cobra.Command{
	Use:   "delete <namespaceId>",
	Short: "delete namespace",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		internal.VerboseLog("delete namespace %s", args[0])
		_, err := nacosClient.DeleteNamespace(args[0])
		internal.CheckErr(err)
		internal.Info("delete namespace %s success", args[0])
	},
}

func init() {
	rootCmd.AddCommand(nsCmd)
	nsCmd.AddCommand(nsGetCmd)
	nsCmd.AddCommand(nsCreateCmd)
	nsCmd.AddCommand(updateCmd)
	nsCmd.AddCommand(nsDeleteCmd)

	nsGetCmd.Flags().BoolVarP(&all, "all", "a", false, "get all namespaces")
}
