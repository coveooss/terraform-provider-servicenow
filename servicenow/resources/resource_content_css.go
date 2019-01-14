package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const contentCssName = "name"
const contentCssType = "type"
const contentCssUrl = "url"
const contentCssStyle = "style"

// ResourceContentCss is holding the info about a style sheet to be included.
func ResourceContentCss() *schema.Resource {
	return &schema.Resource{
		Create: createResourceContentCss,
		Read:   readResourceContentCss,
		Update: updateResourceContentCss,
		Delete: deleteResourceContentCss,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			contentCssName: {
				Type:     schema.TypeString,
				Required: true,
			},
			contentCssType: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "local",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					warns, errs = validateStringValue(val.(string), key, []string{"url", "local"})
					return
				},
				Description: "The type of this content management style sheet. Can be 'url' or 'local'.",
			},
			contentCssUrl: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Used when 'type' is set to 'url'. Must be an absolute url to an external style sheet file.",
			},
			contentCssStyle: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Used when 'type' is set to 'local'. The raw CSS content of this style sheet.",
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceContentCss(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	contentCss, err := client.GetContentCss(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromContentCss(data, contentCss)

	return nil
}

func createResourceContentCss(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	contentCss, err := client.CreateContentCss(resourceToContentCss(data))
	if err != nil {
		return err
	}

	resourceFromContentCss(data, contentCss)

	return readResourceContentCss(data, serviceNowClient)
}

func updateResourceContentCss(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateContentCss(resourceToContentCss(data))
	if err != nil {
		return err
	}

	return readResourceContentCss(data, serviceNowClient)
}

func deleteResourceContentCss(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteContentCss(data.Id())
}

func resourceFromContentCss(data *schema.ResourceData, contentCss *client.ContentCss) {
	data.SetId(contentCss.Id)
	data.Set(contentCssName, contentCss.Name)
	data.Set(contentCssType, contentCss.Type)
	data.Set(contentCssUrl, contentCss.URL)
	data.Set(contentCssStyle, contentCss.Style)
	data.Set(commonScope, contentCss.Scope)
}

func resourceToContentCss(data *schema.ResourceData) *client.ContentCss {
	contentCss := client.ContentCss{
		Name:  data.Get(contentCssName).(string),
		Type:  data.Get(contentCssType).(string),
		URL:   data.Get(contentCssUrl).(string),
		Style: data.Get(contentCssStyle).(string),
	}
	contentCss.Id = data.Id()
	contentCss.Scope = data.Get(commonScope).(string)
	return &contentCss
}
