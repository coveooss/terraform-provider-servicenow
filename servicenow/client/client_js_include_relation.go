package client

import (
	"fmt"
)

const endpointJsIncludeRelation = "m2m_sp_dependency_js_include.do"

// JsIncludeRelation represents the json response for a JsIncludeRelation in ServiceNow.
type JsIncludeRelation struct {
	BaseResult
	JsIncludeId  string `json:"sp_js_include"`
	DependencyId string `json:"sp_dependency"`
	Order        int    `json:"order,string"`
}

// JsIncludeRelationResults is the object returned by ServiceNow API when saving or retrieving records.
type JsIncludeRelationResults struct {
	Records []JsIncludeRelation `json:"records"`
}

// GetJsIncludeRelation retrieves a specific JsIncludeRelation in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetJsIncludeRelation(id string) (*JsIncludeRelation, error) {
	relationResults := JsIncludeRelationResults{}
	if err := client.getObject(endpointJsIncludeRelation, id, &relationResults); err != nil {
		return nil, err
	}

	return &relationResults.Records[0], nil
}

// CreateJsIncludeRelation creates a widget dependency relation in ServiceNow and returns the newly created relation.
func (client *ServiceNowClient) CreateJsIncludeRelation(relation *JsIncludeRelation) (*JsIncludeRelation, error) {
	relationResults := JsIncludeRelationResults{}
	if err := client.createObject(endpointJsIncludeRelation, relation, &relationResults); err != nil {
		return nil, err
	}

	return &relationResults.Records[0], nil
}

// UpdateJsIncludeRelation updates a widget dependency relation in ServiceNow.
func (client *ServiceNowClient) UpdateJsIncludeRelation(relation *JsIncludeRelation) error {
	return client.updateObject(endpointJsIncludeRelation, relation.Id, relation)
}

// DeleteJsIncludeRelation deletes a widget dependency relation in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteJsIncludeRelation(id string) error {
	return client.deleteObject(endpointJsIncludeRelation, id)
}

func (results JsIncludeRelationResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
