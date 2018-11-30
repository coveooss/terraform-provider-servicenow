package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
)

const NAME = "name"
const DESCRIPTION = "description"
const CATEGORY = "category"
const DIRECT = "direct"
const CLIENT_SCRIPT = "client_script"
const PROCESSING_SCRIPT = "processing_script"
const HTML = "html"

func ResourceUiPage() *schema.Resource {
	return &schema.Resource {
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,

		Schema: map[string]*schema.Schema {
			NAME: &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			DESCRIPTION: &schema.Schema {
				Type:     schema.TypeString,
				Optional: true,
				Default: "",
			},
			CATEGORY: &schema.Schema {
				Type:     schema.TypeString,
				Optional: true,
				Default: "general",
			},
			DIRECT: &schema.Schema {
				Type:     schema.TypeBool,
				Optional: true,
				Default: false,
			},
			CLIENT_SCRIPT: &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			PROCESSING_SCRIPT: &schema.Schema {
				Type:     schema.TypeString,
				Required: true,
			},
			HTML: &schema.Schema {
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
	data.Set(NAME, page.Name)
	data.Set(DESCRIPTION, page.Description)
	data.Set(DIRECT, page.Direct)
	data.Set(HTML, page.Html)
	data.Set(PROCESSING_SCRIPT, page.ProcessingScript)
	data.Set(CLIENT_SCRIPT, page.ClientScript)
	data.Set(CATEGORY, page.Category)
}

func resourceToUiPage(data *schema.ResourceData) *client.UiPage {
	return &client.UiPage {
		Id: data.Id(),
		Name: data.Get(NAME).(string),
		Description: data.Get(DESCRIPTION).(string),
		Direct: data.Get(DIRECT).(bool),
		Html: data.Get(HTML).(string),
		ProcessingScript: data.Get(PROCESSING_SCRIPT).(string),
		ClientScript: data.Get(CLIENT_SCRIPT).(string),
		Category: data.Get(CATEGORY).(string),
	}
}