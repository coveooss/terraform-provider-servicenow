package client

// EndpointRole is the endpoint to manage role records.
const EndpointRole = "sys_user_role.do"

// Role is the json response for a role in ServiceNow.
type Role struct {
	BaseResult
	Name              string `json:"name,omitempty"`
	Description       string `json:"description"`
	ElevatedPrivilege bool   `json:"elevated_privilege,string"`
	Suffix            string `json:"suffix"`
	AssignableBy      string `json:"assignable_by"`
}
