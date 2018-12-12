package client

import (
	"fmt"
)

const endpointWidget = "sp_widget.do"

// Widget is the json response for a Widget in ServiceNow.
type Widget struct {
	BaseResult
	CustomId     string `json:"id"`
	Name         string `json:"name"`
	Template     string `json:"template"`
	Css          string `json:"css"`
	Public       bool   `json:"public,string"`
	Roles        string `json:"roles"`
	Link         string `json:"link"`
	Description  string `json:"description"`
	ClientScript string `json:"client_script"`
	ServerScript string `json:"script"`
	DemoData     string `json:"demo_data"`
	OptionSchema string `json:"option_schema"`
	HasPreview   bool   `json:"has_preview,string"`
	DataTable    string `json:"data_table"`
	ControllerAs string `json:"controller_as"`
}

// WidgetResults is the object returned by ServiceNow API when saving or retrieving records.
type WidgetResults struct {
	Records []Widget `json:"records"`
}

// GetWidget retrieves a specific Widget in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetWidget(id string) (*Widget, error) {
	widgetResults := WidgetResults{}
	if err := client.getObject(endpointWidget, id, &widgetResults); err != nil {
		return nil, err
	}

	return &widgetResults.Records[0], nil
}

// CreateWidget creates a new Widget in ServiceNow and returns the newly created page. The new page should
// include the GUID (sys_id) created in ServiceNow.
func (client *ServiceNowClient) CreateWidget(widget *Widget) (*Widget, error) {
	widgetResults := WidgetResults{}
	if err := client.createObject(endpointWidget, widget, &widgetResults); err != nil {
		return nil, err
	}

	return &widgetResults.Records[0], nil
}

// UpdateWidget updates a Widget in ServiceNow.
func (client *ServiceNowClient) UpdateWidget(widget *Widget) error {
	return client.updateObject(endpointWidget, widget.Id, widget)
}

// DeleteWidget deletes a Widget in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteWidget(id string) error {
	return client.deleteObject(endpointWidget, id)
}

func (results WidgetResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
