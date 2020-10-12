package configure

import (
	"fmt"
	"github.com/ddexterpark/merakictl/api"
	user_agent "github.com/ddexterpark/merakictl/user-agent"
	"io"
)

type ClientPolicy struct {
	Mac           string `json:"mac"`
	DevicePolicy  string `json:"devicePolicy"`
	GroupPolicyID string `json:"groupPolicyId"`
}

// GetClientPolicy -  Return The Policy Assigned To A Client On The Network
func GetClientPolicy(networkId, clientId string) (ClientPolicy, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/clients/%s/policy", api.BaseUrl(), networkId, clientId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = ClientPolicy{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}

type SplashAuthorization struct {
	Ssids struct {
		Num0 struct {
			IsAuthorized bool   `json:"isAuthorized"`
			AuthorizedAt string `json:"authorizedAt"`
			ExpiresAt    string `json:"expiresAt"`
		} `json:"0"`
		Num2 struct {
			IsAuthorized bool `json:"isAuthorized"`
		} `json:"2"`
	} `json:"ssids"`
}

// Return The Splash Authorization For A Client For Each SSID They've Associated With Through Splash
func GetSplashAuthorization(networkId, clientId string) (SplashAuthorization, interface{}) {
	baseurl := fmt.Sprintf("%s/networks/%s/clients/%s/splashAuthorizationStatus", api.BaseUrl(), networkId, clientId)
	var payload io.ReadSeeker
	session := api.Session(baseurl, "GET", payload)
	var results = SplashAuthorization{}
	user_agent.UnMarshalJSON(session.Body, &results)
	traceback := user_agent.TraceBack(session)
	return results, traceback
}