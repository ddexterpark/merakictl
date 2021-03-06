package sm

import (
	"github.com/ddexterpark/dashboard-api-golang/api/products/sm/configure"
	"github.com/ddexterpark/merakictl/shell"
	"github.com/spf13/cobra"
)


var GetApnsCert = &cobra.Command{
	Use:   "apnsCert",
	Short: "Get the organization's APNS certificate.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, _, _ := shell.ResolveFlags(cmd.Flags())
if orgId == "" {
	orgId = args[0]
}
		metadata := configure.GetAPNSCertificate(orgId)
		shell.Display(metadata, "ApnsCert", cmd.Flags())
	},
}

var GetBypassActivationLockAttempts = &cobra.Command{
Use:   "bypassActivationLockAttempts",
Short: "Bypass activation lock attempt status.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}

	attemptId := args[0]

metadata := configure.GetBypassActivationLockAttempts(networkId, attemptId)
shell.Display(metadata, "BypassActivationLockAttempts", cmd.Flags())
},
}

var PostBypassActivationLockAttempts = &cobra.Command{
	Use:   "bypassActivationLockAttempts",
	Short: "Bypass activation lock attempt status.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.BypassActivationLockAttempts
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostBypassActivationLockAttempts(networkId,  input)
		shell.Display(metadata, "BypassActivationLockAttempts", cmd.Flags())
	},
}

var GetCerts = &cobra.Command{
Use:   "certs",
Short: "List the certs on a device",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}

	deviceId := args[0]
metadata := configure.GetDeviceCerts(networkId, deviceId)
shell.Display(metadata, "Certs", cmd.Flags())
},
}


var GetDeviceProfiles = &cobra.Command{
Use:   "deviceProfiles",
Short: "Get the profiles associated with a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetDeviceProfiles(networkId, deviceId)
shell.Display(metadata, "DeviceProfiles", cmd.Flags())
},
}

var PutFields = &cobra.Command{
	Use:   "fields",
	Short: "Modify the fields of a device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.Fields
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PutFields(networkId,  input)
		shell.Display(metadata, "Fields", cmd.Flags())
	},
}

var GetNetworkAdapters = &cobra.Command{
Use:   "networkAdapters",
Short: "List the network adapters of a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetNetworkAdapters(networkId, deviceId)
shell.Display(metadata, "NetworkAdapters", cmd.Flags())
},
}


var GetRestrictions = &cobra.Command{
Use:   "restrictions",
Short: "List the restrictions on a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetDeviceRestrictions(networkId, deviceId)
shell.Display(metadata, "Restrictions", cmd.Flags())
},
}


var GetSecurityCenters = &cobra.Command{
Use:   "securityCenters",
Short: "List the security centers on a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetDeviceSecurityCenters(networkId, deviceId)
shell.Display(metadata, "SecurityCenters", cmd.Flags())
},
}


var GetDeviceSoftware = &cobra.Command{
Use:   "deviceSoftware",
Short: "Get a list of softwares associated with a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetDeviceSoftwares(networkId, deviceId)
shell.Display(metadata, "DeviceSoftware", cmd.Flags())
},
}


var GetWlanLists = &cobra.Command{
Use:   "wlanLists",
Short: "List the saved SSID names on a device.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	deviceId := args[0]
metadata := configure.GetDeviceWlanLists(networkId, deviceId)
shell.Display(metadata, "WlanLists", cmd.Flags())
},
}

var PostModifyTags = &cobra.Command{
	Use:   "modifyTags",
	Short: "Add, delete, or put the tags of a set of devices.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.ModifyTags
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostModifyTags(networkId,  input)
		shell.Display(metadata, "ModifyTags", cmd.Flags())
	},
}

var PostCheckin = &cobra.Command{
	Use:   "checkin",
	Short: "Force check-in a set of devices.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.SmIds
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostCheckin(networkId,  input)
		shell.Display(metadata, "Checkin", cmd.Flags())
	},
}

var GetDevices = &cobra.Command{
Use:   "devices",
Short: "List the devices enrolled in an SM network with various specified fields and filters.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}

	fields, _ := cmd.Flags().GetString("fields")
	wifiMacs, _ := cmd.Flags().GetString("wifiMacs")
	serials, _ := cmd.Flags().GetString("serials")
	ids, _ := cmd.Flags().GetString("ids")
	scope, _ := cmd.Flags().GetString("scope")
	perPage, _ := cmd.Flags().GetString("perPage")
	startingAfter, _ := cmd.Flags().GetString("startingAfter")
	endingBefore, _ := cmd.Flags().GetString("endingBefore")

metadata := configure.GetSMDevices(networkId,
	fields, wifiMacs, serials, ids, scope, perPage,
	startingAfter, endingBefore)
shell.Display(metadata, "Devices", cmd.Flags())
},
}

var PostLockDevices = &cobra.Command{
	Use:   "lockDevices",
	Short: "Lock a set of devices.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.SmIds
		input, _ := shell.ReadConfigFile(cmd, &format)

		metadata := configure.PostLock(networkId,  input)
		shell.Display(metadata, "LockDevices", cmd.Flags())
	},
}

