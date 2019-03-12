package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const applicationName = "name"
const applicationScope = "scope"
const applicationVersion = "version"

// ResourceApplication manages an Application in ServiceNow.
func ResourceApplication() *schema.Resource {
	return &schema.Resource{
		Create: createResourceApplication,
		Read:   readResourceApplication,
		Update: updateResourceApplication,
		Delete: deleteResourceApplication,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			applicationName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the Application to retrieve from the ServiceNow instance.",
			},
			applicationScope: {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The unique scope of the application. Normally in the format x_[companyCode]_[shortAppId]. Cannot be changed once the application is created.",
			},
			applicationVersion: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1.0.0",
			},
		},
	}
}

func readResourceApplication(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	application := &client.Application{}
	if err := snowClient.GetObject(client.EndpointApplication, data.Id(), application); err != nil {
		data.SetId("")
		return err
	}

	resourceFromApplication(data, application)

	return nil
}

func createResourceApplication(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	application := resourceToApplication(data)
	if err := snowClient.CreateObject(client.EndpointApplication, application); err != nil {
		return err
	}

	resourceFromApplication(data, application)

	return readResourceApplication(data, serviceNowClient)
}

func updateResourceApplication(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointApplication, resourceToApplication(data)); err != nil {
		return err
	}

	return readResourceApplication(data, serviceNowClient)
}

func deleteResourceApplication(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointApplication, data.Id())
}

func resourceFromApplication(data *schema.ResourceData, application *client.Application) {
	data.SetId(application.ID)
	data.Set(applicationName, application.Name)
	data.Set(applicationScope, application.Scope)
	data.Set(applicationVersion, application.Version)
}

func resourceToApplication(data *schema.ResourceData) *client.Application {
	application := client.Application{
		Name:    data.Get(applicationName).(string),
		Scope:   data.Get(applicationScope).(string),
		Version: data.Get(applicationVersion).(string),
	}
	application.ID = data.Id()
	return &application
}
