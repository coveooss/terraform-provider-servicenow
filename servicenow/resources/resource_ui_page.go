package resources

import (
	"github.com/coveooss/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const uiPageName = "name"
const uiPageDescription = "description"
const uiPageCategory = "category"
const uiPageDirect = "direct"
const uiPageClientScript = "client_script"
const uiPageProcessingScript = "processing_script"
const uiPageHTML = "html"
const uiPageEndpoint = "endpoint"

// ResourceUIPage manages a UI Page in ServiceNow.
func ResourceUIPage() *schema.Resource {
	return &schema.Resource{
		Create: createResourceUIPage,
		Read:   readResourceUIPage,
		Update: updateResourceUIPage,
		Delete: deleteResourceUIPage,

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
			uiPageHTML: {
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

func readResourceUIPage(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	uiPage := &client.UIPage{}
	if err := snowClient.GetObject(client.EndpointUIPage, data.Id(), uiPage); err != nil {
		data.SetId("")
		return err
	}

	resourceFromUIPage(data, uiPage)

	return nil
}

func createResourceUIPage(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	uiPage := resourceToUIPage(data)
	if err := snowClient.CreateObject(client.EndpointUIPage, uiPage); err != nil {
		return err
	}

	resourceFromUIPage(data, uiPage)

	return readResourceUIPage(data, serviceNowClient)
}

func updateResourceUIPage(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointUIPage, resourceToUIPage(data)); err != nil {
		return err
	}

	return readResourceUIPage(data, serviceNowClient)
}

func deleteResourceUIPage(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointUIPage, data.Id())
}

func resourceFromUIPage(data *schema.ResourceData, page *client.UIPage) {
	data.SetId(page.ID)
	data.Set(uiPageName, page.Name)
	data.Set(uiPageDescription, page.Description)
	data.Set(uiPageDirect, page.Direct)
	data.Set(uiPageHTML, page.HTML)
	data.Set(uiPageProcessingScript, page.ProcessingScript)
	data.Set(uiPageClientScript, page.ClientScript)
	data.Set(uiPageCategory, page.Category)
	data.Set(uiPageEndpoint, page.Endpoint)
	data.Set(commonProtectionPolicy, page.ProtectionPolicy)
	data.Set(commonScope, page.Scope)
}

func resourceToUIPage(data *schema.ResourceData) *client.UIPage {
	uiPage := client.UIPage{
		Name:             data.Get(uiPageName).(string),
		Description:      data.Get(uiPageDescription).(string),
		Direct:           data.Get(uiPageDirect).(bool),
		HTML:             data.Get(uiPageHTML).(string),
		ProcessingScript: data.Get(uiPageProcessingScript).(string),
		ClientScript:     data.Get(uiPageClientScript).(string),
		Category:         data.Get(uiPageCategory).(string),
	}
	uiPage.ID = data.Id()
	uiPage.ProtectionPolicy = data.Get(commonProtectionPolicy).(string)
	uiPage.Scope = data.Get(commonScope).(string)
	return &uiPage
}
