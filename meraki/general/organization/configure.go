package organization

import (
	"github.com/ddexterpark/dashboard-api-golang/api/general/organizations/configure"
	shell "github.com/ddexterpark/merakictl/utilities"
	"github.com/spf13/cobra"
)



var GetActionBatches = &cobra.Command{
	Use:   "ActionBatches",
	Short: "Return The List Of Action Batches In The Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		status := cmd.Flag("status").Value.String()
		metadata := configure.GetActionBatches(org, status)
		shell.Display(metadata, "ActionBatches", cmd.Flags())

	},
}

var GetActionBatch = &cobra.Command{
	Use:   "ActionBatch",
	Short: "Return The List Of Action Batches In The Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		actionBatchId := args[0]
		metadata := configure.GetActionBatch(org, actionBatchId)
		shell.Display(metadata, "ActionBatch", cmd.Flags())
	},
}

var DelActionBatch = &cobra.Command{
	Use:   "ActionBatch",
	Short: "Delete Action Batches In The Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		actionBatchId := args[0]
		metadata := configure.DelActionBatch(org, actionBatchId)
		shell.Display(metadata, "ActionBatch", cmd.Flags())
	},
}

var PutActionBatch = &cobra.Command{
	Use:   "ActionBatch",
	Short: "Update The List Of Action Batches In The Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		actionBatchId := args[0]
		metadata := configure.PutActionBatch(org, actionBatchId, format)
		shell.Display(metadata, "ActionBatch", cmd.Flags())
	},
}

var PostActionBatch = &cobra.Command{
	Use:   "ActionBatch",
	Short: "Create an Action Batch.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostActionBatch(org, format)
		shell.Display(metadata, "ActionBatch", cmd.Flags())
	},
}

var GetAdmins = &cobra.Command{
	Use:   "Admins",
	Short: "Display a list of dashboard administrators.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := configure.GetAdmins(org)
		shell.Display(metadata, "Admins", cmd.Flags())

	},
}

var DelAdmins = &cobra.Command{
	Use:   "Admins",
	Short: "Delete a dashboard administrator.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		adminId := args[0]
		metadata := configure.DelAdmin(org, adminId)
		shell.Display(metadata, "Admins", cmd.Flags())

	},
}

var PostAdmins = &cobra.Command{
	Use:   "Admins",
	Short: "Create a new dashboard administrator.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		adminId := args[0]
		metadata := configure.PostAdmin(org, adminId)
		shell.Display(metadata, "Admins", cmd.Flags())

	},
}

var PutAdmins = &cobra.Command{
	Use:   "Admins",
	Short: "Create a new dashboard administrator.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		adminId := args[0]
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutAdmin(org, adminId, format)
		shell.Display(metadata, "Admins", cmd.Flags())

	},
}

var GetBrandingPolicies = &cobra.Command{
	Use:   "BrandingPolicies",
	Short: "Return The Branding Policy IDs Of An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := configure.GetBrandingPolicies(org)
		shell.Display(metadata, "BrandingPolicies", cmd.Flags())

	},
}

var PutBrandingPolicyPriorities = &cobra.Command{
	Use:   "BrandingPolicyPriorities",
	Short: "Update the priority ordering of an organization's branding policies.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutBrandingPolicyPriorities(org, format)
		shell.Display(metadata, "BrandingPolicyPriorities", cmd.Flags())

	},
}

var GetBrandingPolicy = &cobra.Command{
	Use:   "BrandingPolicy",
	Short: "Return The Branding Policy IDs Of An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		brandingId := args[0]
		metadata := configure.GetBrandingPolicy(org, brandingId)
		shell.Display(metadata, "BrandingPolicy", cmd.Flags())
	},
}

var DelBrandingPolicy = &cobra.Command{
	Use:   "BrandingPolicy",
	Short: "Delete The Branding Policy IDs Of An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		brandingId := args[0]
		metadata := configure.DelBrandingPolicy(org, brandingId)
		shell.Display(metadata, "BrandingPolicy", cmd.Flags())
	},
}

var PutBrandingPolicy = &cobra.Command{
	Use:   "BrandingPolicy",
	Short: "Return The Branding Policy IDs Of An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		brandingId := args[0]
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutBrandingPolicy(org, brandingId, format)
		shell.Display(metadata, "BrandingPolicy", cmd.Flags())
	},
}

