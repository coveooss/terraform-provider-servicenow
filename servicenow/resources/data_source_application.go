package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

// DataSourceApplication reads an Application in ServiceNow.
func DataSourceApplication() *schema.Resource {
	// Copy the schema from the resource.
	resourceSchema := ResourceApplication().Schema

	// Change required parameters. For the data source, name is required and everything else is computed.
	for _, schema := range resourceSchema {
		schema.Computed = true
		schema.Required = false
		schema.Optional = false
		schema.Default = nil
	}

	resourceSchema[applicationName].Computed = false
	resourceSchema[applicationName].Required = true

	return &schema.Resource{
		Read:   readDataSourceApplication,
		Schema: resourceSchema,
	}
}

func readDataSourceApplication(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	application, err := client.GetApplicationByName(data.Get(applicationName).(string))
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromApplication(data, application)

	return nil
}
