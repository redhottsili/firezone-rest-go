package firezone

import (
	"encoding/json"
	"errors"
	"fmt"
)

const TokensUrl string = "/gateway_groups"

type Token struct {
	Id string `json:"id"`
	Token string `json:"token"`
}

type TokensResponse struct {
	Data []Token `json:"data"`
}

type TokenResponse struct {
	Data Token `json:"data"`
}

func (c *Client) CreateToken(gateway_group_id string) (*Token, error) {
	var tokenRes TokenResponse
	theURL := fmt.Sprintf("%s/%s/tokens", TokensUrl, gateway_group_id)

	res, err1 := c.doRequest("POST", theURL, nil)
	err2 := json.Unmarshal(res, &tokenRes)
	token := tokenRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &token, nil
}

func (c *Client) DeleteTokens(gateway_group_id string) ([]Token, error) {
	var tokenRes TokensResponse
	theURL := fmt.Sprintf("%s/%s/tokens", TokensUrl, gateway_group_id)

	res, err1 := c.doRequest("DELETE", theURL, nil)
	err2 := json.Unmarshal(res, &tokenRes)
	tokens := tokenRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return tokens, nil
}

func (c *Client) DeleteToken(gateway_group_id string, id string) (*Token, error) {
	var tokenRes TokenResponse
	theURL := fmt.Sprintf("%s/%s/tokens/%s", TokensUrl, gateway_group_id, id)

	res, err1 := c.doRequest("DELETE", theURL, nil)
	err2 := json.Unmarshal(res, &tokenRes)
	token := tokenRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &token, nil
}