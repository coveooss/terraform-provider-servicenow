package resources

import (
	"github.com/coveo/terraform-provider-servicenow/servicenow/client"
	"github.com/hashicorp/terraform/helper/schema"
)

const oauthEntityName = "name"
const oauthEntityClientUUID = "client_uuid"
const oauthEntityClientID = "client_id"
const oauthEntityAccessTokenLifespan = "access_token_lifespan"
const oauthEntityRefreshTokenLifespan = "refresh_token_lifespan"
const oauthEntityRedirectURL = "redirect_url"
const oauthEntityLogoURL = "logo_url"
const oauthEntityAccess = "access"

// ResourceOAuthEntity manages an OAuthEntity in ServiceNow.
func ResourceOAuthEntity() *schema.Resource {
	return &schema.Resource{
		Create: createResourceOAuthEntity,
		Read:   readResourceOAuthEntity,
		Update: updateResourceOAuthEntity,
		Delete: deleteResourceOAuthEntity,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			oauthEntityName: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the OAuth app.",
			},
			oauthEntityClientUUID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Internal unique identifier of the entity.",
			},
			oauthEntityClientID: {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "OAuth Client ID required during handshake.",
			},
			oauthEntityAccessTokenLifespan: {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1800,
				Description: "Number of seconds a newly created access token is good for.",
			},
			oauthEntityRefreshTokenLifespan: {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     8640000,
				Description: "Number of seconds the refresh token is good for.",
			},
			oauthEntityRedirectURL: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The OAuth app's endpoint to receive the authorization code.",
			},
			oauthEntityLogoURL: {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			oauthEntityAccess: {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "public",
				Description: "Whether this Script can be accessed from only this application scope or all application scopes. Values can be 'package_private' or 'public'.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					warns, errs = validateStringValue(val.(string), key, []string{"package_private", "public"})
					return
				},
			},
			commonScope: getScopeSchema(),
		},
	}
}

func readResourceOAuthEntity(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	oauthEntity := &client.OAuthEntity{}
	if err := snowClient.GetObject(client.EndpointOAuthEntity, data.Id(), oauthEntity); err != nil {
		data.SetId("")
		return err
	}

	resourceFromOAuthEntity(data, oauthEntity)

	return nil
}

func createResourceOAuthEntity(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	oauthEntity := resourceToOAuthEntity(data)
	if err := snowClient.CreateObject(client.EndpointOAuthEntity, oauthEntity); err != nil {
		return err
	}

	resourceFromOAuthEntity(data, oauthEntity)

	return readResourceOAuthEntity(data, serviceNowClient)
}

func updateResourceOAuthEntity(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	if err := snowClient.UpdateObject(client.EndpointOAuthEntity, resourceToOAuthEntity(data)); err != nil {
		return err
	}

	return readResourceOAuthEntity(data, serviceNowClient)
}

func deleteResourceOAuthEntity(data *schema.ResourceData, serviceNowClient interface{}) error {
	snowClient := serviceNowClient.(*client.ServiceNowClient)
	return snowClient.DeleteObject(client.EndpointOAuthEntity, data.Id())
}

func resourceFromOAuthEntity(data *schema.ResourceData, oauthEntity *client.OAuthEntity) {
	data.SetId(oauthEntity.ID)
	data.Set(oauthEntityName, oauthEntity.Name)
	data.Set(oauthEntityClientUUID, oauthEntity.ClientUUID)
	data.Set(oauthEntityClientID, oauthEntity.ClientID)
	data.Set(oauthEntityAccessTokenLifespan, oauthEntity.AccessTokenLifespan)
	data.Set(oauthEntityRefreshTokenLifespan, oauthEntity.RefreshTokenLifespan)
	data.Set(oauthEntityRedirectURL, oauthEntity.RedirectURL)
	data.Set(oauthEntityLogoURL, oauthEntity.LogoURL)
	data.Set(oauthEntityAccess, oauthEntity.Access)
	data.Set(commonScope, oauthEntity.Scope)
}

func resourceToOAuthEntity(data *schema.ResourceData) *client.OAuthEntity {
	oauthEntity := client.OAuthEntity{
		Name:                 data.Get(oauthEntityName).(string),
		ClientUUID:           data.Get(oauthEntityClientID).(string),
		ClientID:             data.Get(oauthEntityClientID).(string),
		AccessTokenLifespan:  data.Get(oauthEntityAccess).(int),
		RefreshTokenLifespan: data.Get(oauthEntityRefreshTokenLifespan).(int),
		RedirectURL:          data.Get(oauthEntityRedirectURL).(string),
		LogoURL:              data.Get(oauthEntityLogoURL).(string),
		Access:               data.Get(oauthEntityAccess).(string),
	}
	oauthEntity.ID = data.Id()
	oauthEntity.Scope = data.Get(commonScope).(string)
	return &oauthEntity
}
