package network

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/networks/configure"
	organizations "github.com/ddexterpark/dashboard-api-golang/api/general/organizations/configure"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/spf13/cobra"
)

var GetAlertConfig = &cobra.Command{
	Use:   "AlertConfig",
	Short: "Return The Alert Configuration For This Network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetAlertSettings(networkId)
		shell.Display(metadata, "AlertConfig", cmd.Flags())
	},
}

var PutAlertConfig = &cobra.Command{
	Use:   "AlertConfig",
	Short: "Update the alert configuration for this network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		metadata := configure.PutAlertSettings(networkId, format)
		shell.Display(metadata, "AlertConfig", cmd.Flags())
	},
}


// clients
var GetClients = &cobra.Command{
	Use:   "Clients",
	Short: "List the clients that have used this network in the timespan.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		t0, _ := cmd.Flags().GetString("t0")
		t1, _ := cmd.Flags().GetString("t1")
		timespan, _ := cmd.Flags().GetString("timespan")
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")

		metadata := organizations.GetClients(networkId, t0, t1,
			timespan, perPage, startingAfter, endingBefore)
		shell.Display(metadata, "clients", cmd.Flags())
	},
}



var GetClientPolicy = &cobra.Command{
	Use:   "ClientPolicy",
	Short: "Return the policy assigned to a client on the network. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		clientId := args[0]
		metadata := configure.GetClientPolicy(networkId, clientId)
		shell.Display(metadata, "ClientPolicy", cmd.Flags())
	},
}

var GetClientSplashAuthorization = &cobra.Command{
	Use:   "ClientSplashAuthorization",
	Short: "Return the splash authorization for a client, for each SSID they've associated with through splash. Only enabled SSIDs with Click-through splash enabled will be included. Clients can be identified by a client key or either the MAC or IP depending on whether the network uses Track-by-IP.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		clientId := args[0]
		metadata := configure.GetSplashAuthorization(networkId, clientId)
		shell.Display(metadata, "ClientSplashAuthorization", cmd.Flags())
	},
}

var PutClientSplashAuthorization = &cobra.Command{
	Use:   "ClientSplashAuthorization",
	Short: "Update a client's splash authorization.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		clientId := args[0]

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		metadata := configure.PutSplashAuthorization(networkId, clientId, format)
		shell.Display(metadata, "ClientSplashAuthorization", cmd.Flags())
	},
}

var PostProvisionClient = &cobra.Command{
	Use:   "ProvisionClient",
	Short: "Update a client's splash authorization.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		metadata := configure.PostProvisionClient(networkId, format)
		shell.Display(metadata, "ProvisionClient", cmd.Flags())
	},
}



var GetDevices = &cobra.Command{
	Use:   "Devices",
	Short: "List the devices in a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetDevices(networkId)
		shell.Display(metadata, "Devices", cmd.Flags())
	},
}

var PostClaimSerials = &cobra.Command{
	Use:   "ClaimSerials",
	Short: "Claim devices into a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		format := args[0]
		if format == "" {
			// Read Config File
			var format interface{}
			shell.RenderInput(&format)
		}

		metadata := configure.PostClaimSerials(networkId, format)
		shell.Display(metadata, "ClaimSerials", cmd.Flags())
	},
}

var PostRemoveSerials = &cobra.Command{
	Use:   "RemoveSerials",
	Short: "Remove a single device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		format := args[0]
		if format == "" {
			// Read Config File
			var format interface{}
			shell.RenderInput(&format)
		}

		metadata := configure.PostUnClaimSerials(networkId, format)
		shell.Display(metadata, "RemoveSerials", cmd.Flags())
	},
}

var GetFirmwareUpgrades = &cobra.Command{
	Use:   "FirmwareUpgrades",
	Short: "Get current maintenance window for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetFirmwareUpgrades(networkId)
		shell.Display(metadata, "FirmwareUpgrades", cmd.Flags())
	},
}

