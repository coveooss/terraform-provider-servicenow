package client

// EndpointJsIncludeRelation is the endpoint to manage JS Include Relation records.
const EndpointJsIncludeRelation = "m2m_sp_dependency_js_include.do"

// JsIncludeRelation represents the json response for a JsIncludeRelation in ServiceNow.
type JsIncludeRelation struct {
	BaseResult
	JsIncludeID  string `json:"sp_js_include"`
	DependencyID string `json:"sp_dependency"`
	Order        int    `json:"order,string"`
}
