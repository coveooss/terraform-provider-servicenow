package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const restMethodName = "name"
const restMethodMessageID = "rest_message_id"
const restMethodHTTPMethod = "http_method"
const restMethodRestEndpoint = "rest_endpoint"
const restMethodAuthenticationType = "authentication_type"
const restMethodQualifiedName = "qualified_name"

// ResourceRestMethod is holding the info about a REST method to be included in a REST message.
func ResourceRestMethod() *schema.Resource {
	return &schema.Resource{
		Create: createResourceRestMethod,
		Read:   readResourceRestMethod,
		Update: updateResourceRestMethod,
		Delete: deleteResourceRestMethod,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			restMethodName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A unique identifier for this HTTP method record.",
			},
			restMethodMessageID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The REST message record ID this method is based on.",
			},
			restMethodHTTPMethod: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The HTTP method this record implements. Can be 'get', 'post', 'put', 'patch' or 'delete'.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					warns, errs = validateStringValue(val.(string), key, []string{"get", "post", "put", "patch", "delete"})
					return
				},
			},
			restMethodRestEndpoint: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The URL of the REST web service provider this method sends requests to. Can contain variables in the format '${variable}'.",
			},
			restMethodQualifiedName: {
				Type:     schema.TypeString,
				Computed: true,
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceRestMethod(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	restMethod := &client.RestMethod{}
	if err := snowClient.GetObject(client.EndpointRestMethod, data.Id(), restMethod); err != nil {
		data.SetId("")
		return err
	}

	resourceFromRestMethod(data, restMethod)

	return nil
}

func createResourceRestMethod(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	restMethod := resourceToRestMethod(data)
	if err := snowClient.CreateObject(client.EndpointRestMethod, restMethod); err != nil {
		return err
	}

	resourceFromRestMethod(data, restMethod)

	return readResourceRestMethod(data, serviceNowClient)
}

func updateResourceRestMethod(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointRestMethod, resourceToRestMethod(data)); err != nil {
		return err
	}

	return readResourceRestMethod(data, serviceNowClient)
}

func deleteResourceRestMethod(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointRestMethod, data.Id())
}

func resourceFromRestMethod(data *schema.ResourceData, restMethod *client.RestMethod) {
	data.SetId(restMethod.ID)
	data.Set(restMethodName, restMethod.Name)
	data.Set(restMethodMessageID, restMethod.MessageID)
	data.Set(restMethodHTTPMethod, restMethod.HTTPMethod)
	data.Set(restMethodRestEndpoint, restMethod.RestEndpoint)
	data.Set(restMethodQualifiedName, restMethod.QualifiedName)
	data.Set(commonScope, restMethod.Scope)
}

func resourceToRestMethod(data *schema.ResourceData) *client.RestMethod {
	restMethod := client.RestMethod{
		Name:               data.Get(restMethodName).(string),
		MessageID:          data.Get(restMethodMessageID).(string),
		HTTPMethod:         data.Get(restMethodHTTPMethod).(string),
		RestEndpoint:       data.Get(restMethodRestEndpoint).(string),
		AuthenticationType: "inherit_from_parent",
	}
	restMethod.ID = data.Id()
	restMethod.Scope = data.Get(commonScope).(string)
	return &restMethod
}
