package client

// EndpointScriptedRestApi is the endpoint to manage Scripted Rest Api records.
const EndpointScriptedRestApi = "sys_ws_definition.do"

// EndpointScriptedRestApi is the json response for a Scripted Rest Api in ServiceNow.
type ScriptedRestApi struct {
	BaseResult
	Active             bool   `json:"active,string"`
	Consumes           string `json:"consumes,omitempty"`
	ConsumesCustomized bool   `json:"consumes_customized,string"`
	EnforceACL         string `json:"enforce_acl,omitempty"`
	Name               string `json:"name"`
	Produces           string `json:"produces,omitempty"`
	ProducesCustomized bool   `json:"produces_customized,string"`
	ServiceId          string `json:"service_id,omitempty"`
	BaseURI            string `json:"base_uri,omitempty"`
	Namespace          string `json:"namespace,omitempty"`
	DocLink            string `json:"doc_link,omitempty"`
	ShortDescription   string `json:"short_description,omitempty"`
}
