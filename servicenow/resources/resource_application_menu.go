package resources

import (
	"fmt"

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
					v := val.(string)
					if v != "browser" && v != "mobile" {
						errs = append(errs, fmt.Errorf("%q must be 'browser' or 'mobile', got: %s", key, v))
					}
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
				Description: "Comma-separated list of Roles that can view this application.",
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
		},
	}
}

func readResourceApplicationMenu(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	applicationMenu, err := client.GetApplicationMenu(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromApplicationMenu(data, applicationMenu)

	return nil
}

func createResourceApplicationMenu(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	createdApplicationMenu, err := client.CreateApplicationMenu(resourceToApplicationMenu(data))
	if err != nil {
		return err
	}

	resourceFromApplicationMenu(data, createdApplicationMenu)

	return readResourceApplicationMenu(data, serviceNowClient)
}

func updateResourceApplicationMenu(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateApplicationMenu(resourceToApplicationMenu(data))
	if err != nil {
		return err
	}

	return readResourceApplicationMenu(data, serviceNowClient)
}

func deleteResourceApplicationMenu(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteApplicationMenu(data.Id())
}

func resourceFromApplicationMenu(data *schema.ResourceData, applicationMenu *client.ApplicationMenu) {
	data.SetId(applicationMenu.Id)
	data.Set(applicationMenuTitle, applicationMenu.Title)
	data.Set(applicationMenuDescription, applicationMenu.Description)
	data.Set(applicationMenuHint, applicationMenu.Hint)
	data.Set(applicationMenuDeviceType, applicationMenu.DeviceType)
	data.Set(applicationMenuOrder, applicationMenu.Order)
	data.Set(applicationMenuRoles, applicationMenu.Roles)
	data.Set(applicationMenuCategory, applicationMenu.CategoryID)
	data.Set(applicationMenuActive, applicationMenu.Active)
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
	applicationMenu.Id = data.Id()
	return &applicationMenu
}
