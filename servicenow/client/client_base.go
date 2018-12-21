package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ServiceNowClient is the client used to interact with ServiceNow API.
type ServiceNowClient struct {
	BaseUrl string
	Auth    string
}

// RequestResults is the interface for request responses. Each resource should implement it's own
// validate method that will be called by the base client.
type RequestResults interface {
	validate() error
}

// BaseResult is representing the default properties of all results.
type BaseResult struct {
	Id               string       `json:"sys_id,omitempty"`
	ProtectionPolicy string       `json:"sys_policy,omitempty"`
	Status           string       `json:"__status,omitempty"`
	Error            *ErrorDetail `json:"__error,omitempty"`
}

// ErrorDetail is the details of an error. Should be included in the json if status is not success.
type ErrorDetail struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

// NewClient is a factory method used to return a new ServiceNowClient.
func NewClient(baseUrl string, username string, password string) *ServiceNowClient {
	// Concatenate username + password to create a basic authorization header.
	credentials := username + ":" + password
	return &ServiceNowClient{
		BaseUrl: baseUrl,
		Auth:    "Basic " + base64.StdEncoding.EncodeToString([]byte(credentials)),
	}
}

func (client *ServiceNowClient) requestJSON(method string, path string, jsonData interface{}) ([]byte, error) {
	var data *bytes.Buffer

	if jsonData != nil {
		jsonValue, _ := json.Marshal(jsonData)
		data = bytes.NewBuffer(jsonValue)
	} else {
		data = bytes.NewBuffer(nil)
	}

	request, _ := http.NewRequest(method, client.BaseUrl+path, data)

	// Add the needed headers.
	request.Header.Set("Authorization", client.Auth)
	request.Header.Set("Content-Type", "application/json")

	return client.getResponse(request)
}

func (client *ServiceNowClient) getResponse(request *http.Request) ([]byte, error) {
	httpClient := &http.Client{}
	response, err := httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	responseData, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode >= 300 || response.StatusCode < 200 {
		return nil, fmt.Errorf("Response status %s, %s", response.Status, responseData)
	}

	return responseData, nil
}

// getObject retrieves an object via a specific endpoint with a GET method and a specified
// sys_id. The response is parsed and fills the object in parameters. responseObjectOut
// parameter must be a pointer.
func (client *ServiceNowClient) getObject(endpoint string, id string, responseObjectOut RequestResults) error {
	jsonResponse, err := client.requestJSON("GET", endpoint+"?JSONv2&sysparm_query=sys_id="+id, nil)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonResponse, responseObjectOut); err != nil {
		return err
	}

	return responseObjectOut.validate()
}

// createObject creates a new object in ServiceNow, validates the response and returns it.
// responseObjectOut parameter must be a pointer.
func (client *ServiceNowClient) createObject(endpoint string, objectToCreate interface{}, responseObjectOut RequestResults) error {
	jsonResponse, err := client.requestJSON("POST", endpoint+"?JSONv2&sysparm_action=insert", objectToCreate)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonResponse, responseObjectOut); err != nil {
		return err
	}

	return responseObjectOut.validate()
}

// updateObject updates an object using a specific endpoint, sys_id and object data.
func (client *ServiceNowClient) updateObject(endpoint string, id string, object interface{}) error {
	_, err := client.requestJSON("POST", endpoint+"?JSONv2&sysparm_action=update&sysparm_query=sys_id="+id, object)
	return err
}

// deleteObject deletes an object using a specific endpoing and sys_id.
func (client *ServiceNowClient) deleteObject(endpoint string, id string) error {
	_, err := client.requestJSON("POST", endpoint+"?JSONv2&sysparm_action=deleteRecord&sysparm_sys_id="+id, nil)
	return err
}
