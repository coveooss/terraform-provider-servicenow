package client

// EndpointRestMethod is the endpoint to manage REST message method records.
const EndpointRestMethod = "sys_rest_message_fn.do"

// RestMethod represents the json response for a HTTP method in ServiceNow.
type RestMethod struct {
	BaseResult
	Name               string `json:"function_name"`
	MessageID          string `json:"rest_message"`
	HTTPMethod         string `json:"http_method"`
	RestEndpoint       string `json:"rest_endpoint"`
	AuthenticationType string `json:"authentication_type"`
	QualifiedName      string `json:"qualified_name,omitempty"`
}
