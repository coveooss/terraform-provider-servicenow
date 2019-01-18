package client

// EndpointDBTable is the endpoint to manage DB Table records.
const EndpointDBTable = "sys_db_object.do"

// DBTable is the json response for a Table in ServiceNow.
type DBTable struct {
	BaseResult
	Label                string `json:"label"`
	UserRole             string `json:"user_role"`
	Access               string `json:"access"`
	ReadAccess           bool   `json:"read_access,string"`
	CreateAccess         bool   `json:"create_access,string"`
	AlterAccess          bool   `json:"alter_access,string"`
	DeleteAccess         bool   `json:"delete_access,string"`
	WebServiceAccess     bool   `json:"ws_access,string"`
	ConfigurationAccess  bool   `json:"configuration_access,string"`
	Extendable           bool   `json:"is_extendable,string"`
	LiveFeed             bool   `json:"live_feed_enabled,string"`
	CreateAccessControls bool   `json:"create_access_controls,string"`
	CreateModule         bool   `json:"create_module,string"`
	CreateMobileModule   bool   `json:"create_mobile_module,string"`
	Name                 string `json:"name,omitempty"`
}
