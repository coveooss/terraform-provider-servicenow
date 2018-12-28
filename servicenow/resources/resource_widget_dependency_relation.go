package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const widgetDepRelationDependencyId = "dependency_id"
const widgetDepRelationWidgetId = "widget_id"

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
			widgetDepRelationDependencyId: {
				Type:     schema.TypeString,
				Required: true,
			},
			widgetDepRelationWidgetId: {
				Type:     schema.TypeString,
				Required: true,
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceWidgetDepRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	relation, err := client.GetWidgetDependencyRelation(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromWidgetDepRelation(data, relation)

	return nil
}

func createResourceWidgetDepRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	createdPage, err := client.CreateWidgetDependencyRelation(resourceToWidgetDepRelation(data))
	if err != nil {
		return err
	}

	resourceFromWidgetDepRelation(data, createdPage)

	return readResourceWidgetDepRelation(data, serviceNowClient)
}

func updateResourceWidgetDepRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateWidgetDependencyRelation(resourceToWidgetDepRelation(data))
	if err != nil {
		return err
	}

	return readResourceWidgetDepRelation(data, serviceNowClient)
}

func deleteResourceWidgetDepRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteWidgetDependencyRelation(data.Id())
}

func resourceFromWidgetDepRelation(data *schema.ResourceData, relation *client.WidgetDependencyRelation) {
	data.SetId(relation.Id)
	data.Set(widgetDepRelationDependencyId, relation.DependencyId)
	data.Set(widgetDepRelationWidgetId, relation.WidgetId)
	data.Set(commonScope, relation.Scope)
}

func resourceToWidgetDepRelation(data *schema.ResourceData) *client.WidgetDependencyRelation {
	relation := client.WidgetDependencyRelation{
		DependencyId: data.Get(widgetDepRelationDependencyId).(string),
		WidgetId:     data.Get(widgetDepRelationWidgetId).(string),
	}
	relation.Id = data.Id()
	relation.Scope = data.Get(commonScope).(string)
	return &relation
}
