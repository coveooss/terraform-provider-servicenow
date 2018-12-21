package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const applicationModuleTitle = "title"
const applicationModuleMenuID = "application_menu_id"
const applicationModuleHint = "hint"
const applicationModuleOrder = "order"
const applicationModuleRoles = "roles"
const applicationModuleActive = "active"
const applicationModuleOverrideRoles = "override_menu_roles"
const applicationModuleLinkType = "link_type"
const applicationModuleLinkArguments = "arguments"
const applicationModuleWindowName = "window_name"
const applicationModuleTableName = "table_name"

// ResourceApplicationModule is a single link in the application navigator.
func ResourceApplicationModule() *schema.Resource {
	return &schema.Resource{
		Create: createResourceApplicationModule,
		Read:   readResourceApplicationModule,
		Update: updateResourceApplicationModule,
		Delete: deleteResourceApplicationModule,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			applicationModuleTitle: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Display name of the module in the application navigator.",
			},
			applicationModuleMenuID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The application Menu ID where this module should reside.",
			},
			applicationModuleHint: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Defines the text that appears in a tooltip when a user points to this module.",
			},
			applicationModuleOrder: {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "The display order for the module in the application menu.",
			},
			applicationModuleRoles: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Comma-separated list of Roles (names) that can view this application module.",
			},
			applicationModuleActive: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether or not this application module is in enabled.",
			},
			applicationModuleOverrideRoles: {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Show this module when the user has the specified roles. Otherwise the user must have the roles specified by both the application menu and the module.",
			},
			applicationModuleLinkType: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of device where this menu will appear. Can be 'DIRECT' for a UI page link or 'LIST' for a table link.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					warns, errs = validateStringValue(val.(string), key, []string{"DIRECT", "LIST"})
					return
				},
			},
			applicationModuleLinkArguments: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Full name of the UI Page where this module will redirect when link type is 'DIRECT'. When the link type is 'LIST', this is optional.",
			},
			applicationModuleWindowName: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The name of the browser window when clicking on a link. For example '_blank' can create a new tab.",
			},
			applicationModuleTableName: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The full name of the table where this module will redirect when the link type is 'LIST'.",
			},
			commonProtectionPolicy: getProtectionPolicySchema(),
		},
	}
}

func readResourceApplicationModule(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	applicationModule, err := client.GetApplicationModule(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromApplicationModule(data, applicationModule)

	return nil
}

func createResourceApplicationModule(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	createdApplicationModule, err := client.CreateApplicationModule(resourceToApplicationModule(data))
	if err != nil {
		return err
	}

	resourceFromApplicationModule(data, createdApplicationModule)

	return readResourceApplicationModule(data, serviceNowClient)
}

func updateResourceApplicationModule(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateApplicationModule(resourceToApplicationModule(data))
	if err != nil {
		return err
	}

	return readResourceApplicationModule(data, serviceNowClient)
}

func deleteResourceApplicationModule(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteApplicationModule(data.Id())
}

func resourceFromApplicationModule(data *schema.ResourceData, applicationModule *client.ApplicationModule) {
	data.SetId(applicationModule.Id)
	data.Set(applicationModuleTitle, applicationModule.Title)
	data.Set(applicationModuleMenuID, applicationModule.MenuID)
	data.Set(applicationModuleHint, applicationModule.Hint)
	data.Set(applicationModuleOrder, applicationModule.Order)
	data.Set(applicationModuleRoles, applicationModule.Roles)
	data.Set(applicationModuleActive, applicationModule.Active)
	data.Set(applicationModuleOverrideRoles, applicationModule.OverrideMenuRoles)
	data.Set(applicationModuleLinkType, applicationModule.LinkType)
	data.Set(applicationModuleLinkArguments, applicationModule.Arguments)
	data.Set(applicationModuleWindowName, applicationModule.WindowName)
	data.Set(applicationModuleTableName, applicationModule.TableName)
	data.Set(commonProtectionPolicy, applicationModule.ProtectionPolicy)
}

func resourceToApplicationModule(data *schema.ResourceData) *client.ApplicationModule {
	applicationModule := client.ApplicationModule{
		Title:             data.Get(applicationModuleTitle).(string),
		MenuID:            data.Get(applicationModuleMenuID).(string),
		Hint:              data.Get(applicationModuleHint).(string),
		Order:             data.Get(applicationModuleOrder).(int),
		Roles:             data.Get(applicationModuleRoles).(string),
		Active:            data.Get(applicationModuleActive).(bool),
		OverrideMenuRoles: data.Get(applicationModuleOverrideRoles).(bool),
		LinkType:          data.Get(applicationModuleLinkType).(string),
		Arguments:         data.Get(applicationModuleLinkArguments).(string),
		WindowName:        data.Get(applicationModuleWindowName).(string),
		TableName:         data.Get(applicationModuleTableName).(string),
	}
	applicationModule.Id = data.Id()
	applicationModule.ProtectionPolicy = data.Get(commonProtectionPolicy).(string)
	return &applicationModule
}
