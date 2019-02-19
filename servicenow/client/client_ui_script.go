package client

// EndpointUIScript is the endpoint to manage UI Macro records.
const EndpointUIScript = "sys_ui_script.do"

// UIScript is the json response for a UI Macro in ServiceNow.
type UIScript struct {
	BaseResult
	Name        string `json:"script_name"`
	APIName     string `json:"name,omitempty"`
	Description string `json:"description"`
	Script      string `json:"script"`
	Active      bool   `json:"active,string"`
	UIType      string `json:"ui_type"` // All: 10, Mobile: 1, Desktop 0
}
