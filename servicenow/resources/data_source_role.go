package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func DataSourceRole() *schema.Resource {
	// Copy the schema from the resource.
	resourceSchema := ResourceRole().Schema
	setOnlyRequiredSchema(resourceSchema, roleName)

	return &schema.Resource{
		Read:   readDataSourceRole,
		Schema: resourceSchema,
	}
}

func readDataSourceRole(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	role, err := client.GetRoleByName(data.Get(roleName).(string))
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromRole(data, role)

	return nil
}
