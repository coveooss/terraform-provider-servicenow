package client

// EndpointScriptInclude is the endpoint to manage script include records.
const EndpointScriptInclude = "sys_script_include.do"

// ScriptInclude is the json response for a Script Include in ServiceNow.
type ScriptInclude struct {
	BaseResult
	Name           string `json:"name"`
	ClientCallable bool   `json:"client_callable,string"`
	Description    string `json:"description"`
	Script         string `json:"script"`
	Active         bool   `json:"active,string"`
	Access         string `json:"access"`
	APIName        string `json:"api_name,omitempty"`
}
