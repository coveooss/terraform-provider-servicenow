package client

import (
	"fmt"
)

const endpointOAuthEntity = "oauth_entity.do"

// OAuthEntity is the json response for a OAuthEntity in ServiceNow.
type OAuthEntity struct {
	BaseResult
	Name                 string `json:"name"`
	ClientUUID           string `json:"client_uuid"`
	ClientID             string `json:"client_id"`
	AccessTokenLifespan  int    `json:"access_token_lifespan,string"`
	RefreshTokenLifespan int    `json:"refresh_token_lifespan,string"`
	RedirectURL          string `json:"redirect_url"`
	LogoURL              string `json:"logo_url"`
	Access               string `json:"access"`
}

// OAuthEntityResults is the object returned by ServiceNow API when saving or retrieving records.
type OAuthEntityResults struct {
	Records []OAuthEntity `json:"records"`
}

// GetOAuthEntity retrieves a specific OAuthEntity in ServiceNow with it's sys_id.
func (client *ServiceNowClient) GetOAuthEntity(id string) (*OAuthEntity, error) {
	oauthEntityResults := OAuthEntityResults{}
	if err := client.getObject(endpointOAuthEntity, id, &oauthEntityResults); err != nil {
		return nil, err
	}

	return &oauthEntityResults.Records[0], nil
}

// CreateOAuthEntity creates a new OAuthEntity in ServiceNow and returns the newly created page. The new page should
// include the GUID (sys_id) created in ServiceNow.
func (client *ServiceNowClient) CreateOAuthEntity(oauthEntity *OAuthEntity) (*OAuthEntity, error) {
	oauthEntityResults := OAuthEntityResults{}
	if err := client.createObject(endpointOAuthEntity, oauthEntity.Scope, oauthEntity, &oauthEntityResults); err != nil {
		return nil, err
	}

	return &oauthEntityResults.Records[0], nil
}

// UpdateOAuthEntity updates a OAuthEntity in ServiceNow.
func (client *ServiceNowClient) UpdateOAuthEntity(oauthEntity *OAuthEntity) error {
	return client.updateObject(endpointOAuthEntity, oauthEntity.Id, oauthEntity)
}

// DeleteOAuthEntity deletes a OAuthEntity in ServiceNow with the corresponding sys_id.
func (client *ServiceNowClient) DeleteOAuthEntity(id string) error {
	return client.deleteObject(endpointOAuthEntity, id)
}

func (results OAuthEntityResults) validate() error {
	if len(results.Records) <= 0 {
		return fmt.Errorf("no records found")
	} else if len(results.Records) > 1 {
		return fmt.Errorf("more than one record received")
	} else if results.Records[0].Status != "success" {
		return fmt.Errorf("error from ServiceNow -> %s: %s", results.Records[0].Error.Message, results.Records[0].Error.Reason)
	}
	return nil
}
