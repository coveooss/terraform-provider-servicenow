package client

import (
	"fmt"
)

const endpointRestMessage = "sys_rest_message.do"

// RestMessage represents the json response for a REST Message in ServiceNow.
type RestMessage struct {
	BaseResult
	Name               string `json:"name"`
	Description        string `json:"description"`
	RestEndpoint       string `json:"rest_endpoint"`
	Access             string `json:"access"`
	AuthenticationType string `json:"authentication_type"`
}

// RestMessageResults is the object returned by ServiceNow API when saving or retrieving records.
type RestMessageResults struct {
	Records []RestMessage `json:"records"`
}

// GetRestMessage retrieves a specific REST Message in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetRestMessage(id string) (*RestMessage, error) {
	restMessageResults := RestMessageResults{}
	if err := client.getObject(endpointRestMessage, id, &restMessageResults); err != nil {
		return nil, err
	}

	return &restMessageResults.Records[0], nil
}

// CreateRestMessage creates a REST Message in ServiceNow and returns the newly created endpoint.
func (client *ServiceNowClient) CreateRestMessage(restMessage *RestMessage) (*RestMessage, error) {
	restMessageResults := RestMessageResults{}
	if err := client.createObject(endpointRestMessage, restMessage.Scope, restMessage, &restMessageResults); err != nil {
		return nil, err
	}

	return &restMessageResults.Records[0], nil
}

// UpdateRestMessage updates a REST Message in ServiceNow.
func (client *ServiceNowClient) UpdateRestMessage(restMessage *RestMessage) error {
	return client.updateObject(endpointRestMessage, restMessage.Id, restMessage)
}

// DeleteRestMessage deletes a REST Message in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteRestMessage(id string) error {
	return client.deleteObject(endpointRestMessage, id)
}

func (results RestMessageResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
