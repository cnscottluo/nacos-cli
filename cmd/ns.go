package cmd

import (
	"fmt"
	"github.com/cnscottluo/nacos-cli/internal"

	"github.com/spf13/cobra"
)

var all bool

var nsCmd = &cobra.Command{
	Use:   "ns",
	Short: "namespace management",
	Long:  "namespace management",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ns called")
	},
}

var getCmd = &cobra.Command{
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
			internal.Log("get all namespaces")
			result, err := nacosClient.GetNamespaces()
			internal.CheckErr(err)
			internal.TableShow([]interface{}{result}, "Namespace", "Name", "Desc", "Quota", "Count", "Type")
		} else {
			internal.Log("get namespace %s", args[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(nsCmd)
	nsCmd.AddCommand(getCmd)

	nsCmd.Flags().BoolVarP(&all, "all", "a", false, "get all namespaces")
}
