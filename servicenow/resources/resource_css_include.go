package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const cssIncludeSource = "source"
const cssIncludeName = "name"
const cssIncludeURL = "url"
const cssIncludeStyleSheetID = "style_sheet_id"

// ResourceCSSInclude is holding the info about a javascript script to be included.
func ResourceCSSInclude() *schema.Resource {
	return &schema.Resource{
		Create: createResourceCSSInclude,
		Read:   readResourceCSSInclude,
		Update: updateResourceCSSInclude,
		Delete: deleteResourceCSSInclude,

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
			cssIncludeURL: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			cssIncludeStyleSheetID: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The ID of the service portal style sheet to include.",
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceCSSInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	cssInclude := &client.CSSInclude{}
	if err := snowClient.GetObject(client.EndpointCSSInclude, data.Id(), cssInclude); err != nil {
		data.SetId("")
		return err
	}

	resourceFromCSSInclude(data, cssInclude)

	return nil
}

func createResourceCSSInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	cssInclude := resourceToCSSInclude(data)
	if err := snowClient.CreateObject(client.EndpointCSSInclude, cssInclude); err != nil {
		return err
	}

	resourceFromCSSInclude(data, cssInclude)

	return readResourceCSSInclude(data, serviceNowClient)
}

func updateResourceCSSInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointCSSInclude, resourceToCSSInclude(data)); err != nil {
		return err
	}

	return readResourceCSSInclude(data, serviceNowClient)
}

func deleteResourceCSSInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointCSSInclude, data.Id())
}

func resourceFromCSSInclude(data *schema.ResourceData, cssInclude *client.CSSInclude) {
	data.SetId(cssInclude.ID)
	data.Set(cssIncludeSource, cssInclude.Source)
	data.Set(cssIncludeName, cssInclude.Name)
	data.Set(cssIncludeURL, cssInclude.URL)
	data.Set(cssIncludeStyleSheetID, cssInclude.StyleSheetID)
	data.Set(commonScope, cssInclude.Scope)
}

func resourceToCSSInclude(data *schema.ResourceData) *client.CSSInclude {
	cssInclude := client.CSSInclude{
		Source:       data.Get(cssIncludeSource).(string),
		Name:         data.Get(cssIncludeName).(string),
		URL:          data.Get(cssIncludeURL).(string),
		StyleSheetID: data.Get(cssIncludeStyleSheetID).(string),
	}
	cssInclude.ID = data.Id()
	cssInclude.Scope = data.Get(commonScope).(string)
	return &cssInclude
}
