package resources

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const applicationCategoryName = "name"
const applicationCategoryOrder = "order"
const applicationCategoryStyle = "style"

// DataSourceApplicationCategory is holding the info about an application category.
func DataSourceApplicationCategory() *schema.Resource {
	resourceSchema := map[string]*schema.Schema{
		applicationCategoryName: {
			Type:     schema.TypeString,
			Required: true,
		},
		applicationCategoryOrder: {
			Type:     schema.TypeString,
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
	snowClient := serviceNowClient.(client.ServiceNowClient)
	applicationCategory := &client.ApplicationCategory{}
	if err := snowClient.GetObjectByName(client.EndpointApplicationCategory, data.Get(applicationCategoryName).(string), applicationCategory); err != nil {
		data.SetId("")
		return err
	}

	resourceFromApplicationCategory(data, applicationCategory)

	return nil
}

func resourceFromApplicationCategory(data *schema.ResourceData, applicationCategory *client.ApplicationCategory) {
	data.SetId(applicationCategory.ID)
	data.Set(applicationCategoryName, applicationCategory.Name)
	data.Set(applicationCategoryOrder, applicationCategory.Order)
	data.Set(applicationCategoryStyle, applicationCategory.Style)
	data.Set(commonScope, applicationCategory.Scope)
}
