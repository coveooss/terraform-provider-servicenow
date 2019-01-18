package client

// EndpointRestMessageHeader is the endpoint to manage REST message header records.
const EndpointRestMessageHeader = "sys_rest_message_headers.do"

// RestMessageHeader represents the json response for a HTTP header in ServiceNow.
type RestMessageHeader struct {
	BaseResult
	Name      string `json:"name"`
	Value     string `json:"value"`
	MessageID string `json:"rest_message"`
}
