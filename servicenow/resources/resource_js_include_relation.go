package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const jsIncludeRelationDependencyId = "dependency_id"
const jsIncludeRelationJsIncludeId = "js_include_id"
const jsIncludeRelationOrder = "order"

// ResourceJsIncludeRelation is holding the info about the relation between a js include and a widget dependency.
func ResourceJsIncludeRelation() *schema.Resource {
	return &schema.Resource{
		Create: createResourceJsIncludeRelation,
		Read:   readResourceJsIncludeRelation,
		Update: updateResourceJsIncludeRelation,
		Delete: deleteResourceJsIncludeRelation,

		Schema: map[string]*schema.Schema{
			jsIncludeRelationDependencyId: {
				Type:     schema.TypeString,
				Required: true,
			},
			jsIncludeRelationJsIncludeId: {
				Type:     schema.TypeString,
				Required: true,
			},
			jsIncludeRelationOrder: {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
		},
	}
}

func readResourceJsIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	jsIncludeRelation, err := client.GetJsIncludeRelation(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromJsIncludeRelation(data, jsIncludeRelation)

	return nil
}

func createResourceJsIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	createdPage, err := client.CreateJsIncludeRelation(resourceToJsIncludeRelation(data))
	if err != nil {
		return err
	}

	resourceFromJsIncludeRelation(data, createdPage)

	return readResourceJsIncludeRelation(data, serviceNowClient)
}

func updateResourceJsIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateJsIncludeRelation(resourceToJsIncludeRelation(data))
	if err != nil {
		return err
	}

	return readResourceJsIncludeRelation(data, serviceNowClient)
}

func deleteResourceJsIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteJsIncludeRelation(data.Id())
}

func resourceFromJsIncludeRelation(data *schema.ResourceData, jsIncludeRelation *client.JsIncludeRelation) {
	data.SetId(jsIncludeRelation.Id)
	data.Set(jsIncludeRelationDependencyId, jsIncludeRelation.DependencyId)
	data.Set(jsIncludeRelationJsIncludeId, jsIncludeRelation.JsIncludeId)
	data.Set(jsIncludeRelationOrder, jsIncludeRelation.Order)
}

func resourceToJsIncludeRelation(data *schema.ResourceData) *client.JsIncludeRelation {
	jsIncludeRelation := client.JsIncludeRelation{
		DependencyId: data.Get(jsIncludeRelationDependencyId).(string),
		JsIncludeId:  data.Get(jsIncludeRelationJsIncludeId).(string),
		Order:        data.Get(jsIncludeRelationOrder).(int),
	}
	jsIncludeRelation.Id = data.Id()
	return &jsIncludeRelation
}
