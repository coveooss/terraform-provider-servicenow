package resources

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const systemPropertyRelationCategoryID = "category_id"
const systemPropertyRelationPropertyID = "property_id"
const systemPropertyRelationOrder = "order"

// ResourceSystemPropertyRelation manages a System Property in ServiceNow.
func ResourceSystemPropertyRelation() *schema.Resource {
	return &schema.Resource{
		Create: createResourceSystemPropertyRelation,
		Read:   readResourceSystemPropertyRelation,
		Update: updateResourceSystemPropertyRelation,
		Delete: deleteResourceSystemPropertyRelation,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			systemPropertyRelationCategoryID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "System Property Category ID to link.",
			},
			systemPropertyRelationPropertyID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the System Property to link.",
			},
			systemPropertyRelationOrder: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1",
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceSystemPropertyRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	systemPropertyRelation := &client.SystemPropertyRelation{}
	if err := snowClient.GetObject(client.EndpointSystemPropertyRelation, data.Id(), systemPropertyRelation); err != nil {
		data.SetId("")
		return err
	}

	resourceFromSystemPropertyRelation(data, systemPropertyRelation)

	return nil
}

func createResourceSystemPropertyRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	systemPropertyRelation := resourceToSystemPropertyRelation(data)
	if err := snowClient.CreateObject(client.EndpointSystemPropertyRelation, systemPropertyRelation); err != nil {
		return err
	}

	resourceFromSystemPropertyRelation(data, systemPropertyRelation)

	return readResourceSystemPropertyRelation(data, serviceNowClient)
}

func updateResourceSystemPropertyRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointSystemPropertyRelation, resourceToSystemPropertyRelation(data)); err != nil {
		return err
	}

	return readResourceSystemPropertyRelation(data, serviceNowClient)
}

func deleteResourceSystemPropertyRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointSystemPropertyRelation, data.Id())
}

func resourceFromSystemPropertyRelation(data *schema.ResourceData, systemPropertyRelation *client.SystemPropertyRelation) {
	data.SetId(systemPropertyRelation.ID)
	data.Set(systemPropertyRelationCategoryID, systemPropertyRelation.CategoryID)
	data.Set(systemPropertyRelationPropertyID, systemPropertyRelation.PropertyID)
	data.Set(systemPropertyRelationOrder, systemPropertyRelation.Order)
	data.Set(commonScope, systemPropertyRelation.Scope)
}

func resourceToSystemPropertyRelation(data *schema.ResourceData) *client.SystemPropertyRelation {
	systemPropertyRelation := client.SystemPropertyRelation{
		CategoryID: data.Get(systemPropertyRelationCategoryID).(string),
		PropertyID: data.Get(systemPropertyRelationPropertyID).(string),
		Order:      data.Get(systemPropertyRelationOrder).(string),
	}
	systemPropertyRelation.ID = data.Id()
	systemPropertyRelation.Scope = data.Get(commonScope).(string)
	return &systemPropertyRelation
}
