package resources

import (
	"fmt"
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const jsIncludeSource = "source"
const jsIncludeDisplayName = "display_name"
const jsIncludeUrl = "url"
const jsIncludeUiScriptId = "ui_script_id"

// ResourceJsInclude is holding the info about a javascript script to be included.
func ResourceJsInclude() *schema.Resource {
	return &schema.Resource{
		Create: createResourceJsInclude,
		Read:   readResourceJsInclude,
		Update: updateResourceJsInclude,
		Delete: deleteResourceJsInclude,

		Schema: map[string]*schema.Schema{
			jsIncludeSource: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "url",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v != "url" && v != "local" {
						errs = append(errs, fmt.Errorf("%q must be 'url' or 'local', got: %s", key, v))
					}
					return
				},
			},
			jsIncludeDisplayName: {
				Type:     schema.TypeString,
				Required: true,
			},
			jsIncludeUrl: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			jsIncludeUiScriptId: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func readResourceJsInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	jsInclude, err := client.GetJsInclude(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromJsInclude(data, jsInclude)

	return nil
}

func createResourceJsInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	createdPage, err := client.CreateJsInclude(resourceToJsInclude(data))
	if err != nil {
		return err
	}

	resourceFromJsInclude(data, createdPage)

	return readResourceJsInclude(data, serviceNowClient)
}

func updateResourceJsInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateJsInclude(resourceToJsInclude(data))
	if err != nil {
		return err
	}

	return readResourceJsInclude(data, serviceNowClient)
}

func deleteResourceJsInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteJsInclude(data.Id())
}

func resourceFromJsInclude(data *schema.ResourceData, jsInclude *client.JsInclude) {
	data.SetId(jsInclude.Id)
	data.Set(jsIncludeSource, jsInclude.Source)
	data.Set(jsIncludeDisplayName, jsInclude.DisplayName)
	data.Set(jsIncludeUrl, jsInclude.Url)
	data.Set(jsIncludeUiScriptId, jsInclude.UiScriptId)
}

func resourceToJsInclude(data *schema.ResourceData) *client.JsInclude {
	jsInclude := client.JsInclude{
		Source:      data.Get(jsIncludeSource).(string),
		DisplayName: data.Get(jsIncludeDisplayName).(string),
		Url:         data.Get(jsIncludeUrl).(string),
		UiScriptId:  data.Get(jsIncludeUiScriptId).(string),
	}
	jsInclude.Id = data.Id()
	return &jsInclude
}