var PostBrandingPolicy = &cobra.Command{
	Use:   "BrandingPolicy",
	Short: "Return The Branding Policy IDs Of An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostBrandingPolicy(org, format)
		shell.Display(metadata, "BrandingPolicy", cmd.Flags())
	},
}

var GetConfigurationTemplates = &cobra.Command{
Use:   "ConfigurationTemplates",
Short: "List The Configuration Templates For This Organization.",
Run: func(cmd *cobra.Command, args []string) {
org, _, _ := shell.ResolveFlags(cmd.Flags())
if org == "" {
org = args[0]
}
metadata := configure.GetConfigTemplates(org)
shell.Display(metadata, "ConfigurationTemplates", cmd.Flags())
},
}

var GetConfigurationTemplate = &cobra.Command{
Use:   "ConfigurationTemplate",
Short: "Return a Configuration Template For This Organization.",
Run: func(cmd *cobra.Command, args []string) {
org, _, _ := shell.ResolveFlags(cmd.Flags())
if org == "" {
org = args[1]
}
configTemplateId := args[0]
metadata := configure.GetConfigTemplate(org, configTemplateId)
shell.Display(metadata, "ConfigurationTemplate", cmd.Flags())
},
}

var DelConfigurationTemplate = &cobra.Command{
	Use:   "ConfigurationTemplate",
	Short: "Delete a Configuration Template For This Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		configTemplateId := args[0]
		metadata := configure.DelConfigTemplate(org, configTemplateId)
		shell.Display(metadata, "ConfigurationTemplate", cmd.Flags())
	},
}

var PutConfigurationTemplate = &cobra.Command{
	Use:   "ConfigurationTemplate",
	Short: "Update a Configuration Template For This Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		configTemplateId := args[0]

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutConfigTemplate(org, configTemplateId, format)
		shell.Display(metadata, "ConfigurationTemplate", cmd.Flags())
	},
}

var PostConfigurationTemplate = &cobra.Command{
	Use:   "ConfigurationTemplate",
	Short: "Create a Configuration Template For This Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostConfigTemplate(org, format)
		shell.Display(metadata, "ConfigurationTemplate", cmd.Flags())
	},
}

var GetDevices = &cobra.Command{
	Use:   "Devices",
	Short: "List the devices in an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		configurationUpdatedAfter, _ := cmd.Flags().GetString("configurationUpdatedAfter")
		metadata := configure.GetOrganizationDevices(org, perPage, startingAfter, configurationUpdatedAfter)
		shell.Display(metadata, "Devices", cmd.Flags())
	},
}


var GetInventoryDevices = &cobra.Command{
	Use:   "InventoryDevices",
	Short: "Return The Device Inventory For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")
		usedState, _ := cmd.Flags().GetString("usedState")
		search, _ := cmd.Flags().GetString("search")

		metadata := configure.GetInventoryDevices(org, perPage, startingAfter, endingBefore, usedState, search)
		shell.Display(metadata, "InventoryDevices", cmd.Flags())
	},
}

var GetInventoryDevice = &cobra.Command{
	Use:   "InventoryDevice",
	Short: "Return The Device Inventory For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}

		deviceId := args[0]

		metadata := configure.GetInventoryDevice(org, deviceId)
		shell.Display(metadata, "InventoryDevice", cmd.Flags())
	},
}

var GetLicences = &cobra.Command{
	Use:   "Licences",
	Short: "List The Licenses For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, networkId, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		if networkId == "" {
			networkId = args[1]
		}

		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")
		deviceSerial, _ := cmd.Flags().GetString("deviceSerial")
		state, _ := cmd.Flags().GetString("state")

		metadata := configure.GetLicenses(org, perPage,
			startingAfter, endingBefore, deviceSerial, networkId, state)
			shell.Display(metadata, "Licences", cmd.Flags())
	},
}

var PostAssignSeats = &cobra.Command{
	Use:   "AssignSeats",
	Short: "List A Single License For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, serial := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		if serial == "" {
			serial = args[1]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostAssignSeats(org, serial, format)
		shell.Display(metadata, "AssignSeats", cmd.Flags())
	},
}

var PostMoveSeats = &cobra.Command{
	Use:   "MoveSeats",
	Short: "List A Single License For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, serial := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		if serial == "" {
			serial = args[1]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostMoveSeats(org, serial, format)
		shell.Display(metadata, "MoveSeats", cmd.Flags())
	},
}

