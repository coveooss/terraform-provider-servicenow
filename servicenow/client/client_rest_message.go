package client

// EndpointRestMessage is the endpoint to manage REST message records.
const EndpointRestMessage = "sys_rest_message.do"

// RestMessage represents the json response for a REST Message in ServiceNow.
type RestMessage struct {
	BaseResult
	Name               string `json:"name"`
	Description        string `json:"description"`
	RestEndpoint       string `json:"rest_endpoint"`
	Access             string `json:"access"`
	AuthenticationType string `json:"authentication_type"`
}
