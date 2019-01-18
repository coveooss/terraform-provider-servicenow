package client

// EndpointSystemProperty is the endpoint to manage system property records.
const EndpointSystemProperty = "sys_properties.do"

// SystemProperty is the json response for a system property in ServiceNow.
type SystemProperty struct {
	BaseResult
	Suffix      string `json:"suffix"`
	Type        string `json:"type"`
	Choices     string `json:"choices"`
	IsPrivate   bool   `json:"is_private,string"`
	IgnoreCache bool   `json:"ignore_cache,string"`
	Description string `json:"description"`
	WriteRoles  string `json:"write_roles"`
	ReadRoles   string `json:"read_roles"`
	Name        string `json:"name,omitempty"`
	Value       string `json:"value,omitempty"`
}
