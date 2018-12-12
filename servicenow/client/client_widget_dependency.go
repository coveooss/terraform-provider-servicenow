package client

import (
	"fmt"
)

const endpointWidgetDependency = "sp_dependency.do"

// WidgetDependency represents the json response for a Widget Dependency in ServiceNow.
type WidgetDependency struct {
	BaseResult
	Name     string `json:"name"`
	Module   string `json:"module"`
	PageLoad bool   `json:"page_load,string"`
}

// WidgetDependencyResults is the object returned by ServiceNow API when saving or retrieving records.
type WidgetDependencyResults struct {
	Records []WidgetDependency `json:"records"`
}

// GetWidgetDependency retrieves a specific Widget Dependency in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetWidgetDependency(id string) (*WidgetDependency, error) {
	widgetDependencyResults := WidgetDependencyResults{}
	if err := client.getObject(endpointWidgetDependency, id, &widgetDependencyResults); err != nil {
		return nil, err
	}

	return &widgetDependencyResults.Records[0], nil
}

// CreateWidgetDependency creates a Widget Dependency in ServiceNow and returns the newly created widget dependency.
func (client *ServiceNowClient) CreateWidgetDependency(widgetDependency *WidgetDependency) (*WidgetDependency, error) {
	widgetDependencyResults := WidgetDependencyResults{}
	if err := client.createObject(endpointWidgetDependency, widgetDependency, &widgetDependencyResults); err != nil {
		return nil, err
	}

	return &widgetDependencyResults.Records[0], nil
}

// UpdateWidgetDependency updates a Widget Dependency in ServiceNow.
func (client *ServiceNowClient) UpdateWidgetDependency(widgetDependency *WidgetDependency) error {
	return client.updateObject(endpointWidgetDependency, widgetDependency.Id, widgetDependency)
}

// DeleteWidgetDependency deletes a Widget Dependency in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteWidgetDependency(id string) error {
	return client.deleteObject(endpointWidgetDependency, id)
}

func (results WidgetDependencyResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
