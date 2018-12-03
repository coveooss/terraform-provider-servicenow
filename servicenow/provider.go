package servicenow

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/coveo/terraform-provider-servicenow/servicenow/resources"
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider is a Terraform Provider to that manages objects in a ServiceNow instance.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"instance_url": {
				Type:        schema.TypeString,
				Description: "The Url of the ServiceNow instance to work with.",
				Required:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "Username used to manage resources in the ServiceNow instance using Basic authentication.",
				Required:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Password of the user to manage resources.",
				Required:    true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"servicenow_ui_page": resources.ResourceUiPage(),
		},
		ConfigureFunc: configure,
	}
}

func configure(data *schema.ResourceData) (interface{}, error) {
	// Create a new client to talk to the instance.
	client := client.NewClient(
		data.Get("instance_url").(string),
		data.Get("username").(string),
		data.Get("password").(string))

	return client, nil
}
