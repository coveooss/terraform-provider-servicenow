package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const applicationMenuTitle = "title"
const applicationMenuDescription = "description"
const applicationMenuHint = "hint"
const applicationMenuDeviceType = "device_type"
const applicationMenuOrder = "order"
const applicationMenuRoles = "roles"
const applicationMenuCategory = "category_id"
const applicationMenuActive = "active"

// ResourceApplicationMenu is a group of modules in the application navigator.
func ResourceApplicationMenu() *schema.Resource {
	return &schema.Resource{
		Create: createResourceApplicationMenu,
		Read:   readResourceApplicationMenu,
		Update: updateResourceApplicationMenu,
		Delete: deleteResourceApplicationMenu,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			applicationMenuTitle: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name appearing in the menu (translated).",
			},
			applicationMenuDescription: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Used to provide a more detailed explanation of what this application does.",
			},
			applicationMenuHint: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Defines the text that appears in a tooltip when a user points to a link to this application.",
			},
			applicationMenuDeviceType: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "browser",
				Description: "Type of device where this menu will appear. Can be 'browser' or 'mobile'.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					warns, errs = validateStringValue(val.(string), key, []string{"browser", "mobile"})
					return
				},
			},
			applicationMenuOrder: {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     100,
				Description: "The display order for the application menu.",
			},
			applicationMenuRoles: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Comma-separated list of Roles (names) that can view this application.",
			},
			applicationMenuCategory: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Specifies the menu category ID, which defines the navigation menu style. The default value is Custom Applications.",
			},
			applicationMenuActive: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether or not this application is in use.",
				Default:     true,
			},
			commonProtectionPolicy: getProtectionPolicySchema(),
			commonScope:            getScopeSchema(),
		},
	}
}

func readResourceApplicationMenu(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	applicationMenu := &client.ApplicationMenu{}
	if err := snowClient.GetObject(client.EndpointApplicationMenu, data.Id(), applicationMenu); err != nil {
		data.SetId("")
		return err
	}

	resourceFromApplicationMenu(data, applicationMenu)

	return nil
}

func createResourceApplicationMenu(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	applicationMenu := resourceToApplicationMenu(data)
	if err := snowClient.CreateObject(client.EndpointApplicationMenu, applicationMenu); err != nil {
		return err
	}

	resourceFromApplicationMenu(data, applicationMenu)

	return readResourceApplicationMenu(data, serviceNowClient)
}

func updateResourceApplicationMenu(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointApplicationMenu, resourceToApplicationMenu(data)); err != nil {
		return err
	}

	return readResourceApplicationMenu(data, serviceNowClient)
}

func deleteResourceApplicationMenu(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointApplicationMenu, data.Id())
}

func resourceFromApplicationMenu(data *schema.ResourceData, applicationMenu *client.ApplicationMenu) {
	data.SetId(applicationMenu.ID)
	data.Set(applicationMenuTitle, applicationMenu.Title)
	data.Set(applicationMenuDescription, applicationMenu.Description)
	data.Set(applicationMenuHint, applicationMenu.Hint)
	data.Set(applicationMenuDeviceType, applicationMenu.DeviceType)
	data.Set(applicationMenuOrder, applicationMenu.Order)
	data.Set(applicationMenuRoles, applicationMenu.Roles)
	data.Set(applicationMenuCategory, applicationMenu.CategoryID)
	data.Set(applicationMenuActive, applicationMenu.Active)
	data.Set(commonProtectionPolicy, applicationMenu.ProtectionPolicy)
	data.Set(commonScope, applicationMenu.Scope)
}

func resourceToApplicationMenu(data *schema.ResourceData) *client.ApplicationMenu {
	applicationMenu := client.ApplicationMenu{
		Title:       data.Get(applicationMenuTitle).(string),
		Description: data.Get(applicationMenuDescription).(string),
		Hint:        data.Get(applicationMenuHint).(string),
		DeviceType:  data.Get(applicationMenuDeviceType).(string),
		Order:       data.Get(applicationMenuOrder).(int),
		Roles:       data.Get(applicationMenuRoles).(string),
		CategoryID:  data.Get(applicationMenuCategory).(string),
		Active:      data.Get(applicationMenuActive).(bool),
	}
	applicationMenu.ID = data.Id()
	applicationMenu.ProtectionPolicy = data.Get(commonProtectionPolicy).(string)
	applicationMenu.Scope = data.Get(commonScope).(string)
	return &applicationMenu
}
