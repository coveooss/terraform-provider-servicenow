package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

func DataSourceDBTable() *schema.Resource {
	// Copy the schema from the resource.
	resourceSchema := ResourceDBTable().Schema

	// Change required parameters. For the data source, name is required and everything else is computed.
	for _, schema := range resourceSchema {
		schema.Computed = true
		schema.Required = false
		schema.Optional = false
		schema.Default = nil
	}

	resourceSchema[dbTableName].Computed = false
	resourceSchema[dbTableName].Required = true

	return &schema.Resource{
		Read:   readDataSourceDBTable,
		Schema: resourceSchema,
	}
}

func readDataSourceDBTable(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	dbTable, err := client.GetDBTableByName(data.Get(dbTableName).(string))
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromDBTable(data, dbTable)

	return nil
}
