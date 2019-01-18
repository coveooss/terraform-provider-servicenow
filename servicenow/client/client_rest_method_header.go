package client

// EndpointRestMethodHeader is the endpoint to manage REST Message method header records.
const EndpointRestMethodHeader = "sys_rest_message_fn_headers.do"

// RestMethodHeader represents the json response for a HTTP header in ServiceNow.
type RestMethodHeader struct {
	BaseResult
	Name     string `json:"name"`
	Value    string `json:"value"`
	MethodID string `json:"rest_message_function"`
}
