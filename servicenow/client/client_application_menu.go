package client

import (
	"fmt"
)

const endpointApplicationMenu = "sys_app_application.do"

// ApplicationMenu is the json response for an application menu in ServiceNow.
type ApplicationMenu struct {
	BaseResult
	Title       string `json:"title"`
	Description string `json:"description"`
	Hint        string `json:"hint"`
	DeviceType  string `json:"device_type"`
	Order       int    `json:"order,string"`
	Roles       string `json:"roles"`
	CategoryID  string `json:"category"`
	Active      bool   `json:"active,string"`
}

// ApplicationMenuResults is the object returned by ServiceNow API when saving or retrieving records.
type ApplicationMenuResults struct {
	Records []ApplicationMenu `json:"records"`
}

// GetApplicationMenu retrieves a specific ApplicationMenu in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetApplicationMenu(id string) (*ApplicationMenu, error) {
	applicationMenuPageResults := ApplicationMenuResults{}
	if err := client.getObject(endpointApplicationMenu, id, &applicationMenuPageResults); err != nil {
		return nil, err
	}

	return &applicationMenuPageResults.Records[0], nil
}

// CreateApplicationMenu creates a new ApplicationMenu in ServiceNow and returns the newly created applicationMenu. The new applicationMenu should
// include the GUID (sys_id) created in ServiceNow.
func (client *ServiceNowClient) CreateApplicationMenu(applicationMenu *ApplicationMenu) (*ApplicationMenu, error) {
	applicationMenuPageResults := ApplicationMenuResults{}
	if err := client.createObject(endpointApplicationMenu, applicationMenu, &applicationMenuPageResults); err != nil {
		return nil, err
	}

	return &applicationMenuPageResults.Records[0], nil
}

// UpdateApplicationMenu updates a ApplicationMenu in ServiceNow.
func (client *ServiceNowClient) UpdateApplicationMenu(applicationMenu *ApplicationMenu) error {
	return client.updateObject(endpointApplicationMenu, applicationMenu.Id, applicationMenu)
}

// DeleteApplicationMenu deletes a ApplicationMenu in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteApplicationMenu(id string) error {
	return client.deleteObject(endpointApplicationMenu, id)
}

func (results ApplicationMenuResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
