package firezone

import (
	"encoding/json"
	"errors"
	"fmt"
)

const GatewaysUrl string = "/gateway_groups"

type Gateway struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Ipv4 string `json:"ipv4"`
	Ipv6 string `json:"ipv6"`
	Online bool `json:"online"`
}

type GatewaysResponse struct {
	Data []Gateway `json:"data"`
}

type GatewayResponse struct {
	Data Gateway `json:"data"`
}

func (c *Client) ReadGateways(gateway_group_id string) ([]Gateway, error) {
	var gatewaysRes GatewaysResponse
	theURL := fmt.Sprintf("%s/%s/gateways", GatewaysUrl, gateway_group_id)

	res, err1 := c.doRequest("GET", theURL, nil)
	err2 := json.Unmarshal(res, &gatewaysRes)
	gateways := gatewaysRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return gateways, nil
}

func (c *Client) ReadGateway(gateway_group_id string, id string) (*Gateway, error) {
	var gatewayRes GatewayResponse
	theURL := fmt.Sprintf("%s/%s/gateways/%s", GatewaysUrl, gateway_group_id, id)

	res, err1 := c.doRequest("GET", theURL, nil)
	err2 := json.Unmarshal(res, &gatewayRes)
	gateway := gatewayRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &gateway, nil
}

func (c *Client) DeleteGateway(gateway_group_id string, id string) (*Gateway, error) {
	var gatewayRes GatewayResponse
	theURL := fmt.Sprintf("%s/%s/gateways/%s", GatewaysUrl, gateway_group_id, id)

	res, err1 := c.doRequest("DELETE", theURL, nil)
	err2 := json.Unmarshal(res, &gatewayRes)
	gateway := gatewayRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &gateway, nil
}