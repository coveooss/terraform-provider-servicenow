package client

// EndpointWidgetDependency is the endpoint to manage widget dependency records.
const EndpointWidgetDependency = "sp_dependency.do"

// WidgetDependency represents the json response for a Widget Dependency in ServiceNow.
type WidgetDependency struct {
	BaseResult
	Name     string `json:"name"`
	Module   string `json:"module"`
	PageLoad bool   `json:"page_load,string"`
}
