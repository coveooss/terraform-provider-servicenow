package client

// EndpointApplicationMenu is the endpoint to manage application menu records.
const EndpointApplicationMenu = "sys_app_application.do"

// ApplicationMenu is the json response for an application menu in ServiceNow.
type ApplicationMenu struct {
	BaseResult
	Title       string `json:"title"`
	Description string `json:"description"`
	Hint        string `json:"hint"`
	DeviceType  string `json:"device_type"`
	Order       int    `json:"order,string"`
	Roles       string `json:"roles"`
	CategoryID  string `json:"category"`
	Active      bool   `json:"active,string"`
}
