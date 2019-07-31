package resources

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const jsIncludeSource = "source"
const jsIncludeDisplayName = "display_name"
const jsIncludeURL = "url"
const jsIncludeUIScriptID = "ui_script_id"

// ResourceJsInclude is holding the info about a javascript script to be included.
func ResourceJsInclude() *schema.Resource {
	return &schema.Resource{
		Create: createResourceJsInclude,
		Read:   readResourceJsInclude,
		Update: updateResourceJsInclude,
		Delete: deleteResourceJsInclude,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			jsIncludeSource: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "url",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					warns, errs = validateStringValue(val.(string), key, []string{"url", "local"})
					return
				},
			},
			jsIncludeDisplayName: {
				Type:     schema.TypeString,
				Required: true,
			},
			jsIncludeURL: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			jsIncludeUIScriptID: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceJsInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	jsInclude := &client.JsInclude{}
	if err := snowClient.GetObject(client.EndpointJsInclude, data.Id(), jsInclude); err != nil {
		data.SetId("")
		return err
	}

	resourceFromJsInclude(data, jsInclude)

	return nil
}

func createResourceJsInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	jsInclude := resourceToJsInclude(data)
	if err := snowClient.CreateObject(client.EndpointJsInclude, jsInclude); err != nil {
		return err
	}

	resourceFromJsInclude(data, jsInclude)

	return readResourceJsInclude(data, serviceNowClient)
}

func updateResourceJsInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointJsInclude, resourceToJsInclude(data)); err != nil {
		return err
	}

	return readResourceJsInclude(data, serviceNowClient)
}

func deleteResourceJsInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointJsInclude, data.Id())
}

func resourceFromJsInclude(data *schema.ResourceData, jsInclude *client.JsInclude) {
	data.SetId(jsInclude.ID)
	data.Set(jsIncludeSource, jsInclude.Source)
	data.Set(jsIncludeDisplayName, jsInclude.DisplayName)
	data.Set(jsIncludeURL, jsInclude.URL)
	data.Set(jsIncludeUIScriptID, jsInclude.UIScriptID)
	data.Set(commonScope, jsInclude.Scope)
}

func resourceToJsInclude(data *schema.ResourceData) *client.JsInclude {
	jsInclude := client.JsInclude{
		Source:      data.Get(jsIncludeSource).(string),
		DisplayName: data.Get(jsIncludeDisplayName).(string),
		URL:         data.Get(jsIncludeURL).(string),
		UIScriptID:  data.Get(jsIncludeUIScriptID).(string),
	}
	jsInclude.ID = data.Id()
	jsInclude.Scope = data.Get(commonScope).(string)
	return &jsInclude
}