var PostMoveDevices = &cobra.Command{
	Use:   "moveDevices",
	Short: "Move a set of devicesto a new network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.SmIds
		input, _ := shell.ReadConfigFile(cmd, &format)

		metadata := configure.PostMove(networkId,  input)
		shell.Display(metadata, "MoveDevices", cmd.Flags())
	},
}

var PostRefreshDevice = &cobra.Command{
	Use:   "refreshDevice",
	Short: "Refresh the details of a device",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		deviceId := args[0]

		metadata := configure.PostRefreshDetails(networkId, deviceId)
		shell.Display(metadata, "RefreshDevice", cmd.Flags())
	},
}

var PostUnEnrollDevice = &cobra.Command{
	Use:   "unEnrollDevice",
	Short: "UnEnroll a device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}
		var format configure.Unenroll
		deviceId := args[0]
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostUnenroll(networkId, deviceId,  input)
		shell.Display(metadata, "UnEnrollDevice", cmd.Flags())
	},
}

var PostWipeDevice = &cobra.Command{
	Use:   "wipeDevice",
	Short: "Wipe a device.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}
		var format configure.SmIds
		input, _ := shell.ReadConfigFile(cmd, &format)
		metadata := configure.PostWipeDevice(networkId,  input)
		shell.Display(metadata, "WipeDevice", cmd.Flags())
	},
}

var GetProfiles = &cobra.Command{
Use:   "profiles",
Short: "List all profiles in a network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}
metadata := configure.GetProfiles(networkId)
shell.Display(metadata, "Profiles", cmd.Flags())
},
}


var GetTargetGroups = &cobra.Command{
Use:   "targetGroups",
Short: "List the target groups in this network.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[0]
}

	withDetails, _ := cmd.Flags().GetString("withDetails")

metadata := configure.GetTargetGroups(networkId, withDetails)
shell.Display(metadata, "TargetGroups", cmd.Flags())
},
}

var PostTargetGroup = &cobra.Command{
	Use:   "targetGroup",
	Short: "Add a target group.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		withDetails, _ := cmd.Flags().GetString("withDetails")

		metadata := configure.PostTargetGroups(networkId, withDetails)
		shell.Display(metadata, "TargetGroup", cmd.Flags())
	},
}

var DelTargetGroup = &cobra.Command{
	Use:   "targetGroup",
	Short: "Delete a target group from a network.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[0]
		}

		withDetails, _ := cmd.Flags().GetString("withDetails")

		metadata := configure.DelTargetGroup(networkId, withDetails)
		shell.Display(metadata, "TargetGroup", cmd.Flags())
	},
}

var GetTargetGroup = &cobra.Command{
	Use:   "targetGroup",
	Short: "Return a target group.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		targetGroupId := args[0]
		withDetails, _ := cmd.Flags().GetString("withDetails")

		metadata := configure.GetTargetGroup(networkId, targetGroupId, withDetails)
		shell.Display(metadata, "TargetGroup", cmd.Flags())
	},
}

var PutTargetGroup = &cobra.Command{
	Use:   "targetGroup",
	Short: "Update a target group.",
	Run: func(cmd *cobra.Command, args []string) {
		_, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if networkId == "" {
			networkId = args[1]
		}

		targetGroupId := args[0]
		withDetails, _ := cmd.Flags().GetString("withDetails")

		metadata := configure.PutTargetGroup(networkId, targetGroupId, withDetails)
		shell.Display(metadata, "TargetGroup", cmd.Flags())
	},
}

var GetUserSoftware = &cobra.Command{
Use:   "userSoftware",
Short: "Get a list of softwares associated with a user.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}
	userId  := args[0]

metadata := configure.GetSmUserSoftware(networkId, userId)
shell.Display(metadata, "UserSoftware", cmd.Flags())
},
}


var GetUsers = &cobra.Command{
Use:   "users",
Short: "List the owners in an SM network with various specified fields and filters.",
Run: func(cmd *cobra.Command, args []string) {
_, networkId, _ := shell.ResolveFlags(cmd.Flags())
if networkId == "" {
	networkId = args[1]
}

	ids, _ := cmd.Flags().GetString("ids")
	usernames, _ := cmd.Flags().GetString("usernames")
	emails, _ := cmd.Flags().GetString("emails")
	scope, _ := cmd.Flags().GetString("scope")

metadata := configure.GetSmUsers(networkId, ids, usernames, emails, scope)
shell.Display(metadata, "Users", cmd.Flags())
},
}


var GetVPPAccounts = &cobra.Command{
Use:   "vppAccounts",
Short: "Get a hash containing the unparsed token of the VPP account with the given ID.",
Run: func(cmd *cobra.Command, args []string) {
orgId, _, _ := shell.ResolveFlags(cmd.Flags())
if orgId == "" {
	orgId = args[0]
}
metadata := configure.GetVPPAccounts(orgId)
shell.Display(metadata, "VPPAccounts", cmd.Flags())
},
}

var GetVPPAccount = &cobra.Command{
	Use:   "vppAccount",
	Short: "List the VPP accounts in the organization.",
	Run: func(cmd *cobra.Command, args []string) {
		orgId, _, _ := shell.ResolveFlags(cmd.Flags())
		if orgId == "" {
			orgId = args[1]
		}

		vppAccountId := args[0]
		metadata := configure.GetVPPAccount(orgId, vppAccountId)
		shell.Display(metadata, "VPPAccount", cmd.Flags())
	},
}