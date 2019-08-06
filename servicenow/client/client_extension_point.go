package client

// EndpointExtensionPoint is the endpoint to manage Extension Point records.
const EndpointExtensionPoint = "sys_extension_point.do"

// ExtensionPoint represents the json response for a Extension Point in ServiceNow.
type ExtensionPoint struct {
	BaseResult
	Name          string `json:"name"`
	RestrictScope bool   `json:"restrict_scope,string"`
	Description   string `json:"description"`
	Example       string `json:"example"`
	APIName       string `json:"api_name,omitempty"`
}
