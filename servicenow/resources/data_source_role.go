package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

// DataSourceRole reads the informations about a single Role in ServiceNow.
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
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	role := &client.Role{}
	if err := snowClient.GetObjectByName(client.EndpointRole, data.Get(roleName).(string), role); err != nil {
		data.SetId("")
		return err
	}

	resourceFromRole(data, role)

	return nil
}
