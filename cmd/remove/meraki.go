package remove


import (
	"github.com/ddexterpark/merakictl/meraki/products/mv"
	"github.com/spf13/cobra"
	//"github.com/ddexterpark/merakictl/meraki/general/device"
	"github.com/ddexterpark/merakictl/meraki/general/network"
	"github.com/ddexterpark/merakictl/meraki/general/organization"
	//"github.com/ddexterpark/merakictl/meraki/products/mx"
	//"github.com/ddexterpark/merakictl/meraki/products/mv"
	//"github.com/ddexterpark/merakictl/meraki/products/mg"
	"github.com/ddexterpark/merakictl/meraki/products/insight"
	//"github.com/ddexterpark/merakictl/meraki/products/sm"
	//"github.com/ddexterpark/merakictl/meraki/products/ms"
	//"github.com/ddexterpark/merakictl/meraki/products/mr"
)

var OrgCmd = &cobra.Command{
	Use:   "organization",
	Aliases: []string{"org"},
	Short: "Root for all organization CLI commands.",
}

func init() {
	OrgCmd.AddCommand(organization.DelOrganization)
	OrgCmd.AddCommand(organization.DelActionBatch)
	OrgCmd.AddCommand(organization.DelAdmins)
	OrgCmd.AddCommand(organization.DelBrandingPolicy)
	OrgCmd.AddCommand(organization.DelConfigurationTemplate)
	OrgCmd.AddCommand(organization.DelIDP)
	OrgCmd.AddCommand(organization.DelSamlRole)
}

var NetCmd = &cobra.Command{
	Use:     "network",
	Aliases: []string{"net"},
	Short:   "Entrypoint for general network subcommands.",
}

func init() {
	NetCmd.AddCommand(network.DelFloorplan)
	NetCmd.AddCommand(network.DelGroupPolicy)
	NetCmd.AddCommand(network.DelHTTPServer)
	NetCmd.AddCommand(network.DelMerakiAuthUser)
	NetCmd.AddCommand(network.DelMQTTBroker)
	NetCmd.AddCommand(network.DelNetwork)
	NetCmd.AddCommand(network.DelPIIRequest)

}

var DeviceCmd = &cobra.Command{
	Use:   "device",
	Aliases: []string{"dev"},
	Short: "Root for general device CLI commands.",
	Long:  "Entrypoint for Device subcommands.",
}

func init() {
	//DeviceCmd.AddCommand()
}

var MXCmd = &cobra.Command{
	Use:   "appliance",
	Aliases: []string{"mx"},
	Short: "Root for all mx appliance CLI commands.",
	Long:  "Entrypoint for mx subcommands.",
}

func init() {
	MXCmd.AddCommand()
}

var MSCmd = &cobra.Command{
	Use:   "switch",
	Aliases: []string{"ms"},
	Short: "Root for all Switch CLI commands.",
}

func init() {
	MSCmd.AddCommand()
}

var MRCmd = &cobra.Command{
	Use:   "wireless",
	Aliases: []string{"mr"},
	Short: "Root for all wireless MR CLI commands.",
}

func init() {
	MRCmd.AddCommand()
}

var MGCmd = &cobra.Command{
	Use:   "gateway",
	Aliases: []string{"mg"},
	Short: "Entrypoint for mg cellular gateway subcommands.",
}

func init() {
	MGCmd.AddCommand()
}


var MVCmd = &cobra.Command{
	Use:   "camera",
	Aliases: []string{"mv"},
	Short:  "Entrypoint for mv camera subcommands.",
}

func init() {
	MVCmd.AddCommand(mv.DelQualityRetentionProfile)
}

var SMCmd = &cobra.Command{
	Use:   "systems",
	Aliases: []string{"sm"},
	Short: "Entrypoint for SM subcommands.",
}

func init() {
	SMCmd.AddCommand()
}

var InsightCmd = &cobra.Command{
	Use:   "insight",
	Aliases: []string{"in"},
	Short: "Entrypoint for Insight subcommands.",
}

func init() {
	InsightCmd.AddCommand(insight.DelMonitoredMediaServer)

}