var PostMoveLicenses = &cobra.Command{
	Use:   "MoveLicenses",
	Short: "List A Single License For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, serial := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		if serial == "" {
			serial = args[1]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostMoveLicenses(org, serial, format)
		shell.Display(metadata, "MoveLicenses", cmd.Flags())
	},
}

var PostRenewSeats = &cobra.Command{
	Use:   "RenewSeats",
	Short: "List A Single License For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, serial := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		if serial == "" {
			serial = args[1]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostRenewSeats(org, serial, format)
		shell.Display(metadata, "RenewSeats", cmd.Flags())
	},
}

var GetLicence = &cobra.Command{
	Use:   "Licence",
	Short: "List A Single License For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		licenceId, _ := cmd.Flags().GetString("licenceId")
		metadata := configure.GetLicense(org, licenceId)
		shell.Display(metadata, "Licence", cmd.Flags())
	},
}

var PutLicence = &cobra.Command{
	Use:   "Licence",
	Short: "Update a license.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		licenceId, _ := cmd.Flags().GetString("licenceId")
		metadata := configure.PutLicense(org, licenceId)
		shell.Display(metadata, "Licence", cmd.Flags())
	},
}

var GetLoginSecurity = &cobra.Command{
	Use:   "loginSecurity",
	Short: "Returns The Login Security Settings For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		metadata := configure.GetLoginSecurity(org)
		shell.Display(metadata, "LoginSecurity", cmd.Flags())
	},
}

var PutLoginSecurity = &cobra.Command{
	Use:   "loginSecurity",
	Short: "Update The Login Security Settings For An Organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutLoginSecurity(org, format)
		shell.Display(metadata, "LoginSecurity", cmd.Flags())
	},
}

var GetNetworks = &cobra.Command{
	Use:   "Networks",
	Short: "List the networks that the user has privileges on in an organization.",
	Run: func(cmd *cobra.Command, args []string) {

		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		perPage, _ := cmd.Flags().GetString("perPage")
		startingAfter, _ := cmd.Flags().GetString("startingAfter")
		endingBefore, _ := cmd.Flags().GetString("endingBefore")

		configTemplateId, _ := cmd.Flags().GetString("configTemplateId")
		tagsFilterType, _ := cmd.Flags().GetString("tagsFilterType")
		tags, _ := cmd.Flags().GetString("tags")

		metadata := configure.GetNetworks(org, configTemplateId,
			tagsFilterType, startingAfter, endingBefore, tags, perPage)

		shell.Display(metadata, "Networks", cmd.Flags())

	},
}


var PostNetworks = &cobra.Command{
	Use:   "Networks",
	Short: "List the networks that the user has privileges on in an organization.",
	Run: func(cmd *cobra.Command, args []string) {

		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostNetworks(org, format)

		shell.Display(metadata, "Networks", cmd.Flags())

	},
}

var PostCombineNetworks = &cobra.Command{
	Use:   "CombineNetworks",
	Short: "Combine multiple networks into a single network.",
	Run: func(cmd *cobra.Command, args []string) {

		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostCombineNetworks(org, format)

		shell.Display(metadata, "CombineNetworks", cmd.Flags())

	},
}

var GetIDPS = &cobra.Command{
	Use:   "IDPS",
	Short: "List the SAML IdPs in your organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		metadata := configure.GetIDPS(org)
		shell.Display(metadata, "IDPS", cmd.Flags())
	},
}

var GetIDP = &cobra.Command{
	Use:   "IDP",
	Short: "List a SAML IdP in your organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}

		ldp := args[0]
		metadata := configure.GetIDP(org, ldp)
		shell.Display(metadata, "IDP", cmd.Flags())
	},
}

var DelIDP = &cobra.Command{
	Use:   "IDP",
	Short: "List a SAML IdP in your organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}

		ldp := args[0]
		metadata := configure.DelIDP(org, ldp)
		shell.Display(metadata, "IDP", cmd.Flags())
	},
}

var PutIDP = &cobra.Command{
	Use:   "IDP",
	Short: "List a SAML IdP in your organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}

		ldp := args[0]
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutIDP(org, ldp, format)
		shell.Display(metadata, "IDP", cmd.Flags())
	},
}

