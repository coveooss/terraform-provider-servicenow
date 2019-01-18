package client

// EndpointUIMacro is the endpoint to manage UI Macro records.
const EndpointUIMacro = "sys_ui_macro.do"

// UIMacro is the json response for a UI Macro in ServiceNow.
type UIMacro struct {
	BaseResult
	Name        string `json:"name"`
	Description string `json:"description"`
	APIName     string `json:"scoped_name"`
	XML         string `json:"xml"`
	Active      bool   `json:"active,string"`
}