// MISSING BODY PARAM in dashboard-api-golang
var PutFirmwareUpgrades = &cobra.Command{
	Use:   "FirmwareUpgrades",
	Short: "Update current maintenance window for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.PutFirmwareUpgrades(networkId)
		shell.Display(metadata, "FirmwareUpgrades", cmd.Flags())
	},
}

var GetFloorplans = &cobra.Command{
	Use:   "Floorplans",
	Short: "List the floor plans that belong to your network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetFloorPlans(networkId)
		shell.Display(metadata, "Floorplans", cmd.Flags())
	},
}

var GetFloorplan = &cobra.Command{
	Use:   "Floorplan",
	Short: "Find a floor plan by ID.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		floorplanId := args[0]
		metadata := configure.GetFloorPlan(networkId, floorplanId)
		shell.Display(metadata, "Floorplan", cmd.Flags())
	},
}

var DelFloorplan = &cobra.Command{
	Use:   "Floorplan",
	Short: "Delete a floor plan by ID.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		floorplanId := args[0]
		metadata := configure.DelFloorPlan(networkId, floorplanId)
		shell.Display(metadata, "Floorplan", cmd.Flags())
	},
}

var PutFloorplan = &cobra.Command{
	Use:   "Floorplan",
	Short: "Find a floor plan by ID.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		floorplanId := args[0]

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		metadata := configure.PutFloorPlan(networkId, floorplanId, format)
		shell.Display(metadata, "Floorplan", cmd.Flags())
	},
}

var PostFloorplan = &cobra.Command{
	Use:   "Floorplan",
	Short: "Find a floor plan by ID.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		metadata := configure.PostFloorPlan(networkId, format)
		shell.Display(metadata, "Floorplan", cmd.Flags())
	},
}

var GetGroupPolicies = &cobra.Command{
	Use:   "GroupPolicies",
	Short: "List the group policies in a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetGroupPolicies(networkId)
		shell.Display(metadata, "GroupPolicies", cmd.Flags())
	},
}

var GetGroupPolicy = &cobra.Command{
	Use:   "GroupPolicy ",
	Short: "Display a group policy.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		groupPolicyId := args[0]
		metadata := configure.GetGroupPolicy(networkId, groupPolicyId)
		shell.Display(metadata, "GroupPolicy", cmd.Flags())
	},
}

var DelGroupPolicy = &cobra.Command{
	Use:   "GroupPolicy ",
	Short: "Delete a group policy.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		groupPolicyId := args[0]
		metadata := configure.DelGroupPolicy(networkId, groupPolicyId)
		shell.Display(metadata, "GroupPolicy", cmd.Flags())
	},
}

var PutGroupPolicy = &cobra.Command{
	Use:   "GroupPolicy ",
	Short: "Update a group policy.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		groupPolicyId := args[0]

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		metadata := configure.PutGroupPolicy(networkId, groupPolicyId, format)
		shell.Display(metadata, "GroupPolicy", cmd.Flags())
	},
}

var PostGroupPolicy = &cobra.Command{
	Use:   "GroupPolicy ",
	Short: "Update a group policy.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		metadata := configure.PostGroupPolicy(networkId, format)
		shell.Display(metadata, "GroupPolicy", cmd.Flags())
	},
}

var GetMerakiAuthUsers = &cobra.Command{
	Use:   "MerakiAuthUsers",
	Short: "List the users configured under Meraki Authentication for a network (splash guest or RADIUS users for a mr network, or client VPN users for a wired network).",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetMerakiAuthUsers(networkId)
		shell.Display(metadata, "MerakiAuthUsers", cmd.Flags())
	},
}


var GetMerakiAuthUser = &cobra.Command{
	Use:   "MerakiAuthUser",
	Short: "Return the Meraki Auth splash guest, RADIUS, or client VPN user.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		merakiAuthUserId := args[0]
		metadata := configure.GetMerakiAuthUser(networkId, merakiAuthUserId)
		shell.Display(metadata, "MerakiAuthUser", cmd.Flags())
	},
}

