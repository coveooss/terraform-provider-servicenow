package client

import (
	"fmt"
)

const endpointContentCss = "content_css.do"

// ContentCss represents the json response for a Content Management Style Sheet in ServiceNow.
type ContentCss struct {
	BaseResult
	Name  string `json:"name"`
	Type  string `json:"type"`
	URL   string `json:"url"`
	Style string `json:"style"`
}

// ContentCssResults is the object returned by ServiceNow API when saving or retrieving records.
type ContentCssResults struct {
	Records []ContentCss `json:"records"`
}

// GetContentCss retrieves a specific Content Management Style Sheet in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetContentCss(id string) (*ContentCss, error) {
	contentCssResults := ContentCssResults{}
	if err := client.getObject(endpointContentCss, id, &contentCssResults); err != nil {
		return nil, err
	}

	return &contentCssResults.Records[0], nil
}

// CreateContentCss creates a Content Management Style Sheet in ServiceNow and returns the newly created JS Include.
func (client *ServiceNowClient) CreateContentCss(contentCss *ContentCss) (*ContentCss, error) {
	contentCssResults := ContentCssResults{}
	if err := client.createObject(endpointContentCss, contentCss.Scope, contentCss, &contentCssResults); err != nil {
		return nil, err
	}

	return &contentCssResults.Records[0], nil
}

// UpdateContentCss updates a Content Management Style Sheet in ServiceNow.
func (client *ServiceNowClient) UpdateContentCss(contentCss *ContentCss) error {
	return client.updateObject(endpointContentCss, contentCss.Id, contentCss)
}

// DeleteContentCss deletes a Content Management Style Sheet in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteContentCss(id string) error {
	return client.deleteObject(endpointContentCss, id)
}

func (results ContentCssResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
