package client

// EndpointApplicationModule is the endpoint to manage application modules records.
const EndpointApplicationModule = "sys_app_module.do"

// ApplicationModule is the json response for an application menu in ServiceNow.
type ApplicationModule struct {
	BaseResult
	Title             string `json:"title"`
	MenuID            string `json:"application"`
	Hint              string `json:"hint"`
	Order             int    `json:"order,string"`
	Roles             string `json:"roles"`
	Active            bool   `json:"active,string"`
	OverrideMenuRoles bool   `json:"override_menu_roles,string"`
	LinkType          string `json:"link_type"`
	Arguments         string `json:"query"`
	WindowName        string `json:"window_name"`
	TableName         string `json:"name"`
}
