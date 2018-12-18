package client

import (
	"encoding/json"
	"fmt"
)

const endpointRole = "sys_user_role.do"

// Role is the json response for a UI role in ServiceNow.
type Role struct {
	BaseResult
	Name              string `json:"name"`
	Description       string `json:"description"`
	ElevatedPrivilege bool   `json:"elevated_privilege,string"`
	Suffix            string `json:"suffix"`
	AssignableBy      string `json:"assignable_by"`
}

// RoleResults is the object returned by ServiceNow API when saving or retrieving records.
type RoleResults struct {
	Records []Role `json:"records"`
}

// GetRole retrieves a specific Role in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetRole(id string) (*Role, error) {
	rolePageResults := RoleResults{}
	if err := client.getObject(endpointRole, id, &rolePageResults); err != nil {
		return nil, err
	}

	return &rolePageResults.Records[0], nil
}

// GetRoleByName retrieves a specific Role in ServiceNow with it's name attribute.
func (client *ServiceNowClient) GetRoleByName(name string) (*Role, error) {
	jsonResponse, err := client.requestJSON("GET", endpointRole+"?JSONv2&sysparm_query=name="+name, nil)
	if err != nil {
		return nil, err
	}

	rolePageResults := RoleResults{}
	if err := json.Unmarshal(jsonResponse, &rolePageResults); err != nil {
		return nil, err
	}

	if err := rolePageResults.validate(); err != nil {
		return nil, err
	}

	return &rolePageResults.Records[0], nil
}

// CreateRole creates a new Role in ServiceNow and returns the newly created role. The new role should
// include the GUID (sys_id) created in ServiceNow.
func (client *ServiceNowClient) CreateRole(role *Role) (*Role, error) {
	rolePageResults := RoleResults{}
	if err := client.createObject(endpointRole, role, &rolePageResults); err != nil {
		return nil, err
	}

	return &rolePageResults.Records[0], nil
}

// UpdateRole updates a Role in ServiceNow.
func (client *ServiceNowClient) UpdateRole(role *Role) error {
	return client.updateObject(endpointRole, role.Id, role)
}

// DeleteRole deletes a Role in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteRole(id string) error {
	return client.deleteObject(endpointRole, id)
}

func (results RoleResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
