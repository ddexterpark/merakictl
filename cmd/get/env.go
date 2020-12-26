package get

import (
	"github.com/ddexterpark/dashboard-api-golang/api"
	"github.com/ddexterpark/dashboard-api-golang/shell"
	"github.com/kr/pretty"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// apikey -> api.Token()
var environmentalvariables = &cobra.Command{
	Use:   "env",
	Short: "Prints out environmental variables",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		api.EnvironmentalVariables()

		apiKey := viper.GetString("MERAKI_DASHBOARD_API_KEY")
		baseUrl :=  viper.GetString("MERAKI_DASHBOARD_API_URL")
		version :=  viper.GetString("MERAKI_DASHBOARD_API_VERSION")

		obfuscated := string(apiKey[len(apiKey)-6:])
		pretty.Println("API KEY is : *****************"+obfuscated)
		pretty.Println("API Version is :", version)
		pretty.Println("API BaseUrl is :", baseUrl)


	},
}

var test = &cobra.Command{
	Use:   "test",
	Short: "test stuff.",
	Long:  pretty.Sprint("test"),
	Run: func(cmd *cobra.Command, args []string) {
		pretty.Println(shell.ResolveFlags(cmd.Flags()))
	},
}