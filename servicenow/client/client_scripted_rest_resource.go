package client

// EndpointScriptedRestResource is the endpoint to manage Scripted Rest Resources records.
const EndpointScriptedRestResource = "sys_ws_operation.do"

// EndpointScriptedRestResource is the json response for a Scripted Rest Resources in ServiceNow.
type ScriptedRestResource struct {
	BaseResult
	Name                     string `json:"name"`
	Active                   bool   `json:"active,string"`
	EnforceACL               string `json:"enforce_acl,omitempty"`
	RequiresACLAuthorization bool   `json:"requires_acl_authorization,string"`
	RequiresAuthentication   bool   `json:"requires_authentication,string"`
	RequiresSNCInternalRole  bool   `json:"requires_snc_internal_role,string"`
	Produces                 string `json:"produces,omitempty"`
	ProducesCustomized       bool   `json:"produces_customized,string"`
	ShortDescription         string `json:"short_description,omitempty"`
	OperationScript          string `json:"operation_script"`
	RelativePath             string `json:"relative_path,omitempty"`
	RequestExample           string `json:"request_example,omitempty"`
	HTTPMethod               string `json:"http_method"`
	Consumes                 string `json:"consumes,omitempty"`
	ConsumesCustomized       bool   `json:"consumes_customized,string"`
	OperationURI             string `json:"operation_uri,omitempty"`
	WebServiceDefinition     string `json:"web_service_definition,omitempty"`
	WebServiceVersion        string `json:"web_service_version,omitempty"`
}
