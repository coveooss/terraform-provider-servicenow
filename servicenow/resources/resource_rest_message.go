package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const restMessageName = "name"
const restMessageDescription = "description"
const restMessageRestEndpoint = "rest_endpoint"
const restMessageAccess = "access"
const restMessageAuthenticationType = "authentication_type" // No auth is supported

// ResourceRestMessage is holding the info about a REST message configuration to be included.
func ResourceRestMessage() *schema.Resource {
	return &schema.Resource{
		Create: createResourceRestMessage,
		Read:   readResourceRestMessage,
		Update: updateResourceRestMessage,
		Delete: deleteResourceRestMessage,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			restMessageName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Descriptive name for this REST message.",
			},
			restMessageRestEndpoint: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL of the REST web service provider this REST message sends requests to.  Can contain variables in the format '${variable}'.",
			},
			restMessageDescription: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Description for this REST message.",
			},
			restMessageAccess: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "package_private",
				Description: "Whether this REST message can be accessed from only this application scope or all application scopes. Values can be 'package_private' or 'public'.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					warns, errs = validateStringValue(val.(string), key, []string{"package_private", "public"})
					return
				},
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceRestMessage(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	restMessage := &client.RestMessage{}
	if err := snowClient.GetObject(client.EndpointRestMessage, data.Id(), restMessage); err != nil {
		data.SetId("")
		return err
	}

	resourceFromRestMessage(data, restMessage)

	return nil
}

func createResourceRestMessage(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	restMessage := resourceToRestMessage(data)
	if err := snowClient.CreateObject(client.EndpointRestMessage, restMessage); err != nil {
		return err
	}

	resourceFromRestMessage(data, restMessage)

	return readResourceRestMessage(data, serviceNowClient)
}

func updateResourceRestMessage(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointRestMessage, resourceToRestMessage(data)); err != nil {
		return err
	}

	return readResourceRestMessage(data, serviceNowClient)
}

func deleteResourceRestMessage(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointRestMessage, data.Id())
}

func resourceFromRestMessage(data *schema.ResourceData, restMessage *client.RestMessage) {
	data.SetId(restMessage.ID)
	data.Set(restMessageName, restMessage.Name)
	data.Set(restMessageDescription, restMessage.Description)
	data.Set(restMessageRestEndpoint, restMessage.RestEndpoint)
	data.Set(restMessageAccess, restMessage.Access)
	data.Set(commonScope, restMessage.Scope)
}

func resourceToRestMessage(data *schema.ResourceData) *client.RestMessage {
	restMessage := client.RestMessage{
		Name:               data.Get(restMessageName).(string),
		Description:        data.Get(restMessageDescription).(string),
		RestEndpoint:       data.Get(restMessageRestEndpoint).(string),
		Access:             data.Get(restMessageAccess).(string),
		AuthenticationType: "no_authentication",
	}
	restMessage.ID = data.Id()
	restMessage.Scope = data.Get(commonScope).(string)
	return &restMessage
}
