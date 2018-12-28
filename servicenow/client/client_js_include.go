package client

import (
	"fmt"
)

const endpointJsInclude = "sp_js_include.do"

// JsInclude represents the json response for a Js Include in ServiceNow.
type JsInclude struct {
	BaseResult
	Source      string `json:"source"`
	DisplayName string `json:"display_name"`
	Url         string `json:"url"`
	UiScriptId  string `json:"sys_ui_script"`
}

// JsIncludeResults is the object returned by ServiceNow API when saving or retrieving records.
type JsIncludeResults struct {
	Records []JsInclude `json:"records"`
}

// GetJsInclude retrieves a specific Js Include in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetJsInclude(id string) (*JsInclude, error) {
	jsIncludeResults := JsIncludeResults{}
	if err := client.getObject(endpointJsInclude, id, &jsIncludeResults); err != nil {
		return nil, err
	}

	return &jsIncludeResults.Records[0], nil
}

// CreateJsInclude creates a Js Include in ServiceNow and returns the newly created JS Include.
func (client *ServiceNowClient) CreateJsInclude(jsInclude *JsInclude) (*JsInclude, error) {
	jsIncludeResults := JsIncludeResults{}
	if err := client.createObject(endpointJsInclude, jsInclude.Scope, jsInclude, &jsIncludeResults); err != nil {
		return nil, err
	}

	return &jsIncludeResults.Records[0], nil
}

// UpdateJsInclude updates a Js Include in ServiceNow.
func (client *ServiceNowClient) UpdateJsInclude(jsInclude *JsInclude) error {
	return client.updateObject(endpointJsInclude, jsInclude.Id, jsInclude)
}

// DeleteJsInclude deletes a Js Include in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteJsInclude(id string) error {
	return client.deleteObject(endpointJsInclude, id)
}

func (results JsIncludeResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
