package client

// EndpointOAuthEntity is the endpoint to manage oauth entity records.
const EndpointOAuthEntity = "oauth_entity.do"

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
