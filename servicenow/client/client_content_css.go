package client

// EndpointContentCSS is the endpoint to manage content css records.
const EndpointContentCSS = "content_css.do"

// ContentCSS represents the json response for a Content Management Style Sheet in ServiceNow.
type ContentCSS struct {
	BaseResult
	Name  string `json:"name"`
	Type  string `json:"type"`
	URL   string `json:"url"`
	Style string `json:"style"`
}
