package client

import (
	"encoding/json"
	"fmt"
)

const endpointApplication = "sys_app.do"

// Application is the json response for an application in ServiceNow.
type Application struct {
	BaseResult
	Name    string `json:"name"`
	Scope   string `json:"scope"`
	Version string `json:"version"`
}

// ApplicationResults is the object returned by ServiceNow API when saving or retrieving records.
type ApplicationResults struct {
	Records []Application `json:"records"`
}

// GetApplication retrieves a specific Application in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetApplication(id string) (*Application, error) {
	applicationPageResults := ApplicationResults{}
	if err := client.getObject(endpointApplication, id, &applicationPageResults); err != nil {
		return nil, err
	}

	return &applicationPageResults.Records[0], nil
}

// GetApplicationByName retrieves a specific Application in ServiceNow with it's name attribute.
func (client *ServiceNowClient) GetApplicationByName(name string) (*Application, error) {
	jsonResponse, err := client.requestJSON("GET", endpointApplication+"?JSONv2&sysparm_query=name="+name, nil)
	if err != nil {
		return nil, err
	}

	applicationPageResults := ApplicationResults{}
	if err := json.Unmarshal(jsonResponse, &applicationPageResults); err != nil {
		return nil, err
	}

	if err := applicationPageResults.validate(); err != nil {
		return nil, err
	}

	return &applicationPageResults.Records[0], nil
}

// CreateApplication creates a new Application in ServiceNow and returns the newly created application. The new application should
// include the GUID (sys_id) created in ServiceNow.
func (client *ServiceNowClient) CreateApplication(application *Application) (*Application, error) {
	applicationPageResults := ApplicationResults{}
	if err := client.createObject(endpointApplication, application, &applicationPageResults); err != nil {
		return nil, err
	}

	return &applicationPageResults.Records[0], nil
}

// UpdateApplication updates a Application in ServiceNow.
func (client *ServiceNowClient) UpdateApplication(application *Application) error {
	return client.updateObject(endpointApplication, application.Id, application)
}

// DeleteApplication deletes a Application in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteApplication(id string) error {
	return client.deleteObject(endpointApplication, id)
}

func (results ApplicationResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
