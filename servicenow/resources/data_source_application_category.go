package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const applicationCategoryName = "name"
const applicationCategoryOrder = "order"
const applicationCategoryStyle = "style"

// ResourceApplicationCategory is holding the info about an application category.
func DataSourceApplicationCategory() *schema.Resource {
	resourceSchema := map[string]*schema.Schema{
		applicationCategoryName: {
			Type:     schema.TypeString,
			Required: true,
		},
		applicationCategoryOrder: {
			Type:     schema.TypeInt,
			Computed: true,
		},
		applicationCategoryStyle: {
			Type:     schema.TypeString,
			Computed: true,
		},
		commonScope: getScopeSchema(),
	}
	setOnlyRequiredSchema(resourceSchema, applicationCategoryName)

	return &schema.Resource{
		Read:   readResourceApplicationCategory,
		Schema: resourceSchema,
	}
}

func readResourceApplicationCategory(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	applicationCategory, err := client.GetApplicationCategoryByName(data.Get(applicationCategoryName).(string))
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromApplicationCategory(data, applicationCategory)

	return nil
}

func resourceFromApplicationCategory(data *schema.ResourceData, applicationCategory *client.ApplicationCategory) {
	data.SetId(applicationCategory.Id)
	data.Set(applicationCategoryName, applicationCategory.Name)
	data.Set(applicationCategoryOrder, applicationCategory.Order)
	data.Set(applicationCategoryStyle, applicationCategory.Style)
	data.Set(commonScope, applicationCategory.Scope)
}
