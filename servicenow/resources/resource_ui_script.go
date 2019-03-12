package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const uiScriptName = "name"
const uiScriptDescription = "description"
const uiScriptScript = "script"
const uiScriptActive = "active"
const uiScriptUIType = "type"
const uiScriptAPIName = "api_name"

// ResourceUIScript manages a UI Script in ServiceNow which can be added to any other UI component.
func ResourceUIScript() *schema.Resource {
	return &schema.Resource{
		Create: createResourceUIScript,
		Read:   readResourceUIScript,
		Update: updateResourceUIScript,
		Delete: deleteResourceUIScript,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			uiScriptName: {
				Type:     schema.TypeString,
				Required: true,
			},
			uiScriptDescription: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			uiScriptScript: {
				Type:     schema.TypeString,
				Required: true,
			},
			uiScriptActive: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			uiScriptUIType: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "all",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					warns, errs = validateStringValue(val.(string), key, []string{"all", "desktop", "mobile"})
					return
				},
			},
			uiScriptAPIName: {
				Type:     schema.TypeString,
				Computed: true,
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceUIScript(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	uiScript := &client.UIScript{}
	if err := snowClient.GetObject(client.EndpointUIScript, data.Id(), uiScript); err != nil {
		data.SetId("")
		return err
	}

	resourceFromUIScript(data, uiScript)

	return nil
}

func createResourceUIScript(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	uiScript := resourceToUIScript(data)
	if err := snowClient.CreateObject(client.EndpointUIScript, uiScript); err != nil {
		return err
	}

	resourceFromUIScript(data, uiScript)

	return readResourceUIScript(data, serviceNowClient)
}

func updateResourceUIScript(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointUIScript, resourceToUIScript(data)); err != nil {
		return err
	}

	return readResourceUIScript(data, serviceNowClient)
}

func deleteResourceUIScript(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointUIScript, data.Id())
}

func resourceFromUIScript(data *schema.ResourceData, script *client.UIScript) {
	var typeString string
	switch script.UIType {
	case "1":
		typeString = "mobile"
	case "0":
		typeString = "desktop"
	default:
		typeString = "all"
	}

	data.SetId(script.ID)
	data.Set(uiScriptName, script.Name)
	data.Set(uiScriptDescription, script.Description)
	data.Set(uiScriptScript, script.Script)
	data.Set(uiScriptActive, script.Active)
	data.Set(uiScriptUIType, typeString)
	data.Set(uiScriptAPIName, script.APIName)
}

func resourceToUIScript(data *schema.ResourceData) *client.UIScript {
	var typeInt string
	switch data.Get(uiScriptUIType).(string) {
	case "mobile":
		typeInt = "1"
	case "desktop":
		typeInt = "0"
	default:
		typeInt = "10"
	}

	uiScript := client.UIScript{
		Name:        data.Get(uiScriptName).(string),
		Description: data.Get(uiScriptDescription).(string),
		Script:      data.Get(uiScriptScript).(string),
		Active:      data.Get(uiScriptActive).(bool),
		UIType:      typeInt,
	}
	uiScript.ID = data.Id()
	uiScript.Scope = data.Get(commonScope).(string)
	return &uiScript
}
