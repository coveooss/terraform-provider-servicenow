package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const restMessageHeaderName = "name"
const restMessageHeaderValue = "value"
const restMessageHeaderMessageID = "rest_message_id"

// ResourceRestMessageHeader is holding the info about a header to be applied to a REST method.
func ResourceRestMessageHeader() *schema.Resource {
	return &schema.Resource{
		Create: createResourceRestMessageHeader,
		Read:   readResourceRestMessageHeader,
		Update: updateResourceRestMessageHeader,
		Delete: deleteResourceRestMessageHeader,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			restMessageHeaderName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the header to add to the HTTP request.",
			},
			restMessageHeaderValue: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The value of the header to add to the HTTP request.",
			},
			restMessageHeaderMessageID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The REST message record ID this header will be applied to.",
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceRestMessageHeader(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	restMessageHeader, err := client.GetRestMessageHeader(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromRestMessageHeader(data, restMessageHeader)

	return nil
}

func createResourceRestMessageHeader(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	createdPage, err := client.CreateRestMessageHeader(resourceToRestMessageHeader(data))
	if err != nil {
		return err
	}

	resourceFromRestMessageHeader(data, createdPage)

	return readResourceRestMessageHeader(data, serviceNowClient)
}

func updateResourceRestMessageHeader(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateRestMessageHeader(resourceToRestMessageHeader(data))
	if err != nil {
		return err
	}

	return readResourceRestMessageHeader(data, serviceNowClient)
}

func deleteResourceRestMessageHeader(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteRestMessageHeader(data.Id())
}

func resourceFromRestMessageHeader(data *schema.ResourceData, restMessageHeader *client.RestMessageHeader) {
	data.SetId(restMessageHeader.Id)
	data.Set(restMessageHeaderName, restMessageHeader.Name)
	data.Set(restMessageHeaderValue, restMessageHeader.Value)
	data.Set(restMessageHeaderMessageID, restMessageHeader.MessageID)
	data.Set(commonScope, restMessageHeader.Scope)
}

func resourceToRestMessageHeader(data *schema.ResourceData) *client.RestMessageHeader {
	restMessageHeader := client.RestMessageHeader{
		Name:      data.Get(restMessageHeaderName).(string),
		Value:     data.Get(restMessageHeaderValue).(string),
		MessageID: data.Get(restMessageHeaderMessageID).(string),
	}
	restMessageHeader.Id = data.Id()
	restMessageHeader.Scope = data.Get(commonScope).(string)
	return &restMessageHeader
}
