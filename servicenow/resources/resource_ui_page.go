package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const uiPageName = "name"
const uiPageDescription = "description"
const uiPageCategory = "category"
const uiPageDirect = "direct"
const uiPageClientScript = "client_script"
const uiPageProcessingScript = "processing_script"
const uiPageHtml = "html"
const uiPageEndpoint = "endpoint"

// ResourceUiPage manages a UI Page in ServiceNow.
func ResourceUiPage() *schema.Resource {
	return &schema.Resource{
		Create: createResourceUiPage,
		Read:   readResourceUiPage,
		Update: updateResourceUiPage,
		Delete: deleteResourceUiPage,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			uiPageName: {
				Type:     schema.TypeString,
				Required: true,
			},
			uiPageDescription: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			uiPageCategory: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "general",
			},
			uiPageDirect: {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			uiPageClientScript: {
				Type:     schema.TypeString,
				Required: true,
			},
			uiPageProcessingScript: {
				Type:     schema.TypeString,
				Required: true,
			},
			uiPageHtml: {
				Type:     schema.TypeString,
				Required: true,
			},
			uiPageEndpoint: {
				Type:     schema.TypeString,
				Computed: true,
			},
			commonProtectionPolicy: getProtectionPolicySchema(),
			commonScope:            getScopeSchema(),
		},
	}
}

func readResourceUiPage(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	uiPage, err := client.GetUiPage(data.Id())
	if err != nil {
		data.SetId("")
		return err
	}

	resourceFromUiPage(data, uiPage)

	return nil
}

func createResourceUiPage(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	createdPage, err := client.CreateUiPage(resourceToUiPage(data))
	if err != nil {
		return err
	}

	resourceFromUiPage(data, createdPage)

	return readResourceUiPage(data, serviceNowClient)
}

func updateResourceUiPage(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	err := client.UpdateUiPage(resourceToUiPage(data))
	if err != nil {
		return err
	}

	return readResourceUiPage(data, serviceNowClient)
}

func deleteResourceUiPage(data *schema.ResourceData, serviceNowClient interface{}) error {
	client := serviceNowClient.(*client.ServiceNowClient)
	return client.DeleteUiPage(data.Id())
}

func resourceFromUiPage(data *schema.ResourceData, page *client.UiPage) {
	data.SetId(page.Id)
	data.Set(uiPageName, page.Name)
	data.Set(uiPageDescription, page.Description)
	data.Set(uiPageDirect, page.Direct)
	data.Set(uiPageHtml, page.Html)
	data.Set(uiPageProcessingScript, page.ProcessingScript)
	data.Set(uiPageClientScript, page.ClientScript)
	data.Set(uiPageCategory, page.Category)
	data.Set(uiPageEndpoint, page.Endpoint)
	data.Set(commonProtectionPolicy, page.ProtectionPolicy)
	data.Set(commonScope, page.Scope)
}

func resourceToUiPage(data *schema.ResourceData) *client.UiPage {
	uiPage := client.UiPage{
		Name:             data.Get(uiPageName).(string),
		Description:      data.Get(uiPageDescription).(string),
		Direct:           data.Get(uiPageDirect).(bool),
		Html:             data.Get(uiPageHtml).(string),
		ProcessingScript: data.Get(uiPageProcessingScript).(string),
		ClientScript:     data.Get(uiPageClientScript).(string),
		Category:         data.Get(uiPageCategory).(string),
	}
	uiPage.Id = data.Id()
	uiPage.ProtectionPolicy = data.Get(commonProtectionPolicy).(string)
	uiPage.Scope = data.Get(commonScope).(string)
	return &uiPage
}
