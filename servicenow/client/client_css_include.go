package client

// EndpointCSSInclude is the endpoint to manage CSS includes records.
const EndpointCSSInclude = "sp_css_include.do"

// CSSInclude represents the json response for a CSS Include in ServiceNow.
type CSSInclude struct {
	BaseResult
	Source       string `json:"source"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	StyleSheetID string `json:"sp_css"`
}
