package firezone

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

const IdentitiesUrl string = "/actors"

type Identity struct {
	Id string `json:"id"`
	Actor_id string `json:"actor_id"`
	Provider_id string `json:"provider_id"`
	Provider_identifier string `json:"provider_identifier"`
}

type IdentitiesResponse struct {
	Data []Identity `json:"data"`
}

type IdentityResponse struct {
	Data Identity `json:"data"`
}

type IdentityRequest struct {
	Identity struct {
		Provider_identifier string `json:"provider_identifier"`
	} `json:"identity"`
}

func (c *Client) ReadIdentities(actor_id string) ([]Identity, error) {
	var identitysRes IdentitiesResponse
	theURL := fmt.Sprintf("%s/%s/identities", IdentitiesUrl, actor_id)

	res, err1 := c.doRequest("GET", theURL, nil)
	err2 := json.Unmarshal(res, &identitysRes)
	identitys := identitysRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return identitys, nil
}

func (c *Client) CreateIdentity(actor_id string, provider_id string, identifier string) (*Identity, error) {
	var identityRes IdentityResponse
	bod := IdentityRequest{}
	bod.Identity.Provider_identifier = identifier
	j, err1 := json.Marshal(bod)
	theURL := fmt.Sprintf("%s/%s/providers/%s/identities", IdentitiesUrl, actor_id, provider_id)

	res, err2 := c.doRequest("POST", theURL, bytes.NewBuffer(j))
	err3 := json.Unmarshal(res, &identityRes)
	identity := identityRes.Data

	if err := errors.Join(err1, err2, err3); err != nil {
		return nil, err
	}
	return &identity, nil
}

func (c *Client) ReadIdentity(actor_id string, id string) (*Identity, error) {
	var identityRes IdentityResponse
	theURL := fmt.Sprintf("%s/%s/identities/%s", IdentitiesUrl, actor_id, id)

	res, err1 := c.doRequest("GET", theURL, nil)
	err2 := json.Unmarshal(res, &identityRes)
	identity := identityRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &identity, nil
}

func (c *Client) DeleteIdentity(actor_id string, id string) (*Identity, error) {
	var identityRes IdentityResponse
	theURL := fmt.Sprintf("%s/%s/identities/%s", IdentitiesUrl, actor_id, id)

	res, err1 := c.doRequest("DELETE", theURL, nil)
	err2 := json.Unmarshal(res, &identityRes)
	identity := identityRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &identity, nil
}