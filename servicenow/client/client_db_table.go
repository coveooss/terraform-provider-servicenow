package client

import (
	"fmt"
)

const endpointDBTable = "sys_db_object.do"

// DBTable is the json response for a Table in ServiceNow.
type DBTable struct {
	BaseResult
	Label                string `json:"label"`
	UserRole             string `json:"user_role"`
	Access               string `json:"access"`
	ReadAccess           bool   `json:"read_access,string"`
	CreateAccess         bool   `json:"create_access,string"`
	AlterAccess          bool   `json:"alter_access,string"`
	DeleteAccess         bool   `json:"delete_access,string"`
	WebServiceAccess     bool   `json:"ws_access,string"`
	ConfigurationAccess  bool   `json:"configuration_access,string"`
	Extendable           bool   `json:"is_extendable,string"`
	LiveFeed             bool   `json:"live_feed_enabled,string"`
	CreateAccessControls bool   `json:"create_access_controls,string"`
	CreateModule         bool   `json:"create_module,string"`
	CreateMobileModule   bool   `json:"create_mobile_module,string"`
	Name                 string `json:"name,omitempty"`
}

// DBTableResults is the object returned by ServiceNow API when saving or retrieving records.
type DBTableResults struct {
	Records []DBTable `json:"records"`
}

// GetDBTable retrieves a specific DBTable in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetDBTable(id string) (*DBTable, error) {
	dbTablePageResults := DBTableResults{}
	if err := client.getObject(endpointDBTable, id, &dbTablePageResults); err != nil {
		return nil, err
	}

	return &dbTablePageResults.Records[0], nil
}

// GetDBTableByName retrieves a specific DB Table in ServiceNow with it's name attribute.
func (client *ServiceNowClient) GetDBTableByName(name string) (*DBTable, error) {
	dbTablePageResults := DBTableResults{}
	if err := client.getObjectByName(endpointDBTable, name, &dbTablePageResults); err != nil {
		return nil, err
	}

	return &dbTablePageResults.Records[0], nil
}

// CreateDBTable creates a new DB Table in ServiceNow and returns the newly created Table. The new Table should
// include the GUID (sys_id) created in ServiceNow.
func (client *ServiceNowClient) CreateDBTable(dbTable *DBTable) (*DBTable, error) {
	dbTablePageResults := DBTableResults{}
	if err := client.createObject(endpointDBTable, dbTable.Scope, dbTable, &dbTablePageResults); err != nil {
		return nil, err
	}

	return &dbTablePageResults.Records[0], nil
}

// UpdateDBTable updates a Table in ServiceNow.
func (client *ServiceNowClient) UpdateDBTable(dbTable *DBTable) error {
	return client.updateObject(endpointDBTable, dbTable.Id, dbTable)
}

// DeleteDBTable deletes a Table in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteDBTable(id string) error {
	return client.deleteObject(endpointDBTable, id)
}

func (results DBTableResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
