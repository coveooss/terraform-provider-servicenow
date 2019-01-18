package client

// EndpointApplication is the endpoint to manage application records.
const EndpointApplication = "sys_app.do"

// Application is the json response for an application in ServiceNow.
type Application struct {
	BaseResult
	Name    string `json:"name"`
	Scope   string `json:"scope"`
	Version string `json:"version"`
}
