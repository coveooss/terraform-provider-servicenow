package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const widgetDepRelationDependencyID = "dependency_id"
const widgetDepRelationWidgetID = "widget_id"

// ResourceWidgetDependencyRelation is holding the relationship between a widget and a widget dependency (many-2-many).
func ResourceWidgetDependencyRelation() *schema.Resource {
	return &schema.Resource{
		Create: createResourceWidgetDepRelation,
		Read:   readResourceWidgetDepRelation,
		Update: updateResourceWidgetDepRelation,
		Delete: deleteResourceWidgetDepRelation,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			widgetDepRelationDependencyID: {
				Type:     schema.TypeString,
				Required: true,
			},
			widgetDepRelationWidgetID: {
				Type:     schema.TypeString,
				Required: true,
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceWidgetDepRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	relation := &client.WidgetDependencyRelation{}
	if err := snowClient.GetObject(client.EndpointWidgetDependencyRelation, data.Id(), relation); err != nil {
		data.SetId("")
		return err
	}

	resourceFromWidgetDepRelation(data, relation)

	return nil
}

func createResourceWidgetDepRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	relation := resourceToWidgetDepRelation(data)
	if err := snowClient.CreateObject(client.EndpointWidgetDependencyRelation, relation); err != nil {
		return err
	}

	resourceFromWidgetDepRelation(data, relation)

	return readResourceWidgetDepRelation(data, serviceNowClient)
}

func updateResourceWidgetDepRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointWidgetDependencyRelation, resourceToWidgetDepRelation(data)); err != nil {
		return err
	}

	return readResourceWidgetDepRelation(data, serviceNowClient)
}

func deleteResourceWidgetDepRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointWidgetDependencyRelation, data.Id())
}

func resourceFromWidgetDepRelation(data *schema.ResourceData, relation *client.WidgetDependencyRelation) {
	data.SetId(relation.ID)
	data.Set(widgetDepRelationDependencyID, relation.DependencyID)
	data.Set(widgetDepRelationWidgetID, relation.WidgetID)
	data.Set(commonScope, relation.Scope)
}

func resourceToWidgetDepRelation(data *schema.ResourceData) *client.WidgetDependencyRelation {
	relation := client.WidgetDependencyRelation{
		DependencyID: data.Get(widgetDepRelationDependencyID).(string),
		WidgetID:     data.Get(widgetDepRelationWidgetID).(string),
	}
	relation.ID = data.Id()
	relation.Scope = data.Get(commonScope).(string)
	return &relation
}
