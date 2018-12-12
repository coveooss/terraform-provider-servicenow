package client

import (
	"fmt"
)

const endpointCssIncludeRelation = "m2m_sp_dependency_css_include.do"

// CssIncludeRelation represents the json response for a CssIncludeRelation in ServiceNow.
type CssIncludeRelation struct {
	BaseResult
	CssIncludeId string `json:"sp_css_include"`
	DependencyId string `json:"sp_dependency"`
	Order        int    `json:"order,string"`
}

// CssIncludeRelationResults is the object returned by ServiceNow API when saving or retrieving records.
type CssIncludeRelationResults struct {
	Records []CssIncludeRelation `json:"records"`
}

// GetCssIncludeRelation retrieves a specific CssIncludeRelation in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetCssIncludeRelation(id string) (*CssIncludeRelation, error) {
	relationResults := CssIncludeRelationResults{}
	if err := client.getObject(endpointCssIncludeRelation, id, &relationResults); err != nil {
		return nil, err
	}

	return &relationResults.Records[0], nil
}

// CreateCssIncludeRelation creates a widget dependency relation in ServiceNow and returns the newly created relation.
func (client *ServiceNowClient) CreateCssIncludeRelation(relation *CssIncludeRelation) (*CssIncludeRelation, error) {
	relationResults := CssIncludeRelationResults{}
	if err := client.createObject(endpointCssIncludeRelation, relation, &relationResults); err != nil {
		return nil, err
	}

	return &relationResults.Records[0], nil
}

// UpdateCssIncludeRelation updates a widget dependency relation in ServiceNow.
func (client *ServiceNowClient) UpdateCssIncludeRelation(relation *CssIncludeRelation) error {
	return client.updateObject(endpointCssIncludeRelation, relation.Id, relation)
}

// DeleteCssIncludeRelation deletes a widget dependency relation in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteCssIncludeRelation(id string) error {
	return client.deleteObject(endpointCssIncludeRelation, id)
}

func (results CssIncludeRelationResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
