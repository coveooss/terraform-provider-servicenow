package client

import (
	"fmt"
)

const endpointUiMacro = "sys_ui_macro.do"

// UiMacro is the json response for a UI Macro in ServiceNow.
type UiMacro struct {
	BaseResult
	Name        string `json:"name"`
	Description string `json:"description"`
	APIName     string `json:"scoped_name"`
	Xml         string `json:"xml"`
	Active      bool   `json:"active,string"`
}

// UiMacroResults is the object returned by ServiceNow API when saving or retrieving records.
type UiMacroResults struct {
	Records []UiMacro `json:"records"`
}

// GetUiMacro retrieves a specific UI Macro in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetUiMacro(id string) (*UiMacro, error) {
	uiMacroPageResults := UiMacroResults{}
	if err := client.getObject(endpointUiMacro, id, &uiMacroPageResults); err != nil {
		return nil, err
	}

	return &uiMacroPageResults.Records[0], nil
}

// CreateUiMacro creates a new UiMacro in ServiceNow and returns the newly created UI Macro. The new UI Macro should
// include the GUID (sys_id) created in ServiceNow.
func (client *ServiceNowClient) CreateUiMacro(uiMacro *UiMacro) (*UiMacro, error) {
	uiMacroPageResults := UiMacroResults{}
	if err := client.createObject(endpointUiMacro, uiMacro, &uiMacroPageResults); err != nil {
		return nil, err
	}

	return &uiMacroPageResults.Records[0], nil
}

// UpdateUiMacro updates a UI Macro in ServiceNow.
func (client *ServiceNowClient) UpdateUiMacro(uiMacro *UiMacro) error {
	return client.updateObject(endpointUiMacro, uiMacro.Id, uiMacro)
}

// DeleteUiMacro deletes a UI Macro in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteUiMacro(id string) error {
	return client.deleteObject(endpointUiMacro, id)
}

func (results UiMacroResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
