package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Attributes of the client used to interact with ServiceNow API.
type ServiceNowClient struct {
	BaseUrl string
	Auth    string
}

// Factory method to return a new ServiceNowClient.
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
