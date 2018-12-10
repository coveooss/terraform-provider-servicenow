package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const dependencyId = "dependency_id"
const widgetId = "widget_id"

// ResourceWidgetDependencyRelation is holding the relationship between a widget and a widget dependency (many-2-many).
func ResourceWidgetDependencyRelation() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			dependencyId: {
				Type:     schema.TypeString,
				Required: true,
			},
			widgetId: {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServerRead(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	relation, err := client.GetWidgetDependencyRelation(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	updateResourceData(data, relation)
	return nil
}

func resourceServerCreate(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	createdPage, err := client.CreateWidgetDependencyRelation(resourceToRelation(data))
	if err != nil {
		return err
	}

	updateResourceData(data, createdPage)

	return resourceServerRead(data, serviceNowClient)
}

func resourceServerUpdate(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateWidgetDependencyRelation(resourceToRelation(data))
	if err != nil {
		return err
	}

	return resourceServerRead(data, serviceNowClient)
}

func resourceServerDelete(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteWidgetDependencyRelation(data.Id())
}

func updateResourceData(data *schema.ResourceData, relation *client.WidgetDependencyRelation) {
	data.SetId(relation.Id)
	data.Set(dependencyId, relation.DependencyId)
	data.Set(widgetId, relation.WidgetId)
}

func resourceToRelation(data *schema.ResourceData) *client.WidgetDependencyRelation {
	return &client.WidgetDependencyRelation{
		Id:           data.Id(),
		DependencyId: data.Get(dependencyId).(string),
		WidgetId:     data.Get(widgetId).(string),
	}
}
