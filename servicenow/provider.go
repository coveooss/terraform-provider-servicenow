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
				Sensitive:   true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"servicenow_application":                resources.ResourceApplication(),
			"servicenow_application_menu":           resources.ResourceApplicationMenu(),
			"servicenow_application_module":         resources.ResourceApplicationModule(),
			"servicenow_css_include":                resources.ResourceCssInclude(),
			"servicenow_css_include_relation":       resources.ResourceCssIncludeRelation(),
			"servicenow_js_include":                 resources.ResourceJsInclude(),
			"servicenow_js_include_relation":        resources.ResourceJsIncludeRelation(),
			"servicenow_role":                       resources.ResourceRole(),
			"servicenow_script_include":             resources.ResourceScriptInclude(),
			"servicenow_ui_macro":                   resources.ResourceUiMacro(),
			"servicenow_ui_page":                    resources.ResourceUiPage(),
			"servicenow_widget":                     resources.ResourceWidget(),
			"servicenow_widget_dependency":          resources.ResourceWidgetDependency(),
			"servicenow_widget_dependency_relation": resources.ResourceWidgetDependencyRelation(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"servicenow_application": resources.DataSourceApplication(),
			"servicenow_role":        resources.DataSourceRole(),
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
