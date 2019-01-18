package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

// DataSourceDBTable reads a DB Table in ServiceNow.
func DataSourceDBTable() *schema.Resource {
	// Copy the schema from the resource.
	resourceSchema := ResourceDBTable().Schema
	setOnlyRequiredSchema(resourceSchema, dbTableName)

	return &schema.Resource{
		Read:   readDataSourceDBTable,
		Schema: resourceSchema,
	}
}

func readDataSourceDBTable(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	dbTable := &client.DBTable{}
	if err := snowClient.GetObjectByName(client.EndpointDBTable, data.Get(dbTableName).(string), dbTable); err != nil {
		data.SetId("")
		return err
	}

	resourceFromDBTable(data, dbTable)

	return nil
}
