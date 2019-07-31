package resources

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const contentCSSName = "name"
const contentCSSType = "type"
const contentCSSUrl = "url"
const contentCSSStyle = "style"

// ResourceContentCSS is holding the info about a style sheet to be included.
func ResourceContentCSS() *schema.Resource {
	return &schema.Resource{
		Create: createResourceContentCSS,
		Read:   readResourceContentCSS,
		Update: updateResourceContentCSS,
		Delete: deleteResourceContentCSS,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			contentCSSName: {
				Type:     schema.TypeString,
				Required: true,
			},
			contentCSSType: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "local",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					warns, errs = validateStringValue(val.(string), key, []string{"link", "local"})
					return
				},
				Description: "The type of this content management style sheet. Can be 'link' or 'local'.",
			},
			contentCSSUrl: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Used when 'type' is set to 'link'. Must be an absolute URL to an external style sheet file.",
			},
			contentCSSStyle: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Used when 'type' is set to 'local'. The raw CSS content of this style sheet.",
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceContentCSS(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	contentCSS := &client.ContentCSS{}
	if err := snowClient.GetObject(client.EndpointContentCSS, data.Id(), contentCSS); err != nil {
		data.SetId("")
		return err
	}

	resourceFromContentCSS(data, contentCSS)

	return nil
}

func createResourceContentCSS(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	contentCSS := resourceToContentCSS(data)
	if err := snowClient.CreateObject(client.EndpointContentCSS, contentCSS); err != nil {
		return err
	}

	resourceFromContentCSS(data, contentCSS)

	return readResourceContentCSS(data, serviceNowClient)
}

func updateResourceContentCSS(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointContentCSS, resourceToContentCSS(data)); err != nil {
		return err
	}

	return readResourceContentCSS(data, serviceNowClient)
}

func deleteResourceContentCSS(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointContentCSS, data.Id())
}

func resourceFromContentCSS(data *schema.ResourceData, contentCSS *client.ContentCSS) {
	data.SetId(contentCSS.ID)
	data.Set(contentCSSName, contentCSS.Name)
	data.Set(contentCSSType, contentCSS.Type)
	data.Set(contentCSSUrl, contentCSS.URL)
	data.Set(contentCSSStyle, contentCSS.Style)
	data.Set(commonScope, contentCSS.Scope)
}

func resourceToContentCSS(data *schema.ResourceData) *client.ContentCSS {
	contentCSS := client.ContentCSS{
		Name:  data.Get(contentCSSName).(string),
		Type:  data.Get(contentCSSType).(string),
		URL:   data.Get(contentCSSUrl).(string),
		Style: data.Get(contentCSSStyle).(string),
	}
	contentCSS.ID = data.Id()
	contentCSS.Scope = data.Get(commonScope).(string)
	return &contentCSS
}
