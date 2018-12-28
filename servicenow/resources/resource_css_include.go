package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const cssIncludeSource = "source"
const cssIncludeName = "name"
const cssIncludeUrl = "url"
const cssIncludeCssId = "css_id"

// ResourceCssInclude is holding the info about a javascript script to be included.
func ResourceCssInclude() *schema.Resource {
	return &schema.Resource{
		Create: createResourceCssInclude,
		Read:   readResourceCssInclude,
		Update: updateResourceCssInclude,
		Delete: deleteResourceCssInclude,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			cssIncludeSource: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "url",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					warns, errs = validateStringValue(val.(string), key, []string{"url", "local"})
					return
				},
			},
			cssIncludeName: {
				Type:     schema.TypeString,
				Required: true,
			},
			cssIncludeUrl: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			cssIncludeCssId: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceCssInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	cssInclude, err := client.GetCssInclude(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromCssInclude(data, cssInclude)

	return nil
}

func createResourceCssInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	createdPage, err := client.CreateCssInclude(resourceToCssInclude(data))
	if err != nil {
		return err
	}

	resourceFromCssInclude(data, createdPage)

	return readResourceCssInclude(data, serviceNowClient)
}

func updateResourceCssInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateCssInclude(resourceToCssInclude(data))
	if err != nil {
		return err
	}

	return readResourceCssInclude(data, serviceNowClient)
}

func deleteResourceCssInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteCssInclude(data.Id())
}

func resourceFromCssInclude(data *schema.ResourceData, cssInclude *client.CssInclude) {
	data.SetId(cssInclude.Id)
	data.Set(cssIncludeSource, cssInclude.Source)
	data.Set(cssIncludeName, cssInclude.Name)
	data.Set(cssIncludeUrl, cssInclude.Url)
	data.Set(cssIncludeCssId, cssInclude.CssId)
	data.Set(commonScope, cssInclude.Scope)
}

func resourceToCssInclude(data *schema.ResourceData) *client.CssInclude {
	cssInclude := client.CssInclude{
		Source: data.Get(cssIncludeSource).(string),
		Name:   data.Get(cssIncludeName).(string),
		Url:    data.Get(cssIncludeUrl).(string),
		CssId:  data.Get(cssIncludeCssId).(string),
	}
	cssInclude.Id = data.Id()
	cssInclude.Scope = data.Get(commonScope).(string)
	return &cssInclude
}
