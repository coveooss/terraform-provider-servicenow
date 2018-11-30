package client

import (
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
	"encoding/base64"
	"io/ioutil"
)

type ServiceNowClient struct {
	BaseUrl  string
	Username string
	Password string
}

func (client *ServiceNowClient) requestJSON(method string, path string, jsonData interface{}) ([]byte, error) {
	var data *bytes.Buffer

	if jsonData != nil {
		jsonValue, _ := json.Marshal(jsonData)
		data = bytes.NewBuffer(jsonValue)
	} else {
		data = bytes.NewBuffer(nil)
	}

	request, _ := http.NewRequest(method, client.BaseUrl + path, data)
	client.formRequest(request, "application/json")

	return client.getResponse(request)
}

func (client *ServiceNowClient) formRequest(request *http.Request, contentType string) *http.Request {
	credentials := client.Username + ":" + client.Password
	request.Header.Set("Authorization", "Basic " + base64.StdEncoding.EncodeToString([]byte(credentials)))
	request.Header.Set("Content-Type", "application/json")

	return request
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