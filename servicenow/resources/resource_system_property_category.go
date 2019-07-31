package resources

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const systemPropertyCategoryName = "name"
const systemPropertyCategoryTitleHTML = "title_html"

// ResourceSystemPropertyCategory manages a System Property Category in ServiceNow.
func ResourceSystemPropertyCategory() *schema.Resource {
	return &schema.Resource{
		Create: createResourceSystemPropertyCategory,
		Read:   readResourceSystemPropertyCategory,
		Update: updateResourceSystemPropertyCategory,
		Delete: deleteResourceSystemPropertyCategory,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			systemPropertyCategoryName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Display name of the category.",
			},
			systemPropertyCategoryTitleHTML: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The HTML displayed at the top of the page when configuring properties for this category.",
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceSystemPropertyCategory(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	systemPropertyCategory := &client.SystemPropertyCategory{}
	if err := snowClient.GetObject(client.EndpointSystemPropertyCategory, data.Id(), systemPropertyCategory); err != nil {
		data.SetId("")
		return err
	}

	resourceFromSystemPropertyCategory(data, systemPropertyCategory)

	return nil
}

func createResourceSystemPropertyCategory(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	systemPropertyCategory := resourceToSystemPropertyCategory(data)
	if err := snowClient.CreateObject(client.EndpointSystemPropertyCategory, systemPropertyCategory); err != nil {
		return err
	}

	resourceFromSystemPropertyCategory(data, systemPropertyCategory)

	return readResourceSystemPropertyCategory(data, serviceNowClient)
}

func updateResourceSystemPropertyCategory(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointSystemPropertyCategory, resourceToSystemPropertyCategory(data)); err != nil {
		return err
	}

	return readResourceSystemPropertyCategory(data, serviceNowClient)
}

func deleteResourceSystemPropertyCategory(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointSystemPropertyCategory, data.Id())
}

func resourceFromSystemPropertyCategory(data *schema.ResourceData, systemPropertyCategory *client.SystemPropertyCategory) {
	data.SetId(systemPropertyCategory.ID)
	data.Set(systemPropertyCategoryName, systemPropertyCategory.Name)
	data.Set(systemPropertyCategoryTitleHTML, systemPropertyCategory.TitleHTML)
	data.Set(commonScope, systemPropertyCategory.Scope)
}

func resourceToSystemPropertyCategory(data *schema.ResourceData) *client.SystemPropertyCategory {
	systemPropertyCategory := client.SystemPropertyCategory{
		Name:      data.Get(systemPropertyCategoryName).(string),
		TitleHTML: data.Get(systemPropertyCategoryTitleHTML).(string),
	}
	systemPropertyCategory.ID = data.Id()
	systemPropertyCategory.Scope = data.Get(commonScope).(string)
	return &systemPropertyCategory
}
