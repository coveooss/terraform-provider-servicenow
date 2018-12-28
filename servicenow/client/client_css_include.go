package client

import (
	"fmt"
)

const endpointCssInclude = "sp_css_include.do"

// CssInclude represents the json response for a CSS Include in ServiceNow.
type CssInclude struct {
	BaseResult
	Source string `json:"source"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	CssId  string `json:"sp_css"`
}

// CssIncludeResults is the object returned by ServiceNow API when saving or retrieving records.
type CssIncludeResults struct {
	Records []CssInclude `json:"records"`
}

// GetCssInclude retrieves a specific CSS Include in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetCssInclude(id string) (*CssInclude, error) {
	cssIncludeResults := CssIncludeResults{}
	if err := client.getObject(endpointCssInclude, id, &cssIncludeResults); err != nil {
		return nil, err
	}

	return &cssIncludeResults.Records[0], nil
}

// CreateCssInclude creates a CSS Include in ServiceNow and returns the newly created CSS Include.
func (client *ServiceNowClient) CreateCssInclude(cssInclude *CssInclude) (*CssInclude, error) {
	cssIncludeResults := CssIncludeResults{}
	if err := client.createObject(endpointCssInclude, cssInclude.Scope, cssInclude, &cssIncludeResults); err != nil {
		return nil, err
	}

	return &cssIncludeResults.Records[0], nil
}

// UpdateCssInclude updates a CSS Include in ServiceNow.
func (client *ServiceNowClient) UpdateCssInclude(cssInclude *CssInclude) error {
	return client.updateObject(endpointCssInclude, cssInclude.Id, cssInclude)
}

// DeleteCssInclude deletes a CSS Include in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteCssInclude(id string) error {
	return client.deleteObject(endpointCssInclude, id)
}

func (results CssIncludeResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