var DelMerakiAuthUser = &cobra.Command{
	Use:   "MerakiAuthUser",
	Short: "Delete the Meraki Auth splash guest, RADIUS, or client VPN user.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		merakiAuthUserId := args[0]
		metadata := configure.DelMerakiAuthUser(networkId, merakiAuthUserId)
		shell.Display(metadata, "MerakiAuthUser", cmd.Flags())
	},
}

var PutMerakiAuthUser = &cobra.Command{
	Use:   "MerakiAuthUser",
	Short: "Update the Meraki Auth splash guest, RADIUS, or client VPN user.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		merakiAuthUserId := args[0]
		metadata := configure.PutMerakiAuthUser(networkId, merakiAuthUserId, format)
		shell.Display(metadata, "MerakiAuthUser", cmd.Flags())
	},
}

var PostMerakiAuthUser = &cobra.Command{
	Use:   "MerakiAuthUser",
	Short: "Create the Meraki Auth splash guest, RADIUS, or client VPN user.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostMerakiAuthUser(networkId, format)
		shell.Display(metadata, "MerakiAuthUser", cmd.Flags())
	},
}


var GetMQTTBrokers = &cobra.Command{
	Use:   "MQTTBrokers",
	Short: "List the MQTT brokers for this network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetMqttBrokers(networkId)
		shell.Display(metadata, "MQTTBrokers", cmd.Flags())
	},
}


var GetMQTTBroker = &cobra.Command{
	Use:   "MQTTBroker",
	Short: "Return an MQTT broker.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		mqttBrokerId := args[0]
		metadata := configure.GetMqttBroker(networkId, mqttBrokerId)
		shell.Display(metadata, "MQTTBroker", cmd.Flags())
	},
}

var DelMQTTBroker = &cobra.Command{
	Use:   "MQTTBroker",
	Short: "Delete a MQTT broker.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		mqttBrokerId := args[0]
		metadata := configure.DelMqttBroker(networkId, mqttBrokerId)
		shell.Display(metadata, "MQTTBroker", cmd.Flags())
	},
}

var PutMQTTBroker = &cobra.Command{
	Use:   "MQTTBroker",
	Short: "Return an MQTT broker.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		mqttBrokerId := args[0]

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		metadata := configure.PutMqttBroker(networkId, mqttBrokerId, format)
		shell.Display(metadata, "MQTTBroker", cmd.Flags())
	},
}

var PostMQTTBroker = &cobra.Command{
	Use:   "MQTTBroker",
	Short: "Create a MQTT broker.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		metadata := configure.PostMqttBroker(networkId, format)
		shell.Display(metadata, "MQTTBroker", cmd.Flags())
	},
}

var GetNetflow = &cobra.Command{
	Use:   "Netflow",
	Short: "Return the NetFlow traffic reporting settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetNetFlow(networkId)
		shell.Display(metadata, "Netflow", cmd.Flags())
	},
}

var PutNetflow = &cobra.Command{
	Use:   "Netflow",
	Short: "Return the NetFlow traffic reporting settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		metadata := configure.PutNetFlow(networkId, format)
		shell.Display(metadata, "Netflow", cmd.Flags())
	},
}

var GetChannelUtilization = &cobra.Command{
	Use:   "ChannelUtilization",
	Short: "Get the channel utilization over each radio for all APs in a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		t0, _ := cmd.Flags().GetString("t0")
		t1, _ := cmd.Flags().GetString("t1")
		timespan, _ := cmd.Flags().GetString("timespan")
		resolution, _ := cmd.Flags().GetString("resolution")
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")

		metadata := configure.GetChannelUtilization(networkId, t0, t1,
			timespan, resolution, perPage, startingAfter, endingBefore)
		shell.Display(metadata, "ChannelUtilization", cmd.Flags())
	},
}

var GetPIIRequests = &cobra.Command{
	Use:   "PIIRequests",
	Short: "List the PII requests for this network or organization.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetPiiRequests(networkId)
		shell.Display(metadata, "PIIRequests", cmd.Flags())
	},
}

var GetPIIRequest = &cobra.Command{
	Use:   "PIIRequest",
	Short: "Return a PII request.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		requestId := args[0]
		metadata := configure.GetPiiRequest(networkId, requestId)
		shell.Display(metadata, "PIIRequest", cmd.Flags())
	},
}

