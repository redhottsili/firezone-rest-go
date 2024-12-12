package firezone

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

//WIP

const ResourcesUrl string = "/resources"

type ResourceType string

const (
	IP ResourceType = "ip"
	DNS ResourceType = "dns"
	CIDR ResourceType = "cidr"
)

type Resource struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type ResourceType `json:"type"`
	Address string `json:"address"`
	Address_description string `json:"address_description" json:"address"`
}

type ResourcesResponse struct {
	Data []Resource `json:"data"`
}

type ResourceResponse struct {
	Data Resource `json:"data"`
}

type ResourceRequest struct {
	Resource struct {
		Name string `json:"name"`
		Type ResourceType `json:"type"`
		Address string `json:"address"`
		Address_description string `json:"address_description"`
		Connections []Connection `json:"connections"`
	} `json:"resource"`
}

type Connection struct {
	Gateway_group_id string `json:"gateway_group_id"`
}

func (c *Client) ReadResources() ([]Resource, error) {
	var resourceRes ResourcesResponse

	res, err1 := c.doRequest("GET", ResourcesUrl, nil)
	err2 := json.Unmarshal(res, &resourceRes)
	resources := resourceRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return resources, nil
}

func (c *Client) CreateResource(r Resource, conns []string) (*Resource, error) {
	var resourceRes ResourceResponse
	bod := ResourceRequest{}
	bod.Resource.Name = r.Name
	bod.Resource.Type = r.Type
	bod.Resource.Address = r.Address
	bod.Resource.Address_description = r.Address_description
	for _, conn := range conns {
		bod.Resource.Connections = append(
			bod.Resource.Connections,
			Connection{Gateway_group_id: conn},
		)
	}
	j, err1 := json.Marshal(bod)

	res, err2 := c.doRequest("POST", ResourcesUrl, bytes.NewBuffer(j))
	err3 := json.Unmarshal(res, &resourceRes)
	resource := resourceRes.Data

	if err := errors.Join(err1, err2, err3); err != nil {
		return nil, err
	}
	return &resource, nil
}

func (c *Client) ReadResource(id string) (*Resource, error) {
	var resourceRes ResourceResponse
	theURL := fmt.Sprintf("%s/%s", ResourcesUrl, id)

	res, err1 := c.doRequest("GET", theURL, nil)
	err2 := json.Unmarshal(res, &resourceRes)
	resource := resourceRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &resource, nil
}

func (c *Client) UpdateResource(r Resource, conns []string) (*Resource, error) {
	var resourceRes ResourceResponse
	bod := ResourceRequest{}
	bod.Resource.Name = r.Name
	bod.Resource.Type = r.Type	
	bod.Resource.Address = r.Address
	bod.Resource.Address_description = r.Address_description
	for _, conn := range conns {
		bod.Resource.Connections = append(
			bod.Resource.Connections,
			Connection{Gateway_group_id: conn},
		)
	}
	j, err1 := json.Marshal(bod)
	fmt.Println(string(j))
	theURL := fmt.Sprintf("%s/%s", ResourcesUrl, r.Id)

	res, err2 := c.doRequest("PUT", theURL, bytes.NewBuffer(j))
	err3 := json.Unmarshal(res, &resourceRes)
	resource := resourceRes.Data

	if err := errors.Join(err1, err2, err3); err != nil {
		return nil, err
	}
	return &resource, nil
}

func (c *Client) DeleteResource(id string) (*Resource, error) {
	var resourceRes ResourceResponse
	theURL := fmt.Sprintf("%s/%s", ResourcesUrl, id)

	res, err1 := c.doRequest("DELETE", theURL, nil)
	err2 := json.Unmarshal(res, &resourceRes)
	resource := resourceRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &resource, nil
}