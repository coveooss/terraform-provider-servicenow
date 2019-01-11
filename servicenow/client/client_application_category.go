package client

import (
	"fmt"
)

const endpointApplicationCategory = "sys_app_category.do"

// ApplicationCategory represents the json response for a Application Category in ServiceNow.
type ApplicationCategory struct {
	BaseResult
	Name  string `json:"name"`
	Order int    `json:"default_order,string"`
	Style string `json:"style"`
}

// ApplicationCategoryResults is the object returned by ServiceNow API when saving or retrieving records.
type ApplicationCategoryResults struct {
	Records []ApplicationCategory `json:"records"`
}

// GetApplicationCategoryByName retrieves a specific Application Category in ServiceNow with it's name.
func (client *ServiceNowClient) GetApplicationCategoryByName(name string) (*ApplicationCategory, error) {
	applicationCategoryResults := ApplicationCategoryResults{}
	if err := client.getObjectByName(endpointApplicationCategory, name, &applicationCategoryResults); err != nil {
		return nil, err
	}

	return &applicationCategoryResults.Records[0], nil
}

func (results ApplicationCategoryResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
