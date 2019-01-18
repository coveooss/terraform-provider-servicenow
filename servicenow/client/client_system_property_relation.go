package client

// EndpointSystemPropertyRelation is the endpoint to manage system property category relation records.
const EndpointSystemPropertyRelation = "sys_properties.do"

// SystemPropertyRelation is the json response for a system property relation in ServiceNow.
type SystemPropertyRelation struct {
	BaseResult
	CategoryID string `json:"category"`
	PropertyID string `json:"property"`
	Order      string `json:"order,omitempty"`
}
