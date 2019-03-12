package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Client is the client used to interact with ServiceNow API.
type Client struct {
	BaseURL string
	Auth    string
}

// ServiceNowClient defines possible methods to call on the ServiceNowClient.
type ServiceNowClient interface {
	GetObject(string, string, Record) error
	GetObjectByName(string, string, Record) error
	CreateObject(string, Record) error
	UpdateObject(string, Record) error
	DeleteObject(string, string) error
}

// BaseResult is representing the default properties of all results.
type BaseResult struct {
	ID               string       `json:"sys_id,omitempty"`
	ProtectionPolicy string       `json:"sys_policy,omitempty"`
	Scope            string       `json:"sys_scope,omitempty"`
	Status           string       `json:"__status,omitempty"`
	Error            *ErrorDetail `json:"__error,omitempty"`
}

// Record is the interface for any BaseResult.
type Record interface {
	GetID() string
	GetScope() string
	GetStatus() string
	GetError() *ErrorDetail
}

// BaseResultList represents the response from the API. Records are always returned inside an array.
type BaseResultList struct {
	Records []json.RawMessage
}

// ErrorDetail is the details of an error. Should be included in the json if status is not success.
type ErrorDetail struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

// NewClient is a factory method used to return a new ServiceNowClient.
func NewClient(baseURL string, username string, password string) *Client {
	// Concatenate username + password to create a basic authorization header.
	credentials := username + ":" + password
	return &Client{
		BaseURL: baseURL,
		Auth:    "Basic " + base64.StdEncoding.EncodeToString([]byte(credentials)),
	}
}

// GetID returns the ID of a BaseRecord.
func (record BaseResult) GetID() string {
	return record.ID
}

// GetStatus returns the Status of a BaseRecord.
func (record BaseResult) GetStatus() string {
	return record.Status
}

// GetError returns the Error of a BaseRecord, if any.
func (record BaseResult) GetError() *ErrorDetail {
	return record.Error
}

// GetScope returns the Scope of a BaseRecord.
func (record BaseResult) GetScope() string {
	return record.Scope
}

// validateOnlyOneResultReceived checks if a BaseResultList has exactly one record.
func validateOnlyOneResultReceived(results BaseResultList) error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	}
	return nil
}

// GetObject retrieves an object via a specific endpoint with a GET method and a specified
// sys_id. The response is parsed and fills the object in parameters. responseObjectOut
// parameter must be a pointer.
func (client *Client) GetObject(endpoint string, id string, responseObjectOut Record) error {
	jsonResponse, err := client.requestJSON("GET", endpoint+"?JSONv2&sysparm_query=sys_id="+id, nil)
	if err != nil {
		return err
	}
	return parseResponseToRecord(jsonResponse, responseObjectOut)
}

// GetObjectByName retrieves an object via its name attribute.
func (client *Client) GetObjectByName(endpoint string, name string, responseObjectOut Record) error {
	jsonResponse, err := client.requestJSON("GET", endpoint+"?JSONv2&sysparm_query=name="+url.QueryEscape(name), nil)
	if err != nil {
		return err
	}
	return parseResponseToRecord(jsonResponse, responseObjectOut)
}

// CreateObject creates a new object in ServiceNow, validates the response and fills the object
// with properties received from the service.
func (client *Client) CreateObject(endpoint string, objectToCreate Record) error {
	url := endpoint + "?JSONv2&sysparm_action=insert"
	if objectToCreate.GetScope() != "" {
		url += "&sysparm_record_scope=" + objectToCreate.GetScope()
	}

	jsonResponse, err := client.requestJSON("POST", url, objectToCreate)
	if err != nil {
		return err
	}

	// Replace the object to create with the data from the object created.
	return parseResponseToRecord(jsonResponse, objectToCreate)
}

// UpdateObject updates an object using a specific endpoint, sys_id and object data.
func (client *Client) UpdateObject(endpoint string, object Record) error {
	_, err := client.requestJSON("POST", endpoint+"?JSONv2&sysparm_action=update&sysparm_query=sys_id="+object.GetID(), object)
	return err
}

// DeleteObject deletes an object using a specific endpoing and sys_id.
func (client *Client) DeleteObject(endpoint string, id string) error {
	_, err := client.requestJSON("POST", endpoint+"?JSONv2&sysparm_action=deleteRecord&sysparm_sys_id="+id, nil)
	return err
}

// requestJSON execute an HTTP request and returns the raw response data.
func (client *Client) requestJSON(method string, path string, jsonData interface{}) ([]byte, error) {
	var data *bytes.Buffer

	if jsonData != nil {
		jsonValue, _ := json.Marshal(jsonData)
		data = bytes.NewBuffer(jsonValue)
	} else {
		data = bytes.NewBuffer(nil)
	}

	request, _ := http.NewRequest(method, client.BaseURL+path, data)

	// Add the needed headers.
	request.Header.Set("Authorization", client.Auth)
	request.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	responseData, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode >= 300 || response.StatusCode < 200 {
		return nil, fmt.Errorf("HTTP response status %s, %s", response.Status, responseData)
	}

	return responseData, nil
}

func parseResponseToRecord(jsonResponse []byte, responseObjectOut Record) error {
	// Parse the response in the generic struct and validate it.
	baseResultsList := BaseResultList{}
	if err := json.Unmarshal(jsonResponse, &baseResultsList); err != nil {
		return err
	}

	if err := validateOnlyOneResultReceived(baseResultsList); err != nil {
		return err
	}

	// Parse the Record into its concrete type and validate it.
	if err := json.Unmarshal(baseResultsList.Records[0], responseObjectOut); err != nil {
		return err
	}

	return checkStatus(responseObjectOut)
}

// validate checks if the specified Record is in error or not.
func checkStatus(record Record) error {
	if record.GetStatus() != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", record.GetError().Message, record.GetError().Reason)
	}
	return nil
}
