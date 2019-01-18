package client

// EndpointJsInclude is the endpoint to manage JS Include records.
const EndpointJsInclude = "sp_js_include.do"

// JsInclude represents the json response for a Js Include in ServiceNow.
type JsInclude struct {
	BaseResult
	Source      string `json:"source"`
	DisplayName string `json:"display_name"`
	URL         string `json:"url"`
	UIScriptID  string `json:"sys_ui_script"`
}
