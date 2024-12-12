package firezone

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

const ActorGroupsUrl string = "/actor_groups"

type ActorGroup struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type ActorGroupsResponse struct {
	Data []ActorGroup `json:"data"`
}

type ActorGroupResponse struct {
	Data ActorGroup `json:"data"`
}

type ActorGroupRequest struct {
	Actor_group struct {
		Name string `json:"name"`
	} `json:"actor_group"`
}

func (c *Client) ReadActorGroups() ([]ActorGroup, error) {
	var actorGroupsRes ActorGroupsResponse

	res, err1 := c.doRequest("GET", ActorGroupsUrl, nil)
	err2 := json.Unmarshal(res, &actorGroupsRes)
	actorGroups := actorGroupsRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return actorGroups, nil
}

func (c *Client) CreateActorGroup(name string) (*ActorGroup, error) {
	var actorGroupRes ActorGroupResponse
	bod := ActorGroupRequest{}
	bod.Actor_group.Name = name
	j, err1 := json.Marshal(bod)

	res, err2 := c.doRequest("POST", ActorGroupsUrl, bytes.NewBuffer(j))
	err3 := json.Unmarshal(res, &actorGroupRes)
	actorGroup := actorGroupRes.Data

	if err := errors.Join(err1, err2, err3); err != nil {
		return nil, err
	}
	return &actorGroup, nil
}

func (c *Client) ReadActorGroup(id string) (*ActorGroup, error) {
	var actorGroupRes ActorGroupResponse
	theURL := fmt.Sprintf("%s/%s", ActorGroupsUrl, id)

	res, err1 := c.doRequest("GET", theURL, nil)
	err2 := json.Unmarshal(res, &actorGroupRes)
	actorGroup := actorGroupRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &actorGroup, nil
}

func (c *Client) UpdateActorGroup(aG *ActorGroup) (*ActorGroup, error) {
	var actorGroupRes ActorGroupResponse
	bod := ActorGroupRequest{}
	bod.Actor_group.Name = aG.Name
	j, err1 := json.Marshal(bod)
	theURL := fmt.Sprintf("%s/%s", ActorGroupsUrl, aG.Id)

	res, err2 := c.doRequest("PUT", theURL, bytes.NewBuffer(j))
	err3 := json.Unmarshal(res, &actorGroupRes)
	actorGroup := actorGroupRes.Data

	if err := errors.Join(err1, err2, err3); err != nil {
		return nil, err
	}
	return &actorGroup, nil
}

func (c *Client) DeleteActorGroup(id string) (*ActorGroup, error) {
	var actorGroupRes ActorGroupResponse
	theURL := fmt.Sprintf("%s/%s", ActorGroupsUrl, id)

	res, err1 := c.doRequest("DELETE", theURL, nil)
	err2 := json.Unmarshal(res, &actorGroupRes)
	actorGroup := actorGroupRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &actorGroup, nil
}