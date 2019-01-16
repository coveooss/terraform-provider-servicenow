package client

import (
	"fmt"
)

const endpointRestMessageHeader = "sys_rest_message_headers.do"

// RestMessageHeader represents the json response for a HTTP header in ServiceNow.
type RestMessageHeader struct {
	BaseResult
	Name      string `json:"name"`
	Value     string `json:"value"`
	MessageID string `json:"rest_message"`
}

// RestMessageHeaderResults is the object returned by ServiceNow API when saving or retrieving records.
type RestMessageHeaderResults struct {
	Records []RestMessageHeader `json:"records"`
}

// GetRestMessageHeader retrieves a specific HTTP header in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetRestMessageHeader(id string) (*RestMessageHeader, error) {
	restMessageHeaderResults := RestMessageHeaderResults{}
	if err := client.getObject(endpointRestMessageHeader, id, &restMessageHeaderResults); err != nil {
		return nil, err
	}

	return &restMessageHeaderResults.Records[0], nil
}

// CreateRestMessageHeader creates a HTTP header in ServiceNow and returns the newly created JS Include.
func (client *ServiceNowClient) CreateRestMessageHeader(restMessageHeader *RestMessageHeader) (*RestMessageHeader, error) {
	restMessageHeaderResults := RestMessageHeaderResults{}
	if err := client.createObject(endpointRestMessageHeader, restMessageHeader.Scope, restMessageHeader, &restMessageHeaderResults); err != nil {
		return nil, err
	}

	return &restMessageHeaderResults.Records[0], nil
}

// UpdateRestMessageHeader updates a HTTP header in ServiceNow.
func (client *ServiceNowClient) UpdateRestMessageHeader(restMessageHeader *RestMessageHeader) error {
	return client.updateObject(endpointRestMessageHeader, restMessageHeader.Id, restMessageHeader)
}

// DeleteRestMessageHeader deletes a HTTP header in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteRestMessageHeader(id string) error {
	return client.deleteObject(endpointRestMessageHeader, id)
}

func (results RestMessageHeaderResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
