package firezone

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

const PoliciesUrl string = "/policies"

type Policy struct {
	Id string `json:"id"`
	Description string `json:"description"`
	Actor_group_id string `json:"actor_group_id"`
	Resource_id string `json:"resource_id"`
}

type PoliciesResponse struct {
	Data []Policy `json:"data"`
}

type PolicyResponse struct {
	Data Policy `json:"data"`
}

type PolicyRequest struct {
	Policy struct {
		Description string `json:"description"`
		Actor_group_id string `json:"actor_group_id"`
		Resource_id string `json:"resource_id"`
	} `json:"policy"`
}

func (c *Client) ReadPolicies() ([]Policy, error) {
	var policiesRes PoliciesResponse

	res, err1 := c.doRequest("GET", PoliciesUrl, nil)
	err2 := json.Unmarshal(res, &policiesRes)
	policies := policiesRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return policies, nil
}

func (c *Client) CreatePolicy(p Policy) (*Policy, error) {
	var policyRes PolicyResponse
	bod := PolicyRequest{}
	bod.Policy.Description = p.Description
	bod.Policy.Actor_group_id = p.Actor_group_id
	bod.Policy.Resource_id = p.Resource_id
	j, err1 := json.Marshal(bod)

	res, err2 := c.doRequest("POST", PoliciesUrl, bytes.NewBuffer(j))
	err3 := json.Unmarshal(res, &policyRes)
	policy := policyRes.Data

	if err := errors.Join(err1, err2, err3); err != nil {
		return nil, err
	}
	return &policy, nil
}

func (c *Client) ReadPolicy(id string) (*Policy, error) {
	var policyRes PolicyResponse
	theURL := fmt.Sprintf("%s/%s", PoliciesUrl, id)

	res, err1 := c.doRequest("GET", theURL, nil)
	err2 := json.Unmarshal(res, &policyRes)
	policy := policyRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &policy, nil
}

func (c *Client) UpdatePolicy(p *Policy) (*Policy, error) {
	var policyRes PolicyResponse
	bod := PolicyRequest{}
	bod.Policy.Description = p.Description
	bod.Policy.Actor_group_id = p.Actor_group_id
	bod.Policy.Resource_id = p.Resource_id
	j, err1 := json.Marshal(bod)
	theURL := fmt.Sprintf("%s/%s", PoliciesUrl, p.Id)

	res, err2 := c.doRequest("PUT", theURL, bytes.NewBuffer(j))
	err3 := json.Unmarshal(res, &policyRes)
	policy := policyRes.Data

	if err := errors.Join(err1, err2, err3); err != nil {
		return nil, err
	}
	return &policy, nil
}

func (c *Client) DeletePolicy(id string) (*Policy, error) {
	var policyRes PolicyResponse
	theURL := fmt.Sprintf("%s/%s", PoliciesUrl, id)

	res, err1 := c.doRequest("DELETE", theURL, nil)
	err2 := json.Unmarshal(res, &policyRes)
	policy := policyRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &policy, nil
}