package firezone

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

const GatewayGroupsUrl string = "/gateway_groups"

type GatewayGroup struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type GatewayGroupsResponse struct {
	Data []GatewayGroup `json:"data"`
}

type GatewayGroupResponse struct {
	Data GatewayGroup `json:"data"`
}

type GatewayGroupRequest struct {
	Gateway_group struct {
		Name string `json:"name"`
	} `json:"gateway_group"`
}

func (c *Client) ReadGatewayGroups() ([]GatewayGroup, error) {
	var gatewayGroupsRes GatewayGroupsResponse

	res, err1 := c.doRequest("GET", GatewayGroupsUrl, nil)
	err2 := json.Unmarshal(res, &gatewayGroupsRes)
	gatewayGroups := gatewayGroupsRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return gatewayGroups, nil
}

func (c *Client) CreateGatewayGroup(name string) (*GatewayGroup, error) {
	var gatewayGroupRes GatewayGroupResponse
	bod := GatewayGroupRequest{}
	bod.Gateway_group.Name = name
	j, err1 := json.Marshal(bod)

	res, err2 := c.doRequest("POST", GatewayGroupsUrl, bytes.NewBuffer(j))
	err3 := json.Unmarshal(res, &gatewayGroupRes)
	gatewayGroup := gatewayGroupRes.Data

	if err := errors.Join(err1, err2, err3); err != nil {
		return nil, err
	}
	return &gatewayGroup, nil
}

func (c *Client) ReadGatewayGroup(id string) (*GatewayGroup, error) {
	var gatewayGroupRes GatewayGroupResponse
	theURL := fmt.Sprintf("%s/%s", GatewayGroupsUrl, id)

	res, err1 := c.doRequest("GET", theURL, nil)
	err2 := json.Unmarshal(res, &gatewayGroupRes)
	gatewayGroup := gatewayGroupRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &gatewayGroup, nil
}

func (c *Client) UpdateGatewayGroup(gG *GatewayGroup) (*GatewayGroup, error) {
	var gatewayGroupRes GatewayGroupResponse
	bod := GatewayGroupRequest{}
	bod.Gateway_group.Name = gG.Name
	j, err1 := json.Marshal(bod)
	theURL := fmt.Sprintf("%s/%s", GatewayGroupsUrl, gG.Id)

	res, err2 := c.doRequest("PUT", theURL, bytes.NewBuffer(j))
	err3 := json.Unmarshal(res, &gatewayGroupRes)
	gatewayGroup := gatewayGroupRes.Data

	if err := errors.Join(err1, err2, err3); err != nil {
		return nil, err
	}
	return &gatewayGroup, nil
}

func (c *Client) DeleteGatewayGroup(id string) (*GatewayGroup, error) {
	var gatewayGroupRes GatewayGroupResponse
	theURL := fmt.Sprintf("%s/%s", GatewayGroupsUrl, id)

	res, err1 := c.doRequest("DELETE", theURL, nil)
	err2 := json.Unmarshal(res, &gatewayGroupRes)
	gatewayGroup := gatewayGroupRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &gatewayGroup, nil
}