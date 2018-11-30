package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const name = "name"
const description = "description"
const category = "category"
const direct = "direct"
const clientScript = "client_script"
const processingScript = "processing_script"
const html = "html"

// Resource to manage a UI Page in ServiceNow.
func ResourceUiPage() *schema.Resource {
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema{
			name: {
				Type:     schema.TypeString,
				Required: true,
			},
			description: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			category: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "general",
			},
			direct: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			clientScript: {
				Type:     schema.TypeString,
				Required: true,
			},
			processingScript: {
				Type:     schema.TypeString,
				Required: true,
			},
			html: {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceServerRead(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	uiPage, err := client.GetUiPage(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	updateResourceData(data, uiPage)
	return nil
}

func resourceServerCreate(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	createdPage, err := client.CreateUiPage(resourceToUiPage(data))
	if err != nil {
		return err
	}

	updateResourceData(data, createdPage)

	return resourceServerRead(data, serviceNowClient)
}

func resourceServerUpdate(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateUiPage(resourceToUiPage(data))
	if err != nil {
		return err
	}

	return resourceServerRead(data, serviceNowClient)
}

func resourceServerDelete(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteUiPage(data.Id())
}

func updateResourceData(data *schema.ResourceData, page *client.UiPage) {
	data.SetId(page.Id)
	data.Set(name, page.Name)
	data.Set(description, page.Description)
	data.Set(direct, page.Direct)
	data.Set(html, page.Html)
	data.Set(processingScript, page.ProcessingScript)
	data.Set(clientScript, page.ClientScript)
	data.Set(category, page.Category)
}

func resourceToUiPage(data *schema.ResourceData) *client.UiPage {
	return &client.UiPage{
		Id:               data.Id(),
		Name:             data.Get(name).(string),
		Description:      data.Get(description).(string),
		Direct:           data.Get(direct).(bool),
		Html:             data.Get(html).(string),
		ProcessingScript: data.Get(processingScript).(string),
		ClientScript:     data.Get(clientScript).(string),
		Category:         data.Get(category).(string),
	}
}
