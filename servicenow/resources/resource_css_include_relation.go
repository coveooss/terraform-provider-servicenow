package resources

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const cssIncludeRelationDependencyID = "dependency_id"
const cssIncludeRelationCSSIncludeID = "css_include_id"
const cssIncludeRelationOrder = "order"

// ResourceCSSIncludeRelation is holding the info about the relation between a CSS Include and a widget dependency.
func ResourceCSSIncludeRelation() *schema.Resource {
	return &schema.Resource{
		Create: createResourceCSSIncludeRelation,
		Read:   readResourceCSSIncludeRelation,
		Update: updateResourceCSSIncludeRelation,
		Delete: deleteResourceCSSIncludeRelation,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			cssIncludeRelationDependencyID: {
				Type:     schema.TypeString,
				Required: true,
			},
			cssIncludeRelationCSSIncludeID: {
				Type:     schema.TypeString,
				Required: true,
			},
			cssIncludeRelationOrder: {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceCSSIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	cssIncludeRelation := &client.CSSIncludeRelation{}
	if err := snowClient.GetObject(client.EndpointCSSIncludeRelation, data.Id(), cssIncludeRelation); err != nil {
		data.SetId("")
		return err
	}

	resourceFromCSSIncludeRelation(data, cssIncludeRelation)

	return nil
}

func createResourceCSSIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	cssIncludeRelation := resourceToCSSIncludeRelation(data)
	if err := snowClient.CreateObject(client.EndpointCSSIncludeRelation, cssIncludeRelation); err != nil {
		return err
	}

	resourceFromCSSIncludeRelation(data, cssIncludeRelation)

	return readResourceCSSIncludeRelation(data, serviceNowClient)
}

func updateResourceCSSIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointCSSIncludeRelation, resourceToCSSIncludeRelation(data)); err != nil {
		return err
	}

	return readResourceCSSIncludeRelation(data, serviceNowClient)
}

func deleteResourceCSSIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointCSSIncludeRelation, data.Id())
}

func resourceFromCSSIncludeRelation(data *schema.ResourceData, cssIncludeRelation *client.CSSIncludeRelation) {
	data.SetId(cssIncludeRelation.ID)
	data.Set(cssIncludeRelationDependencyID, cssIncludeRelation.DependencyID)
	data.Set(cssIncludeRelationCSSIncludeID, cssIncludeRelation.CSSIncludeID)
	data.Set(cssIncludeRelationOrder, cssIncludeRelation.Order)
	data.Set(commonScope, cssIncludeRelation.Scope)
}

func resourceToCSSIncludeRelation(data *schema.ResourceData) *client.CSSIncludeRelation {
	cssIncludeRelation := client.CSSIncludeRelation{
		DependencyID: data.Get(cssIncludeRelationDependencyID).(string),
		CSSIncludeID: data.Get(cssIncludeRelationCSSIncludeID).(string),
		Order:        data.Get(cssIncludeRelationOrder).(int),
	}
	cssIncludeRelation.ID = data.Id()
	cssIncludeRelation.Scope = data.Get(commonScope).(string)
	return &cssIncludeRelation
}
