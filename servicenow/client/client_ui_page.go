package client

// EndpointUIPage is the endpoint to manage UI Pages records.
const EndpointUIPage = "sys_ui_page.do"

// UIPage is the json response for a UI page in ServiceNow.
type UIPage struct {
	BaseResult
	Name             string `json:"name"`
	Description      string `json:"description"`
	Direct           bool   `json:"direct,string"`
	HTML             string `json:"html"`
	ProcessingScript string `json:"processing_script"`
	ClientScript     string `json:"client_script"`
	Category         string `json:"category"`
	Endpoint         string `json:"endpoint,omitempty"`
}
