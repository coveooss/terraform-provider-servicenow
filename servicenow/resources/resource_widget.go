package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const widgetId = "identifier"
const widgetName = "name"
const widgetTemplate = "template"
const widgetCss = "css"
const widgetPublic = "public"
const widgetRoles = "roles"
const widgetLink = "link_function"
const widgetDescription = "description"
const widgetClientScript = "client_script"
const widgetServerScript = "server_script"
const widgetDemoData = "demo_data"
const widgetOptionSchema = "option_schema"
const widgetHasPreview = "has_preview"
const widgetDataTable = "data_table"
const widgetControllerAs = "controller_as"

// Resource to manage a Widget in ServiceNow.
func ResourceWidget() *schema.Resource {
	return &schema.Resource{
		Create: createResourceWidget,
		Read:   readResourceWidget,
		Update: updateResourceWidget,
		Delete: deleteResourceWidget,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			widgetId: {
				Type:     schema.TypeString,
				Required: true,
			},
			widgetName: {
				Type:     schema.TypeString,
				Required: true,
			},
			widgetTemplate: {
				Type:     schema.TypeString,
				Required: true,
			},
			widgetCss: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			widgetPublic: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  false,
			},
			widgetRoles: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			widgetLink: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			widgetDescription: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			widgetClientScript: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			widgetServerScript: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			widgetDemoData: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			widgetOptionSchema: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			widgetHasPreview: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  false,
			},
			widgetDataTable: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			widgetControllerAs: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "c",
			},
			commonProtectionPolicy: getProtectionPolicySchema(),
		},
	}
}

func readResourceWidget(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	widget, err := client.GetWidget(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromWidget(data, widget)

	return nil
}

func createResourceWidget(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	createdPage, err := client.CreateWidget(resourceToWidget(data))
	if err != nil {
		return err
	}

	resourceFromWidget(data, createdPage)

	return readResourceWidget(data, serviceNowClient)
}

func updateResourceWidget(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateWidget(resourceToWidget(data))
	if err != nil {
		return err
	}

	return readResourceWidget(data, serviceNowClient)
}

func deleteResourceWidget(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteWidget(data.Id())
}

func resourceFromWidget(data *schema.ResourceData, widget *client.Widget) {
	data.SetId(widget.Id)
	data.Set(widgetId, widget.CustomId)
	data.Set(widgetName, widget.Name)
	data.Set(widgetTemplate, widget.Template)
	data.Set(widgetCss, widget.Css)
	data.Set(widgetPublic, widget.Public)
	data.Set(widgetRoles, widget.Roles)
	data.Set(widgetLink, widget.Link)
	data.Set(widgetDescription, widget.Description)
	data.Set(widgetClientScript, widget.ClientScript)
	data.Set(widgetServerScript, widget.ServerScript)
	data.Set(widgetDemoData, widget.DemoData)
	data.Set(widgetOptionSchema, widget.OptionSchema)
	data.Set(widgetHasPreview, widget.HasPreview)
	data.Set(widgetDataTable, widget.DataTable)
	data.Set(widgetControllerAs, widget.ControllerAs)
	data.Set(commonProtectionPolicy, widget.ProtectionPolicy)
}

func resourceToWidget(data *schema.ResourceData) *client.Widget {
	widget := client.Widget{
		CustomId:     data.Get(widgetId).(string),
		Name:         data.Get(widgetName).(string),
		Template:     data.Get(widgetTemplate).(string),
		Css:          data.Get(widgetCss).(string),
		Public:       data.Get(widgetPublic).(bool),
		Roles:        data.Get(widgetId).(string),
		Link:         data.Get(widgetLink).(string),
		Description:  data.Get(widgetDescription).(string),
		ClientScript: data.Get(widgetClientScript).(string),
		ServerScript: data.Get(widgetServerScript).(string),
		DemoData:     data.Get(widgetDemoData).(string),
		OptionSchema: data.Get(widgetOptionSchema).(string),
		HasPreview:   data.Get(widgetHasPreview).(bool),
		DataTable:    data.Get(widgetDataTable).(string),
		ControllerAs: data.Get(widgetControllerAs).(string),
	}
	widget.Id = data.Id()
	widget.ProtectionPolicy = data.Get(commonProtectionPolicy).(string)
	return &widget
}
