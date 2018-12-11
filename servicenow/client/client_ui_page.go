package client

import (
	"fmt"
)

const endpointUiPage = "sys_ui_page.do"

// UiPage is the json response for a UI page in ServiceNow.
type UiPage struct {
	BaseResult
	Name             string `json:"name"`
	Description      string `json:"description"`
	Direct           bool   `json:"direct,string"`
	Html             string `json:"html"`
	ProcessingScript string `json:"processing_script"`
	ClientScript     string `json:"client_script"`
	Category         string `json:"category"`
}

// UiPageResults is the object returned by ServiceNow API when saving or retrieving records.
type UiPageResults struct {
	Records []UiPage `json:"records"`
}

// GetUiPage retrieves a specific UI Page in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetUiPage(id string) (*UiPage, error) {
	uiPageResults := UiPageResults{}
	if err := client.getObject(endpointUiPage, id, &uiPageResults); err != nil {
		return nil, err
	}

	return &uiPageResults.Records[0], nil
}

// CreateUiPage creates a new UI Page in ServiceNow and returns the newly created page. The new page should
// include the GUID (sys_id) created in ServiceNow.
func (client *ServiceNowClient) CreateUiPage(uiPage *UiPage) (*UiPage, error) {
	uiPageResults := UiPageResults{}
	if err := client.createObject(endpointUiPage, uiPage, &uiPageResults); err != nil {
		return nil, err
	}

	return &uiPageResults.Records[0], nil
}

// UpdateUiPage updates a UI Page in ServiceNow.
func (client *ServiceNowClient) UpdateUiPage(uiPage *UiPage) error {
	return client.updateObject(endpointUiPage, uiPage.Id, uiPage)
}

// DeleteUiPage deletes a UI Page in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteUiPage(id string) error {
	return client.deleteObject(endpointUiPage, id)
}

func (results UiPageResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
