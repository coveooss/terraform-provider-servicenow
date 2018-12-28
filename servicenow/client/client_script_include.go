package client

import (
	"fmt"
)

const endpointScriptInclude = "sys_script_include.do"

// ScriptInclude is the json response for a Script Include in ServiceNow.
type ScriptInclude struct {
	BaseResult
	Name           string `json:"name"`
	ClientCallable bool   `json:"client_callable,string"`
	Description    string `json:"description"`
	Script         string `json:"script"`
	Active         bool   `json:"active,string"`
	Access         string `json:"access"`
	APIName        string `json:"api_name,omitempty"`
}

// ScriptIncludeResults is the object returned by ServiceNow API when saving or retrieving records.
type ScriptIncludeResults struct {
	Records []ScriptInclude `json:"records"`
}

// GetScriptInclude retrieves a specific ScriptInclude in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetScriptInclude(id string) (*ScriptInclude, error) {
	scriptIncludePageResults := ScriptIncludeResults{}
	if err := client.getObject(endpointScriptInclude, id, &scriptIncludePageResults); err != nil {
		return nil, err
	}

	return &scriptIncludePageResults.Records[0], nil
}

// CreateScriptInclude creates a new ScriptInclude in ServiceNow and returns the newly created scriptInclude. The new scriptInclude should
// include the GUID (sys_id) created in ServiceNow.
func (client *ServiceNowClient) CreateScriptInclude(scriptInclude *ScriptInclude) (*ScriptInclude, error) {
	scriptIncludePageResults := ScriptIncludeResults{}
	if err := client.createObject(endpointScriptInclude, scriptInclude, &scriptIncludePageResults); err != nil {
		return nil, err
	}

	return &scriptIncludePageResults.Records[0], nil
}

// UpdateScriptInclude updates a ScriptInclude in ServiceNow.
func (client *ServiceNowClient) UpdateScriptInclude(scriptInclude *ScriptInclude) error {
	return client.updateObject(endpointScriptInclude, scriptInclude.Id, scriptInclude)
}

// DeleteScriptInclude deletes a ScriptInclude in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteScriptInclude(id string) error {
	return client.deleteObject(endpointScriptInclude, id)
}

func (results ScriptIncludeResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
