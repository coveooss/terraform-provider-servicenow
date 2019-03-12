package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const widgetDependencyName = "name"
const widgetDependencyModule = "module"
const widgetDependencyPageLoad = "page_load"

// ResourceWidgetDependency is holding the info about a javascript script to be included.
func ResourceWidgetDependency() *schema.Resource {
	return &schema.Resource{
		Create: createResourceWidgetDependency,
		Read:   readResourceWidgetDependency,
		Update: updateResourceWidgetDependency,
		Delete: deleteResourceWidgetDependency,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			widgetDependencyName: {
				Type:     schema.TypeString,
				Required: true,
			},
			widgetDependencyModule: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			widgetDependencyPageLoad: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceWidgetDependency(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	widgetDependency := &client.WidgetDependency{}
	if err := snowClient.GetObject(client.EndpointWidgetDependency, data.Id(), widgetDependency); err != nil {
		data.SetId("")
		return err
	}

	resourceFromWidgetDependency(data, widgetDependency)

	return nil
}

func createResourceWidgetDependency(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	widgetDependency := resourceToWidgetDependency(data)
	if err := snowClient.CreateObject(client.EndpointWidgetDependency, widgetDependency); err != nil {
		return err
	}

	resourceFromWidgetDependency(data, widgetDependency)

	return readResourceWidgetDependency(data, serviceNowClient)
}

func updateResourceWidgetDependency(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointWidgetDependency, resourceToWidgetDependency(data)); err != nil {
		return err
	}

	return readResourceWidgetDependency(data, serviceNowClient)
}

func deleteResourceWidgetDependency(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointWidgetDependency, data.Id())
}

func resourceFromWidgetDependency(data *schema.ResourceData, widgetDependency *client.WidgetDependency) {
	data.SetId(widgetDependency.ID)
	data.Set(widgetDependencyName, widgetDependency.Name)
	data.Set(widgetDependencyModule, widgetDependency.Module)
	data.Set(widgetDependencyPageLoad, widgetDependency.PageLoad)
	data.Set(commonScope, widgetDependency.Scope)
}

func resourceToWidgetDependency(data *schema.ResourceData) *client.WidgetDependency {
	widgetDependency := client.WidgetDependency{
		Name:     data.Get(widgetDependencyName).(string),
		Module:   data.Get(widgetDependencyModule).(string),
		PageLoad: data.Get(widgetDependencyPageLoad).(bool),
	}
	widgetDependency.ID = data.Id()
	widgetDependency.Scope = data.Get(commonScope).(string)
	return &widgetDependency
}
