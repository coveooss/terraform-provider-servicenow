package client

// EndpointSystemPropertyCategory is the endpoint to manage system property categories records.
const EndpointSystemPropertyCategory = "sys_properties_category.do"

// SystemPropertyCategory is the json response for a system property in ServiceNow.
type SystemPropertyCategory struct {
	BaseResult
	Name      string `json:"name"`
	TitleHTML string `json:"title"`
}
