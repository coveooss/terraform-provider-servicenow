package client

import (
	"fmt"
)

const endpointApplicationModule = "sys_app_module.do"

// ApplicationModule is the json response for an application menu in ServiceNow.
type ApplicationModule struct {
	BaseResult
	Title             string `json:"title"`
	MenuID            string `json:"application"`
	Hint              string `json:"hint"`
	Order             int    `json:"order,string"`
	Roles             string `json:"roles"`
	Active            bool   `json:"active,string"`
	OverrideMenuRoles bool   `json:"override_menu_roles,string"`
	LinkType          string `json:"link_type"`
	Arguments         string `json:"query"`
	WindowName        string `json:"window_name"`
	TableName         string `json:"name"`
}

// ApplicationModuleResults is the object returned by ServiceNow API when saving or retrieving records.
type ApplicationModuleResults struct {
	Records []ApplicationModule `json:"records"`
}

// GetApplicationModule retrieves a specific ApplicationModule in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetApplicationModule(id string) (*ApplicationModule, error) {
	applicationModulePageResults := ApplicationModuleResults{}
	if err := client.getObject(endpointApplicationModule, id, &applicationModulePageResults); err != nil {
		return nil, err
	}

	return &applicationModulePageResults.Records[0], nil
}

// CreateApplicationModule creates a new ApplicationModule in ServiceNow and returns the newly created applicationModule. The new applicationModule should
// include the GUID (sys_id) created in ServiceNow.
func (client *ServiceNowClient) CreateApplicationModule(applicationModule *ApplicationModule) (*ApplicationModule, error) {
	applicationModulePageResults := ApplicationModuleResults{}
	if err := client.createObject(endpointApplicationModule, applicationModule.Scope, applicationModule, &applicationModulePageResults); err != nil {
		return nil, err
	}

	return &applicationModulePageResults.Records[0], nil
}

// UpdateApplicationModule updates a ApplicationModule in ServiceNow.
func (client *ServiceNowClient) UpdateApplicationModule(applicationModule *ApplicationModule) error {
	return client.updateObject(endpointApplicationModule, applicationModule.Id, applicationModule)
}

// DeleteApplicationModule deletes a ApplicationModule in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteApplicationModule(id string) error {
	return client.deleteObject(endpointApplicationModule, id)
}

func (results ApplicationModuleResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
