package client

// EndointACL is the endpoint to manage ACL records.
const EndpointACL = "sys_security_acl.do"

// ACL is the json response for an ACL in ServiceNow.
type ACL struct {
	BaseResult
	Type           string `json:"type"`
	Operation      string `json:"operation"`
	AdminOverrides bool   `json:"admin_overrides,string"`
	Name           string `json:"name"`
	Description    string `json:"description,omitempty"`
	Active         bool   `json:"active,string"`
	Advanced       bool   `json:"advanced,string"`
	Condition      string `json:"condition,omitempty"`
	Script         string `json:"script,omitempty"`
}
