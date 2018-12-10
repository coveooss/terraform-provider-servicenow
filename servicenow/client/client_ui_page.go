package client

import (
	"encoding/json"
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
	Results []UiPage `json:"records"`
}

// GetUiPage retrieves a specific UI Page in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetUiPage(id string) (*UiPage, error) {
	jsonResponse, err := client.requestJSON("GET", endpointUiPage + "?JSONv2&sysparm_query=sys_id="+id, nil)
	if err != nil {
		return nil, err
	}

	uiPageResults := UiPageResults{}
	if err := json.Unmarshal(jsonResponse, &uiPageResults); err != nil {
		return nil, err
	}

	if len(uiPageResults.Results) <= 0 {
		return nil, fmt.Errorf("No UI Page found using sys_id %s", id)
	}

	return &uiPageResults.Results[0], nil
}

// CreateUiPage creates a new UI Page in ServiceNow and returns the newly created page. The new page should
// include the GUID (sys_id) created in ServiceNow.
func (client *ServiceNowClient) CreateUiPage(uiPage *UiPage) (*UiPage, error) {
	jsonResponse, err := client.requestJSON("POST", endpointUiPage + "?JSONv2&sysparm_action=insert", uiPage)
	if err != nil {
		return nil, err
	}

	uiPageResults := UiPageResults{}
	if err = json.Unmarshal(jsonResponse, &uiPageResults); err != nil {
		return nil, err
	}

	var page UiPage
	if len(uiPageResults.Results) <= 0 {
		err = fmt.Errorf("no UI Page were inserted")
	} else {
		page := &uiPageResults.Results[0]
		if page.Status != "success" {
			err = fmt.Errorf("error during insert -> %s: %s", page.Error.Message, page.Error.Reason)
		}
	}

	return &page, err
}

// UpdateUiPage updates a UI Page in ServiceNow.
func (client *ServiceNowClient) UpdateUiPage(uiPage *UiPage) error {
	return client.updateObject(endpointUiPage, uiPage.Id, uiPage)
}

// DeleteUiPage deletes a UI Page in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteUiPage(id string) error {
	return client.deleteObject(endpointUiPage, id)
}