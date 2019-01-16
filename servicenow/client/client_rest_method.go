package client

import (
	"fmt"
)

const endpointRestMethod = "sys_rest_message_fn.do"

// RestMethod represents the json response for a HTTP method in ServiceNow.
type RestMethod struct {
	BaseResult
	Name               string `json:"function_name"`
	MessageID          string `json:"rest_message"`
	HTTPMethod         string `json:"http_method"`
	RestEndpoint       string `json:"rest_endpoint"`
	AuthenticationType string `json:"authentication_type"`
	QualifiedName      string `json:"qualified_name,omitempty"`
}

// RestMethodResults is the object returned by ServiceNow API when saving or retrieving records.
type RestMethodResults struct {
	Records []RestMethod `json:"records"`
}

// GetRestMethod retrieves a specific HTTP method in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetRestMethod(id string) (*RestMethod, error) {
	restMethodResults := RestMethodResults{}
	if err := client.getObject(endpointRestMethod, id, &restMethodResults); err != nil {
		return nil, err
	}

	return &restMethodResults.Records[0], nil
}

// CreateRestMethod creates a HTTP method in ServiceNow and returns the newly created JS Include.
func (client *ServiceNowClient) CreateRestMethod(RestMethod *RestMethod) (*RestMethod, error) {
	restMethodResults := RestMethodResults{}
	if err := client.createObject(endpointRestMethod, RestMethod.Scope, RestMethod, &restMethodResults); err != nil {
		return nil, err
	}

	return &restMethodResults.Records[0], nil
}

// UpdateRestMethod updates a HTTP method in ServiceNow.
func (client *ServiceNowClient) UpdateRestMethod(restMethod *RestMethod) error {
	return client.updateObject(endpointRestMethod, restMethod.Id, restMethod)
}

// DeleteRestMethod deletes a HTTP method in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteRestMethod(id string) error {
	return client.deleteObject(endpointRestMethod, id)
}

func (results RestMethodResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
