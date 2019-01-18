package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const uiMacroName = "name"
const uiMacroDescription = "description"
const uiMacroXML = "xml"
const uiMacroAPIName = "api_name"
const uiMacroActive = "active"

// ResourceUIMacro manages a UI Macro in ServiceNow.
func ResourceUIMacro() *schema.Resource {
	return &schema.Resource{
		Create: createResourceUIMacro,
		Read:   readResourceUIMacro,
		Update: updateResourceUIMacro,
		Delete: deleteResourceUIMacro,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			uiMacroName: {
				Type:     schema.TypeString,
				Required: true,
			},
			uiMacroXML: {
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

func readResourceUIMacro(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	uiMacro := &client.UIMacro{}
	if err := snowClient.GetObject(client.EndpointUIMacro, data.Id(), uiMacro); err != nil {
		data.SetId("")
		return err
	}

	resourceFromUIMacro(data, uiMacro)

	return nil
}

func createResourceUIMacro(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	uiMacro := resourceToUIMacro(data)
	if err := snowClient.CreateObject(client.EndpointUIMacro, uiMacro); err != nil {
		return err
	}

	resourceFromUIMacro(data, uiMacro)

	return readResourceUIMacro(data, serviceNowClient)
}

func updateResourceUIMacro(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointUIMacro, resourceToUIMacro(data)); err != nil {
		return err
	}

	return readResourceUIMacro(data, serviceNowClient)
}

func deleteResourceUIMacro(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointUIMacro, data.Id())
}

func resourceFromUIMacro(data *schema.ResourceData, uiMacro *client.UIMacro) {
	data.SetId(uiMacro.ID)
	data.Set(uiMacroName, uiMacro.Name)
	data.Set(uiMacroDescription, uiMacro.Description)
	data.Set(uiMacroXML, uiMacro.XML)
	data.Set(uiMacroAPIName, uiMacro.APIName)
	data.Set(uiMacroActive, uiMacro.Active)
	data.Set(commonProtectionPolicy, uiMacro.ProtectionPolicy)
	data.Set(commonScope, uiMacro.Scope)
}

func resourceToUIMacro(data *schema.ResourceData) *client.UIMacro {
	uiMacro := client.UIMacro{
		Name:        data.Get(uiMacroName).(string),
		Description: data.Get(uiMacroDescription).(string),
		XML:         data.Get(uiMacroXML).(string),
		APIName:     data.Get(uiMacroAPIName).(string),
		Active:      data.Get(uiMacroActive).(bool),
	}
	uiMacro.ID = data.Id()
	uiMacro.ProtectionPolicy = data.Get(commonProtectionPolicy).(string)
	uiMacro.Scope = data.Get(commonScope).(string)
	return &uiMacro
}
