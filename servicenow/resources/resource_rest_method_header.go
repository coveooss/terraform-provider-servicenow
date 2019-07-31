package resources

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const restMethodHeaderName = "name"
const restMethodHeaderValue = "value"
const restMethodHeaderMethodID = "rest_method_id"

// ResourceRestMethodHeader is holding the info about a header to be applied to a REST method.
func ResourceRestMethodHeader() *schema.Resource {
	return &schema.Resource{
		Create: createResourceRestMethodHeader,
		Read:   readResourceRestMethodHeader,
		Update: updateResourceRestMethodHeader,
		Delete: deleteResourceRestMethodHeader,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			restMethodHeaderName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the header to add to the HTTP request.",
			},
			restMethodHeaderValue: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The value of the header to add to the HTTP request.",
			},
			restMethodHeaderMethodID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The REST method record ID this header will be applied to.",
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceRestMethodHeader(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	restMethodHeader := &client.RestMethodHeader{}
	if err := snowClient.GetObject(client.EndpointRestMethodHeader, data.Id(), restMethodHeader); err != nil {
		data.SetId("")
		return err
	}

	resourceFromRestMethodHeader(data, restMethodHeader)

	return nil
}

func createResourceRestMethodHeader(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	restMethodHeader := resourceToRestMethodHeader(data)
	if err := snowClient.CreateObject(client.EndpointRestMethodHeader, restMethodHeader); err != nil {
		return err
	}

	resourceFromRestMethodHeader(data, restMethodHeader)

	return readResourceRestMethodHeader(data, serviceNowClient)
}

func updateResourceRestMethodHeader(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointRestMethodHeader, resourceToRestMethodHeader(data)); err != nil {
		return err
	}

	return readResourceRestMethodHeader(data, serviceNowClient)
}

func deleteResourceRestMethodHeader(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointRestMethodHeader, data.Id())
}

func resourceFromRestMethodHeader(data *schema.ResourceData, restMethodHeader *client.RestMethodHeader) {
	data.SetId(restMethodHeader.ID)
	data.Set(restMethodHeaderName, restMethodHeader.Name)
	data.Set(restMethodHeaderValue, restMethodHeader.Value)
	data.Set(restMethodHeaderMethodID, restMethodHeader.MethodID)
	data.Set(commonScope, restMethodHeader.Scope)
}

func resourceToRestMethodHeader(data *schema.ResourceData) *client.RestMethodHeader {
	restMethodHeader := client.RestMethodHeader{
		Name:     data.Get(restMethodHeaderName).(string),
		Value:    data.Get(restMethodHeaderValue).(string),
		MethodID: data.Get(restMethodHeaderMethodID).(string),
	}
	restMethodHeader.ID = data.Id()
	restMethodHeader.Scope = data.Get(commonScope).(string)
	return &restMethodHeader
}
