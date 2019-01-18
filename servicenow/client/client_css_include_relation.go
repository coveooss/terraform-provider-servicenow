package client

// EndpointCSSIncludeRelation is the endpoint to manage CSS include relation records.
const EndpointCSSIncludeRelation = "m2m_sp_dependency_css_include.do"

// CSSIncludeRelation represents the json response for a CssIncludeRelation in ServiceNow.
type CSSIncludeRelation struct {
	BaseResult
	CSSIncludeID string `json:"sp_css_include"`
	DependencyID string `json:"sp_dependency"`
	Order        int    `json:"order,string"`
}
