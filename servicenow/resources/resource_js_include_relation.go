package resources

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const jsIncludeRelationDependencyID = "dependency_id"
const jsIncludeRelationJsIncludeID = "js_include_id"
const jsIncludeRelationOrder = "order"

// ResourceJsIncludeRelation is holding the info about the relation between a js include and a widget dependency.
func ResourceJsIncludeRelation() *schema.Resource {
	return &schema.Resource{
		Create: createResourceJsIncludeRelation,
		Read:   readResourceJsIncludeRelation,
		Update: updateResourceJsIncludeRelation,
		Delete: deleteResourceJsIncludeRelation,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			jsIncludeRelationDependencyID: {
				Type:     schema.TypeString,
				Required: true,
			},
			jsIncludeRelationJsIncludeID: {
				Type:     schema.TypeString,
				Required: true,
			},
			jsIncludeRelationOrder: {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceJsIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	jsIncludeRelation := &client.JsIncludeRelation{}
	if err := snowClient.GetObject(client.EndpointJsIncludeRelation, data.Id(), jsIncludeRelation); err != nil {
		data.SetId("")
		return err
	}

	resourceFromJsIncludeRelation(data, jsIncludeRelation)

	return nil
}

func createResourceJsIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	jsIncludeRelation := resourceToJsIncludeRelation(data)
	if err := snowClient.CreateObject(client.EndpointJsIncludeRelation, jsIncludeRelation); err != nil {
		return err
	}

	resourceFromJsIncludeRelation(data, jsIncludeRelation)

	return readResourceJsIncludeRelation(data, serviceNowClient)
}

func updateResourceJsIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointJsIncludeRelation, resourceToJsIncludeRelation(data)); err != nil {
		return err
	}

	return readResourceJsIncludeRelation(data, serviceNowClient)
}

func deleteResourceJsIncludeRelation(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointJsIncludeRelation, data.Id())
}

func resourceFromJsIncludeRelation(data *schema.ResourceData, jsIncludeRelation *client.JsIncludeRelation) {
	data.SetId(jsIncludeRelation.ID)
	data.Set(jsIncludeRelationDependencyID, jsIncludeRelation.DependencyID)
	data.Set(jsIncludeRelationJsIncludeID, jsIncludeRelation.JsIncludeID)
	data.Set(jsIncludeRelationOrder, jsIncludeRelation.Order)
	data.Set(commonScope, jsIncludeRelation.Scope)
}

func resourceToJsIncludeRelation(data *schema.ResourceData) *client.JsIncludeRelation {
	jsIncludeRelation := client.JsIncludeRelation{
		DependencyID: data.Get(jsIncludeRelationDependencyID).(string),
		JsIncludeID:  data.Get(jsIncludeRelationJsIncludeID).(string),
		Order:        data.Get(jsIncludeRelationOrder).(int),
	}
	jsIncludeRelation.ID = data.Id()
	jsIncludeRelation.Scope = data.Get(commonScope).(string)
	return &jsIncludeRelation
}
