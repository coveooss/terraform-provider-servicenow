package resources

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

// DataSourceSystemProperty reads the informations about a single SystemProperty in ServiceNow.
func DataSourceSystemProperty() *schema.Resource {
	// Copy the schema from the resource.
	resourceSchema := ResourceSystemProperty().Schema
	setOnlyRequiredSchema(resourceSchema, systemPropertyName)

	return &schema.Resource{
		Read:   readDataSourceSystemProperty,
		Schema: resourceSchema,
	}
}

func readDataSourceSystemProperty(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	systemProperty := &client.SystemProperty{}
	if err := snowClient.GetObjectByName(client.EndpointSystemProperty, data.Get(systemPropertyName).(string), systemProperty); err != nil {
		data.SetId("")
		return err
	}

	resourceFromSystemProperty(data, systemProperty)

	return nil
}
