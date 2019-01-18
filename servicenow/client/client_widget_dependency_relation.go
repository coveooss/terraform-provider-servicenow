package client

// EndpointWidgetDependencyRelation is the endpoint to manage widget dependency relation records.
const EndpointWidgetDependencyRelation = "m2m_sp_widget_dependency.do"

// WidgetDependencyRelation represents the json response for a widget dependency relation in ServiceNow.
type WidgetDependencyRelation struct {
	BaseResult
	DependencyID string `json:"sp_dependency"`
	WidgetID     string `json:"sp_widget"`
}
