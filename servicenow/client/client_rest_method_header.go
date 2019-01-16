package client

import (
	"fmt"
)

const endpointRestMethodHeader = "sys_rest_message_fn_headers.do"

// RestMethodHeader represents the json response for a HTTP header in ServiceNow.
type RestMethodHeader struct {
	BaseResult
	Name     string `json:"name"`
	Value    string `json:"value"`
	MethodID string `json:"rest_message_function"`
}

// RestMethodHeaderResults is the object returned by ServiceNow API when saving or retrieving records.
type RestMethodHeaderResults struct {
	Records []RestMethodHeader `json:"records"`
}

// GetRestMethodHeader retrieves a specific HTTP header in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetRestMethodHeader(id string) (*RestMethodHeader, error) {
	restMethodHeaderResults := RestMethodHeaderResults{}
	if err := client.getObject(endpointRestMethodHeader, id, &restMethodHeaderResults); err != nil {
		return nil, err
	}

	return &restMethodHeaderResults.Records[0], nil
}

// CreateRestMethodHeader creates a HTTP header in ServiceNow and returns the newly created JS Include.
func (client *ServiceNowClient) CreateRestMethodHeader(restMethodHeader *RestMethodHeader) (*RestMethodHeader, error) {
	restMethodHeaderResults := RestMethodHeaderResults{}
	if err := client.createObject(endpointRestMethodHeader, restMethodHeader.Scope, restMethodHeader, &restMethodHeaderResults); err != nil {
		return nil, err
	}

	return &restMethodHeaderResults.Records[0], nil
}

// UpdateRestMethodHeader updates a HTTP header in ServiceNow.
func (client *ServiceNowClient) UpdateRestMethodHeader(restMethodHeader *RestMethodHeader) error {
	return client.updateObject(endpointRestMethodHeader, restMethodHeader.Id, restMethodHeader)
}

// DeleteRestMethodHeader deletes a HTTP header in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteRestMethodHeader(id string) error {
	return client.deleteObject(endpointRestMethodHeader, id)
}

func (results RestMethodHeaderResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
