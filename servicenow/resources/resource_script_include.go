package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const scriptIncludeName = "name"
const scriptIncludeClientCallable = "client_callable"
const scriptIncludeDescription = "description"
const scriptIncludeScript = "script"
const scriptIncludeActive = "active"
const scriptIncludeAccess = "access"
const scriptIncludeAPIName = "api_name"

// ResourceScriptInclude manages a Script Include in ServiceNow.
func ResourceScriptInclude() *schema.Resource {
	return &schema.Resource{
		Create: createResourceScriptInclude,
		Read:   readResourceScriptInclude,
		Update: updateResourceScriptInclude,
		Delete: deleteResourceScriptInclude,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			scriptIncludeName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Display name of the script. Needed to have an api_name.",
			},
			scriptIncludeScript: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Javascript script to run when this Script Include is called.",
			},
			scriptIncludeDescription: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Describe what the script does.",
			},
			scriptIncludeClientCallable: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether or not this script can be called from the client-side or only server-side.",
			},
			scriptIncludeActive: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether or not this Script Include is enabled.",
			},
			scriptIncludeAccess: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "package_private",
				Description: "Whether this Script can be accessed from only this application scope or all application scopes. Values can be 'package_private' or 'public'.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					warns, errs = validateStringValue(val.(string), key, []string{"package_private", "public"})
					return
				},
			},
			scriptIncludeAPIName: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Full name of the Script Include needed to call it.",
			},
			commonProtectionPolicy: getProtectionPolicySchema(),
		},
	}
}

func readResourceScriptInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	scriptInclude, err := client.GetScriptInclude(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromScriptInclude(data, scriptInclude)

	return nil
}

func createResourceScriptInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	createdScriptInclude, err := client.CreateScriptInclude(resourceToScriptInclude(data))
	if err != nil {
		return err
	}

	resourceFromScriptInclude(data, createdScriptInclude)

	return readResourceScriptInclude(data, serviceNowClient)
}

func updateResourceScriptInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateScriptInclude(resourceToScriptInclude(data))
	if err != nil {
		return err
	}

	return readResourceScriptInclude(data, serviceNowClient)
}

func deleteResourceScriptInclude(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteScriptInclude(data.Id())
}

func resourceFromScriptInclude(data *schema.ResourceData, scriptInclude *client.ScriptInclude) {
	data.SetId(scriptInclude.Id)
	data.Set(scriptIncludeName, scriptInclude.Name)
	data.Set(scriptIncludeClientCallable, scriptInclude.ClientCallable)
	data.Set(scriptIncludeDescription, scriptInclude.Description)
	data.Set(scriptIncludeScript, scriptInclude.Script)
	data.Set(scriptIncludeActive, scriptInclude.Active)
	data.Set(scriptIncludeAccess, scriptInclude.Access)
	data.Set(scriptIncludeAPIName, scriptInclude.APIName)
	data.Set(commonProtectionPolicy, scriptInclude.ProtectionPolicy)
}

func resourceToScriptInclude(data *schema.ResourceData) *client.ScriptInclude {
	scriptInclude := client.ScriptInclude{
		Name:           data.Get(scriptIncludeName).(string),
		ClientCallable: data.Get(scriptIncludeClientCallable).(bool),
		Description:    data.Get(scriptIncludeDescription).(string),
		Script:         data.Get(scriptIncludeScript).(string),
		Active:         data.Get(scriptIncludeActive).(bool),
		Access:         data.Get(scriptIncludeAccess).(string),
	}
	scriptInclude.Id = data.Id()
	scriptInclude.ProtectionPolicy = data.Get(commonProtectionPolicy).(string)
	return &scriptInclude
}