var DelPIIRequest = &cobra.Command{
	Use:   "PIIRequest",
	Short: "Delete a PII request.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		requestId := args[0]
		metadata := configure.DelPiiRequest(networkId, requestId)
		shell.Display(metadata, "PIIRequest", cmd.Flags())
	},
}

var PostPIIRequest = &cobra.Command{
	Use:   "PIIRequest",
	Short: "Create a PII request.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostPiiRequest(networkId, format)
		shell.Display(metadata, "PIIRequest", cmd.Flags())
	},
}

var GetSMDevices = &cobra.Command{
	Use:   "SMDevices",
	Short: "Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		username, _ := cmd.Flags().GetString("username")
		email, _ := cmd.Flags().GetString("email")
		mac, _ := cmd.Flags().GetString("mac")
		serial, _ := cmd.Flags().GetString("serial")
		imei, _ := cmd.Flags().GetString("imei")
		bluetoothMac, _ := cmd.Flags().GetString("bluetoothMac")

		metadata := configure.GetSmDevicesForKey(networkId,
			username, email, mac, serial, imei, bluetoothMac)
		shell.Display(metadata, "SMDevices", cmd.Flags())
	},
}

var GetSMOwners = &cobra.Command{
	Use:   "SMOwners",
	Short: "Given a piece of Personally Identifiable Information (PII), return the Systems Manager device ID(s) associated with that identifier.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		username, _ := cmd.Flags().GetString("username")
		email, _ := cmd.Flags().GetString("email")
		mac, _ := cmd.Flags().GetString("mac")
		serial, _ := cmd.Flags().GetString("serial")
		imei, _ := cmd.Flags().GetString("imei")
		bluetoothMac, _ := cmd.Flags().GetString("bluetoothMac")

		metadata := configure.GetSmOwnersForKey(networkId,
			username, email, mac, serial, imei, bluetoothMac)
		shell.Display(metadata, "SMOwners", cmd.Flags())
	},
}


var GetSettings = &cobra.Command{
	Use:   "Settings",
	Short: "Return the settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetSettings(networkId)
		shell.Display(metadata, "Settings", cmd.Flags())
	},
}

var PutSettings = &cobra.Command{
	Use:   "Settings",
	Short: "Update the settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutSettings(networkId, format)
		shell.Display(metadata, "Settings", cmd.Flags())
	},
}

var GetSNMP = &cobra.Command{
	Use:   "SNMP",
	Short: "Return the SNMP settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetSNMP(networkId)
		shell.Display(metadata, "SNMP", cmd.Flags())
	},
}

// FIX in Dashboard-API-GoLang
var PutSNMP = &cobra.Command{
	Use:   "SNMP",
	Short: "Return the SNMP settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		//metadata := configure.PutSNMP(networkId, format)
		//shell.Display(metadata, "SNMP", cmd.Flags())
	},
}

var GetSyslog = &cobra.Command{
	Use:   "Syslog",
	Short: "List the syslog servers for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetSyslogServers(networkId)
		shell.Display(metadata, "Syslog", cmd.Flags())
	},
}

var PutSyslog = &cobra.Command{
	Use:   "Syslog",
	Short: "Update the syslog servers for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutSyslogServers(networkId, format)
		shell.Display(metadata, "Syslog", cmd.Flags())
	},
}

var GetTrafficAnalysis = &cobra.Command{
	Use:   "TrafficAnalysis",
	Short: "Return the traffic analysis settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetTrafficAnalysis(networkId)
		shell.Display(metadata, "TrafficAnalysis", cmd.Flags())
	},
}

var PutTrafficAnalysis = &cobra.Command{
	Use:   "TrafficAnalysis",
	Short: "Update the traffic analysis settings for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutTrafficAnalysis(networkId, format)
		shell.Display(metadata, "TrafficAnalysis", cmd.Flags())
	},
}