var PostIDP = &cobra.Command{
	Use:   "IDP",
	Short: "List a SAML IdP in your organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostIDP(org, format)
		shell.Display(metadata, "IDP", cmd.Flags())
	},
}

var GetSAML = &cobra.Command{
	Use:   "SAML",
	Short: "Returns the SAML SSO enabled settings for an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		metadata := configure.GetSAML(org)
		shell.Display(metadata, "SAML", cmd.Flags())
	},
}

var PutSAML = &cobra.Command{
	Use:   "SAML",
	Short: "Returns the SAML SSO enabled settings for an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutSAML(org, format)
		shell.Display(metadata, "SAML", cmd.Flags())
	},
}

var GetSamlRoles = &cobra.Command{
	Use:   "SamlRoles",
	Short: "List the SAML roles for this organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}

		metadata := configure.GetSAMLRoles(org)
		shell.Display(metadata, "SamlRoles", cmd.Flags())
	},
}

var GetSamlRole = &cobra.Command{
	Use:   "SamlRole",
	Short: "List a single SAML role for this organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		samlRoleId := args[0]
		metadata := configure.GetSAMLRole(org, samlRoleId)
		shell.Display(metadata, "SamlRole", cmd.Flags())
	},
}

var DelSamlRole = &cobra.Command{
	Use:   "SamlRole",
	Short: "List a single SAML role for this organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		samlRoleId := args[0]
		metadata := configure.DelSAMLRole(org, samlRoleId)
		shell.Display(metadata, "SamlRole", cmd.Flags())
	},
}

var PutSamlRole = &cobra.Command{
	Use:   "SamlRole",
	Short: "List a single SAML role for this organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		samlRoleId := args[0]
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutSAMLRole(org, samlRoleId, format)
		shell.Display(metadata, "SamlRole", cmd.Flags())
	},
}

var PostSamlRole = &cobra.Command{
	Use:   "SamlRole",
	Short: "List a single SAML role for this organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostSAMLRole(org, format)
		shell.Display(metadata, "SamlRole", cmd.Flags())
	},
}

var GetSNMP = &cobra.Command{
	Use:   "SNMP",
	Short: "Return the SNMP settings for an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := configure.GetSNMP(org)
		shell.Display(metadata, "SNMP", cmd.Flags())
	},
}

var PutSNMP = &cobra.Command{
	Use:   "SNMP",
	Short: "Update the SNMP settings for an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PutSNMP(org, format)
		shell.Display(metadata, "SNMP", cmd.Flags())
	},
}

var GetList = &cobra.Command{
	Use:   "List",
	Short: "List the organizations that the user has privileges on.",
	Run: func(cmd *cobra.Command, args []string) {
		metadata := configure.GetOrganizations()
		shell.Display(metadata, "organizations", cmd.Flags())
	},
}

var GetDetail = &cobra.Command{
	Use:   "Details",
	Short: "List a specific organization that the user has privileges on.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := configure.GetOrganization(org)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}


var DelOrganization = &cobra.Command{
	Use:   "fromList",
	Short: "Delete an existing organization",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		metadata := configure.DelOrganization(org)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}

var PutOrganization = &cobra.Command{
	Use:   "org",
	Short: "Update an existing organization",
	Run: func(cmd *cobra.Command, args []string) {

		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[1]
		}
		var name = args[0]

		metadata := configure.PutOrganization(org, name)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}

var PostOrganization = &cobra.Command{
	Use:   "org",
	Short: "Create an organization",
	Run: func(cmd *cobra.Command, args []string) {
		var name = args[0]
		metadata := configure.PostOrganization(name)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}

var PostClaim = &cobra.Command{
	Use:   "org",
	Short: "Claim a list of devices, licenses, and/or orders into an organization.",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostClaim(org, format)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}

var PostClone = &cobra.Command{
	Use:   "org",
	Short: "Create a new organization by cloning the addressed organization",
	Run: func(cmd *cobra.Command, args []string) {
		org, _, _ := shell.ResolveFlags(cmd.Flags())
		if org == "" {
			org = args[0]
		}
		// Read Config File
		var format interface{}
		shell.RenderInput(&format)
		metadata := configure.PostClone(org, format)
		shell.Display(metadata, "organization", cmd.Flags())
	},
}