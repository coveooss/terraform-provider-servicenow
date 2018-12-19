package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const cssIncludeRelationDependencyId = "dependency_id"
const cssIncludeRelationCssIncludeId = "css_include_id"
const cssIncludeRelationOrder = "order"

// ResourceCssIncludeRelation is holding the info about the relation between a CSS Include and a widget dependency.
func ResourceCssIncludeRelation() *schema.Resource {
	return &schema.Resource{
		Create: createResourceCssIncludeRelation,
		Read:   readResourceCssIncludeRelation,
		Update: updateResourceCssIncludeRelation,
		Delete: deleteResourceCssIncludeRelation,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			cssIncludeRelationDependencyId: {
				Type:     schema.TypeString,
				Required: true,
			},
			cssIncludeRelationCssIncludeId: {
				Type:     schema.TypeString,
				Required: true,
			},
			cssIncludeRelationOrder: {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
		},
	}
}

func readResourceCssIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	cssIncludeRelation, err := client.GetCssIncludeRelation(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromCssIncludeRelation(data, cssIncludeRelation)

	return nil
}

func createResourceCssIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	createdPage, err := client.CreateCssIncludeRelation(resourceToCssIncludeRelation(data))
	if err != nil {
		return err
	}

	resourceFromCssIncludeRelation(data, createdPage)

	return readResourceCssIncludeRelation(data, serviceNowClient)
}

func updateResourceCssIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateCssIncludeRelation(resourceToCssIncludeRelation(data))
	if err != nil {
		return err
	}

	return readResourceCssIncludeRelation(data, serviceNowClient)
}

func deleteResourceCssIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteCssIncludeRelation(data.Id())
}

func resourceFromCssIncludeRelation(data *schema.ResourceData, cssIncludeRelation *client.CssIncludeRelation) {
	data.SetId(cssIncludeRelation.Id)
	data.Set(cssIncludeRelationDependencyId, cssIncludeRelation.DependencyId)
	data.Set(cssIncludeRelationCssIncludeId, cssIncludeRelation.CssIncludeId)
	data.Set(cssIncludeRelationOrder, cssIncludeRelation.Order)
}

func resourceToCssIncludeRelation(data *schema.ResourceData) *client.CssIncludeRelation {
	cssIncludeRelation := client.CssIncludeRelation{
		DependencyId: data.Get(cssIncludeRelationDependencyId).(string),
		CssIncludeId: data.Get(cssIncludeRelationCssIncludeId).(string),
		Order:        data.Get(cssIncludeRelationOrder).(int),
	}
	cssIncludeRelation.Id = data.Id()
	return &cssIncludeRelation
}
