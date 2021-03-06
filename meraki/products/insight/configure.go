package insight

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/insight/configure"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)

var GetMonitoredMediaServers = &cobra.Command{
	Use:   "monitoredMediaServers",
	Short: "List the monitored media servers for this organization. Only valid for organizations with Meraki Insight.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, _, _ := shell.ResolveFlags(cmd.Flags())
		if orgId == "" {
			orgId = args[0]
		}
		metadata := configure.GetMonitoredMediaServers(orgId)
		shell.Display(metadata, "MonitoredMediaServers", cmd.Flags())
	},
}

var GetMonitoredMediaServer = &cobra.Command{
	Use:   "monitoredMediaServer",
	Short: "Return a monitored media server for this organization. Only valid for organizations with Meraki Insight.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, _, _ := shell.ResolveFlags(cmd.Flags())
		if orgId == "" {
			orgId = args[1]
		}
		monitoredMediaServerId := args[0]
		metadata := configure.GetMonitoredMediaServer(orgId, monitoredMediaServerId)
		shell.Display(metadata, "MonitoredMediaServer", cmd.Flags())
	},
}

var DelMonitoredMediaServer = &cobra.Command{
	Use:   "monitoredMediaServer",
	Short: "Delete a monitored media server for this organization. Only valid for organizations with Meraki Insight.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, _, _ := shell.ResolveFlags(cmd.Flags())
		if orgId == "" {
			orgId = args[1]
		}
		monitoredMediaServerId := args[0]
		metadata := configure.DelMonitoredMediaServer(orgId, monitoredMediaServerId)
		shell.Display(metadata, "MonitoredMediaServer", cmd.Flags())
	},
}

var PutMonitoredMediaServer = &cobra.Command{
	Use:   "monitoredMediaServer",
	Short: "Return a monitored media server for this organization. Only valid for organizations with Meraki Insight.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, _, _ := shell.ResolveFlags(cmd.Flags())
		if orgId == "" {
			orgId = args[1]
		}
		monitoredMediaServerId := args[0]
		var format configure.MonitoredMediaServer
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutMonitoredMediaServer(orgId, monitoredMediaServerId,  input)
		shell.Display(metadata, "MonitoredMediaServer", cmd.Flags())
	},
}

var PostMonitoredMediaServer = &cobra.Command{
	Use:   "monitoredMediaServer",
	Short: "Return a monitored media server for this organization. Only valid for organizations with Meraki Insight.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, _, _ := shell.ResolveFlags(cmd.Flags())
		if orgId == "" {
			orgId = args[0]
		}
		var format configure.MonitoredMediaServer
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostMonitoredMediaServer(orgId,  input)
		shell.Display(metadata, "MonitoredMediaServer", cmd.Flags())
	},
}