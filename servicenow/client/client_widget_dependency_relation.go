package client

import (
	"fmt"
)

const endpointWidgetDependencyRelation = "m2m_sp_widget_dependency.do"

// WidgetDependencyRelation represents the json response for a widget dependency relation in ServiceNow.
type WidgetDependencyRelation struct {
	BaseResult
	DependencyId string `json:"sp_dependency"`
	WidgetId     string `json:"sp_widget"`
}

// WidgetDependencyRelationResults is the object returned by ServiceNow API when saving or retrieving records.
type WidgetDependencyRelationResults struct {
	Records []WidgetDependencyRelation `json:"records"`
}

// GetWidgetDependencyRelation retrieves a specific widget dependency relation in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetWidgetDependencyRelation(id string) (*WidgetDependencyRelation, error) {
	relationResults := WidgetDependencyRelationResults{}
	if err := client.getObject(endpointWidgetDependencyRelation, id, &relationResults); err != nil {
		return nil, err
	}

	return &relationResults.Records[0], nil
}

// CreateWidgetDependencyRelation creates a widget dependency relation in ServiceNow and returns the newly created relation.
func (client *ServiceNowClient) CreateWidgetDependencyRelation(relation *WidgetDependencyRelation) (*WidgetDependencyRelation, error) {
	relationResults := WidgetDependencyRelationResults{}
	if err := client.createObject(endpointWidgetDependencyRelation, relation, &relationResults); err != nil {
		return nil, err
	}

	return &relationResults.Records[0], nil
}

// UpdateWidgetDependencyRelation updates a widget dependency relation in ServiceNow.
func (client *ServiceNowClient) UpdateWidgetDependencyRelation(relation *WidgetDependencyRelation) error {
	return client.updateObject(endpointWidgetDependencyRelation, relation.Id, relation)
}

// DeleteWidgetDependencyRelation deletes a widget dependency relation in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteWidgetDependencyRelation(id string) error {
	return client.deleteObject(endpointWidgetDependencyRelation, id)
}

func (results WidgetDependencyRelationResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
