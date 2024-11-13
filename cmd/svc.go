package cmd

import (
	"errors"

	"github.com/cnscottluo/nacos-cli/internal"
	"github.com/cnscottluo/nacos-cli/internal/nacos"
	"github.com/spf13/cobra"
)

var svcCmd = &cobra.Command{
	Use:   "svc",
	Short: "service management",
	Long:  `service management.`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var svcListCmd = &cobra.Command{
	Use:   "list",
	Short: "list services",
	Long:  `list services.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkAddr()
	},
	Run: func(cmd *cobra.Command, args []string) {
		services, err := nacosClient.GetServices(namespace, group)
		internal.CheckErr(err)
		internal.ShowTable(
			[]string{"Service"}, internal.GenData(
				&services.Services, func(resp string) []string {
					return []string{resp}
				},
			),
		)
	},
}

var svcGetCmd = &cobra.Command{
	Use:   "get <serviceName>",
	Short: "get service",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("to many arguments")
		}
		if len(args) < 1 {
			return errors.New("serviceName arg is required")
		}
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkAddr()
	},
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]
		service, err := nacosClient.GetService(namespace, group, serviceName)
		internal.CheckErr(err)
		internal.ShowTable(
			[]string{
				"Namespace", "GroupName", "ServiceName", "Metadata", "ProtectThreshold", "Ephemeral",
			}, internal.GenData(
				&[]nacos.ServiceDetailResp{*service}, func(resp nacos.ServiceDetailResp) []string {
					return []string{
						resp.Namespace,
						resp.GroupName,
						resp.ServiceName,
						internal.ToString(resp.Metadata),
						internal.ToString(resp.ProtectThreshold),
						internal.ToString(resp.Ephemeral),
					}
				},
			),
		)
	},
}

var svcInsCmd = &cobra.Command{
	Use:   "ins <serviceName>",
	Short: "instance management",
	Long:  `instance management.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("too many arguments")
		}
		if len(args) < 1 {
			return errors.New("serviceName arg is required")
		}
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkAddr()
	},
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := args[0]
		instances, err := nacosClient.GetServiceInstances(namespace, group, serviceName)
		internal.CheckErr(err)
		internal.ShowTable(
			[]string{
				"IP", "Port", "Weight", "Healthy", "Enabled", "Ephemeral", "ClusterName", "ServiceName", "Metadata",
			}, internal.GenData(
				&instances.Hosts, func(resp nacos.Host) []string {
					return []string{
						resp.Ip,
						internal.ToString(resp.Port),
						internal.ToString(resp.Weight),
						internal.ToString(resp.Healthy),
						internal.ToString(resp.Enabled),
						internal.ToString(resp.Ephemeral),
						resp.ClusterName,
						resp.ServiceName,
						internal.ToString(resp.Metadata),
					}
				},
			),
		)
	},
}

func init() {
	rootCmd.AddCommand(svcCmd)
	svcCmd.AddCommand(svcListCmd)
	svcCmd.AddCommand(svcGetCmd)
	svcCmd.AddCommand(svcInsCmd)

	svcCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "namespace")
	svcCmd.PersistentFlags().StringVarP(&group, "group", "g", "", "group")
}
