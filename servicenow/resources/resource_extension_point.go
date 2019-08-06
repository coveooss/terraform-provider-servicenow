package resources

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const extensionPointName = "name"
const extensionPointDescription = "description"
const extensionPointRestrictScope = "restrict_scope"
const extensionPointExample = "example"
const extensionPointAPIName = "api_name"

// ResourceExtensionPoint is holding the info about a scripted extension point.
func ResourceExtensionPoint() *schema.Resource {
	return &schema.Resource{
		Create: createResourceExtensionPoint,
		Read:   readResourceExtensionPoint,
		Update: updateResourceExtensionPoint,
		Delete: deleteResourceExtensionPoint,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			extensionPointName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique name of the extension point.",
			},
			extensionPointDescription: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Description of the required implementation of the extension point.",
			},
			extensionPointExample: {
				Type:        schema.TypeString,
				Required:    true,
				Default:     "",
				Description: "Example implementation code.",
			},
			extensionPointRestrictScope: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Only allow extension instances within this point's scope.",
			},
			extensionPointAPIName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the extension point API, that is pre-pended with the application scope to which it applies. This is a system-assigned name and cannot be changed.",
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceExtensionPoint(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	extensionPoint := &client.ExtensionPoint{}
	if err := snowClient.GetObject(client.EndpointExtensionPoint, data.Id(), extensionPoint); err != nil {
		data.SetId("")
		return err
	}

	resourceFromExtensionPoint(data, extensionPoint)

	return nil
}

func createResourceExtensionPoint(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	extensionPoint := resourceToExtensionPoint(data)
	if err := snowClient.CreateObject(client.EndpointExtensionPoint, extensionPoint); err != nil {
		return err
	}

	resourceFromExtensionPoint(data, extensionPoint)

	return readResourceExtensionPoint(data, serviceNowClient)
}

func updateResourceExtensionPoint(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointExtensionPoint, resourceToExtensionPoint(data)); err != nil {
		return err
	}

	return readResourceExtensionPoint(data, serviceNowClient)
}

func deleteResourceExtensionPoint(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointExtensionPoint, data.Id())
}

func resourceFromExtensionPoint(data *schema.ResourceData, extensionPoint *client.ExtensionPoint) {
	data.SetId(extensionPoint.ID)
	data.Set(extensionPointName, extensionPoint.Name)
	data.Set(extensionPointDescription, extensionPoint.Description)
	data.Set(extensionPointRestrictScope, extensionPoint.RestrictScope)
	data.Set(extensionPointExample, extensionPoint.Example)
	data.Set(extensionPointAPIName, extensionPoint.APIName)
	data.Set(commonScope, extensionPoint.Scope)
}

func resourceToExtensionPoint(data *schema.ResourceData) *client.ExtensionPoint {
	extensionPoint := client.ExtensionPoint{
		Name:          data.Get(extensionPointName).(string),
		Description:   data.Get(extensionPointDescription).(string),
		RestrictScope: data.Get(extensionPointRestrictScope).(bool),
		Example:       data.Get(extensionPointExample).(string),
	}
	extensionPoint.ID = data.Id()
	extensionPoint.Scope = data.Get(commonScope).(string)
	return &extensionPoint
}
