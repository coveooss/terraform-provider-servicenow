package client

import (
	"encoding/json"
	"fmt"
)

// Represents the object returned by ServiceNow API when saving or retrieving records.
type UiPageResults struct {
	Results []UiPage `json:"records"`
}

// Represents the json response for a UI page in ServiceNow.
type UiPage struct {
	Id               string `json:"sys_id,omitempty"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Direct           bool   `json:"direct,string"`
	Html             string `json:"html"`
	ProcessingScript string `json:"processing_script"`
	ClientScript     string `json:"client_script"`
	Category         string `json:"category"`

	Status string       `json:"__status,omitempty"`
	Error  *ErrorDetail `json:"__error,omitempty"`
}

// Represents the details of an error. Should be included in the json if status is not success.
type ErrorDetail struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

// Retrieve a specific UI Page in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetUiPage(id string) (*UiPage, error) {
	jsonResponse, err := client.requestJSON("GET", "sys_ui_page.do?JSONv2&sysparm_query=sys_id="+id, nil)
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

// Creates a new UI Page in ServiceNow and returns the newly created page. The new page should
// include the GUID (sys_id) created in ServiceNow.
func (client *ServiceNowClient) CreateUiPage(uiPage *UiPage) (*UiPage, error) {
	jsonResponse, err := client.requestJSON("POST", "sys_ui_page.do?JSONv2&sysparm_action=insert", uiPage)
	if err != nil {
		return nil, err
	}

	uiPageResults := UiPageResults{}
	if err = json.Unmarshal(jsonResponse, &uiPageResults); err != nil {
		return nil, err
	}

	var page UiPage
	if len(uiPageResults.Results) <= 0 {
		err = fmt.Errorf("No UI Page were inserted.")
	} else {
		page := &uiPageResults.Results[0]
		if page.Status != "success" {
			err = fmt.Errorf("Error during insert. %s: %s.", page.Error.Message, page.Error.Reason)
		}
	}

	return &page, err
}

// Updates a UI Page in ServiceNow.
func (client *ServiceNowClient) UpdateUiPage(uiPage *UiPage) error {
	_, err := client.requestJSON("POST", "sys_ui_page.do?JSONv2&sysparm_action=update&sysparm_query=sys_id="+uiPage.Id, uiPage)
	return err
}

// Deletes a UI Page in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteUiPage(id string) error {
	_, err := client.requestJSON("POST", "sys_ui_page.do?JSONv2&sysparm_action=deleteRecord&sysparm_sys_id="+id, nil)
	return err
}
