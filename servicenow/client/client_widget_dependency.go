package client

import (
	"encoding/json"
	"fmt"
)

const endpointWidgetDependencyRelation = "m2m_sp_widget_dependency.do"

// WidgetDependencyRelation represents the json response for a UI page in ServiceNow.
type WidgetDependencyRelation struct {
	BaseResult
	DependencyId string `json:"sp_dependency"`
	WidgetId     string `json:"sp_widget"`
}

// WidgetDependencyRelationResults is the object returned by ServiceNow API when saving or retrieving records.
type WidgetDependencyRelationResults struct {
	Results []WidgetDependencyRelation `json:"records"`
}

// GetWidgetDependencyRelation retrieves a specific UI Page in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetWidgetDependencyRelation(id string) (*WidgetDependencyRelation, error) {
	jsonResponse, err := client.requestJSON("GET", endpointWidgetDependencyRelation+"?JSONv2&sysparm_query=sys_id="+id, nil)
	if err != nil {
		return nil, err
	}

	relationResults := WidgetDependencyRelationResults{}
	if err := json.Unmarshal(jsonResponse, &relationResults); err != nil {
		return nil, err
	}

	if len(relationResults.Results) <= 0 {
		return nil, fmt.Errorf("No UI Page found using sys_id %s", id)
	}

	return &relationResults.Results[0], nil
}

// CreateWidgetDependencyRelation creates a widget dependency relation in ServiceNow and returns the newly created relation.
func (client *ServiceNowClient) CreateWidgetDependencyRelation(relation *WidgetDependencyRelation) (*WidgetDependencyRelation, error) {
	jsonResponse, err := client.requestJSON("POST", endpointWidgetDependencyRelation+"?JSONv2&sysparm_action=insert", relation)
	if err != nil {
		return nil, err
	}

	relationResults := WidgetDependencyRelationResults{}
	if err = json.Unmarshal(jsonResponse, &relationResults); err != nil {
		return nil, err
	}

	var createdRelation WidgetDependencyRelation
	if len(relationResults.Results) <= 0 {
		err = fmt.Errorf("nothing was inserted")
	} else {
		createdRelation := &relationResults.Results[0]
		if createdRelation.Status != "success" {
			err = fmt.Errorf("error during insert -> %s: %s", createdRelation.Error.Message, createdRelation.Error.Reason)
		}
	}

	return &createdRelation, err
}

// UpdateWidgetDependencyRelation updates a widget dependency relation in ServiceNow.
func (client *ServiceNowClient) UpdateWidgetDependencyRelation(relation *WidgetDependencyRelation) error {
	return client.updateObject(endpointWidgetDependencyRelation, relation.Id, relation)
}

// DeleteWidgetDependencyRelation deletes a widget dependency relation in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteWidgetDependencyRelation(id string) error {
	return client.deleteObject(endpointWidgetDependencyRelation, id)
}