var GetTrafficShaping = &cobra.Command{
	Use:   "TrafficShaping",
	Short: "Returns the application categories for traffic shaping rules.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetTrafficAnalysis(networkId)
		shell.Display(metadata, "TrafficShaping", cmd.Flags())
	},
}

var GetDSCP = &cobra.Command{
	Use:   "DSCP",
	Short: "Returns the available DSCP tagging options for your traffic shaping rules.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetDscpTaggingOptions(networkId)
		shell.Display(metadata, "DSCP", cmd.Flags())
	},
}

var GetHTTPServers = &cobra.Command{
	Use:   "HTTPServers",
	Short: "List the HTTP servers for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		metadata := configure.GetHTTPServers(networkId)
		shell.Display(metadata, "HTTPServers", cmd.Flags())
	},
}

var GetHTTPServer = &cobra.Command{
	Use:   "HTTPServer",
	Short: "Return an HTTP server for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		httpServerId := args[0]
		metadata := configure.GetHTTPServer(networkId, httpServerId)
		shell.Display(metadata, "HTTPServer", cmd.Flags())
	},
}

var PutHTTPServer = &cobra.Command{
	Use:   "HTTPServer",
	Short: "Update an HTTP server for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		httpServerId := args[0]
		metadata := configure.PutHTTPServer(networkId, httpServerId, format)
		shell.Display(metadata, "HTTPServer", cmd.Flags())
	},
}

// FIX in DASHBOARD-API-GOLANG
var PostHTTPServer = &cobra.Command{
	Use:   "HTTPServer",
	Short: "Create a HTTP server for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)

		//metadata := configure.PostHTTPServer(networkId,format)
		//shell.Display(metadata, "HTTPServer", cmd.Flags())
	},
}

var DelHTTPServer = &cobra.Command{
	Use:   "HTTPServer",
	Short: "Delete an HTTP server for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		httpServerId := args[0]
		metadata := configure.DelHTTPServer(networkId, httpServerId)
		shell.Display(metadata, "HTTPServer", cmd.Flags())
	},
}

var GetWebhookTest = &cobra.Command{
	Use:   "WebhookTest",
	Short: "Return the status of a webhook test for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		weebhookTestId := args[0]
		metadata := configure.GetWebhookTests(networkId, weebhookTestId)
		shell.Display(metadata, "WebhookTest", cmd.Flags())
	},
}

var PostWebhookTest = &cobra.Command{
	Use:   "WebhookTest",
	Short: "Post  a webhook test for a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.PostWebhookTests(networkId)
		shell.Display(metadata, "WebhookTest", cmd.Flags())
	},
}

var GetNetwork = &cobra.Command{
	Use:   "Details",
	Short: "Return a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.GetNetwork(networkId)
		shell.Display(metadata, "Details", cmd.Flags())
	},
}

var PutNetwork = &cobra.Command{
	Use:   "Details",
	Short: "Return a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutNetwork(networkId, format)
		shell.Display(metadata, "Details", cmd.Flags())
	},
}

var DelNetwork = &cobra.Command{
	Use:   "Details",
	Short: "Delete a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		metadata := configure.DelNetwork(networkId)
		shell.Display(metadata, "Details", cmd.Flags())
	},
}

//FIX IN DASHBOARD-API-GOLANG
var PostBindNetwork = &cobra.Command{
	Use:   "BindNetwork",
	Short: "Bind a network to a template.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		//metadata := configure.PostBindNetwork()
		//shell.Display(metadata, "BindNetwork", cmd.Flags())
	},
}

//FIX IN DASHBOARD-API-GOLANG
var PostUnBindNetwork = &cobra.Command{
	Use:   "UnBindNetwork",
	Short: "UnBind a network from a template.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		//metadata := configure.PostUnBindNetwork()
		//shell.Display(metadata, "UnBindNetwork", cmd.Flags())
	},
}

//FIX IN DASHBOARD-API-GOLANG
var PostSplitNetwork = &cobra.Command{
	Use:   "SplitNetwork",
	Short: "Split a network from a template.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		//metadata := configure.PostSplitNetwork()
		//shell.Display(metadata, "SplitNetwork", cmd.Flags())
	},
}