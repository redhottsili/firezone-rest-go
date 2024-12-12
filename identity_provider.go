package firezone

import (
	"encoding/json"
	"errors"
	"fmt"
)

const IdentityProvidersUrl string = "/identity_providers"

type IdentityProvider struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type IdentityProvidersResponse struct {
	Data []IdentityProvider `json:"data"`
}

type IdentityProviderResponse struct {
	Data IdentityProvider `json:"data"`
}

func (c *Client) ReadIdentityProviders() ([]IdentityProvider, error) {
	var identityProvidersRes IdentityProvidersResponse

	res, err1 := c.doRequest("GET", IdentityProvidersUrl, nil)
	err2 := json.Unmarshal(res, &identityProvidersRes)
	identityProviders := identityProvidersRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return identityProviders, nil
}

func (c *Client) ReadIdentityProvider(id string) (*IdentityProvider, error) {
	var identityProviderRes IdentityProviderResponse
	theURL := fmt.Sprintf("%s/%s", IdentityProvidersUrl, id)

	res, err1 := c.doRequest("GET", theURL, nil)
	err2 := json.Unmarshal(res, &identityProviderRes)
	identityProvider := identityProviderRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &identityProvider, nil
}

func (c *Client) DeleteIdentityProvider(id string) (*IdentityProvider, error) {
	var identityProviderRes IdentityProviderResponse
	theURL := fmt.Sprintf("%s/%s", ActorsUrl, id)

	res, err1 := c.doRequest("DELETE", theURL, nil)
	err2 := json.Unmarshal(res, &identityProviderRes)
	identityProvider := identityProviderRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &identityProvider, nil
}