# Merakictl CLI Reference 

*Main Documentation: [English](https://github.com/ddexterpark/merakictl/blob/master/README.md)*

*CLI Reference Overview: [English](https://github.com/ddexterpark/merakictl/tree/master/docs/commands/README.md)*

## SM  
 
 All SM level API calls. 
 
  HTTP | Operation | Syntax | Filters | Description |
 ----- | --------- | ------ | ----------- | ----------- |
  GET  | apnsCert | merakictl get sm apnsCert {organizationId} | | Get the organization's APNS certificate.
  GET  | bypassActivationLockAttempts | merakictl get sm bypassActivationLockAttempts {attemptId} {networkId} | | Bypass activation lock attempt status.
  GET  | certs | merakictl get sm certs {deviceId} {networkId} | | List the certs on a device.
  GET  | deviceProfiles | merakictl get sm deviceProfiles {deviceId} {networkId} | | Get the profiles associated with a device.
  GET  | networkAdapters | merakictl get sm networkAdapters {deviceId} {networkId} | | List the network adapters of a device.
  GET  | restrictions | merakictl get sm restrictions {deviceId} {networkId} | | List the restrictions on a device.
  GET  | securityCenters | merakictl get sm securityCenters {deviceId} {networkId} | | List the security centers on a device.
  GET  | deviceSoftware | merakictl get sm deviceSoftware {deviceId} {networkId} | | Get a list of softwares associated with a device.
  GET  | wlanLists | merakictl get sm wlanLists {deviceId} {networkId} | | List the saved SSID names on a device.
  GET  | devices | merakictl get sm devices {networkId} | --fields --wifiMacs --serials --ids --scope --perPage --startingAfter --endingBefore | List the devices enrolled in an SM network with various specified fields and filters.
  GET  | profiles | merakictl get sm profiles {networkId} | | List all profiles in a network.
  GET  | targetGroups | merakictl get sm targetGroups {networkId} | --withDetails | List the target groups in this network.
  GET  | targetGroup | merakictl get sm targetGroup {targetGroupId} {networkId} | --withDetails | Return a target group.
  GET  | userSoftware | merakictl get sm userSoftware {clientId} {networkId} | | Get a list of softwares associated with a user.
  GET  | users | merakictl get sm users {networkId} | --ids --usernames --emails --scope | List the owners in an SM network with various specified fields and filters.
  GET  | vppAccount | merakictl get sm vppAccount {vppAccountId} {organizationId} | | Get a hash containing the unparsed token of the VPP account with the given ID.
  GET  | vppAccounts | merakictl get sm vppAccounts {organizationId} | | List the VPP accounts in the organization.
  GET  | cellularUsageHistory | merakictl get sm cellularUsageHistory {deviceId} {networkId} | | Return the client's daily cellular data usage history. Usage data is in kilobytes.
  GET  | connectivity | merakictl get sm connectivity {deviceId} {networkId} | --perPage --startingAfter --endingBefore | Returns historical connectivity data (whether a device is regularly checking in to Dashboard).
  GET  | desktopLogs | merakictl get sm desktopLogs {deviceId} {networkId} | --perPage --startingAfter --endingBefore | Return historical records of various Systems Manager network connection details for desktop devices.
  GET  | deviceCommandLogs | merakictl get sm deviceCommandLogs {deviceId} {networkId} | --perPage --startingAfter --endingBefore | Return historical records of commands sent to Systems Manager devices. Note that this will include the name of the Dashboard user who initiated the command if it was generated by a Dashboard admin rather than the automatic behavior of the system; you may wish to filter this out of any reports.
  GET  | performanceHistory | merakictl get sm performanceHistory {deviceId} {networkId} | --perPage --startingAfter --endingBefore | Return historical records of various Systems Manager client metrics for desktop devices.
