package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const uiMacroName = "name"
const uiMacroDescription = "description"
const uiMacroXml = "xml"
const uiMacroAPIName = "api_name"
const uiMacroActive = "active"

// ResourceUiMacro manages a UI Macro in ServiceNow.
func ResourceUiMacro() *schema.Resource {
	return &schema.Resource{
		Create: createResourceUiMacro,
		Read:   readResourceUiMacro,
		Update: updateResourceUiMacro,
		Delete: deleteResourceUiMacro,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			uiMacroName: {
				Type:     schema.TypeString,
				Required: true,
			},
			uiMacroXml: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The body of the UI Macro. Must be in XML format.",
			},
			uiMacroDescription: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			uiMacroAPIName: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				ForceNew:    true,
				Description: "The scoped name of the macro. Normally contains the name field prefixed with the application scope.",
			},
			uiMacroActive: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether or not this Macro is enabled.",
			},
			commonProtectionPolicy: getProtectionPolicySchema(),
			commonScope:            getScopeSchema(),
		},
	}
}

func readResourceUiMacro(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	uiMacro, err := client.GetUiMacro(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromUiMacro(data, uiMacro)

	return nil
}

func createResourceUiMacro(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	createdUiMacro, err := client.CreateUiMacro(resourceToUiMacro(data))
	if err != nil {
		return err
	}

	resourceFromUiMacro(data, createdUiMacro)

	return readResourceUiMacro(data, serviceNowClient)
}

func updateResourceUiMacro(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateUiMacro(resourceToUiMacro(data))
	if err != nil {
		return err
	}

	return readResourceUiMacro(data, serviceNowClient)
}

func deleteResourceUiMacro(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteUiMacro(data.Id())
}

func resourceFromUiMacro(data *schema.ResourceData, uiMacro *client.UiMacro) {
	data.SetId(uiMacro.Id)
	data.Set(uiMacroName, uiMacro.Name)
	data.Set(uiMacroDescription, uiMacro.Description)
	data.Set(uiMacroXml, uiMacro.Xml)
	data.Set(uiMacroAPIName, uiMacro.APIName)
	data.Set(uiMacroActive, uiMacro.Active)
	data.Set(commonProtectionPolicy, uiMacro.ProtectionPolicy)
	data.Set(commonScope, uiMacro.Scope)
}

func resourceToUiMacro(data *schema.ResourceData) *client.UiMacro {
	uiMacro := client.UiMacro{
		Name:        data.Get(uiMacroName).(string),
		Description: data.Get(uiMacroDescription).(string),
		Xml:         data.Get(uiMacroXml).(string),
		APIName:     data.Get(uiMacroAPIName).(string),
		Active:      data.Get(uiMacroActive).(bool),
	}
	uiMacro.Id = data.Id()
	uiMacro.ProtectionPolicy = data.Get(commonProtectionPolicy).(string)
	uiMacro.Scope = data.Get(commonScope).(string)
	return &uiMacro
}
