package firezone

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

const ActorGroupMembershipsUrl string = "/actor_groups"

type ActorGroupMembershipsResponse struct {
	Data []Actor `json:"data"`
}

type ActorGroupMembershipsRequest struct {
	Memberships []ActorID `json:"memberships"`
}

type ActorID struct {
	Actor_id string `json:"actor_id"`
}

type ActorGroupMembershipsUpdateResponse struct {
	Data struct {
		Actor_ids []string `json:"actor_ids"`
	} `json:"data"`
}


func (c *Client) ReadActorGroupMemberships(actor_group_id string) ([]Actor, error) {
	var actorGroupMembershipsRes ActorGroupMembershipsResponse
	theURL := fmt.Sprintf("%s/%s/memberships", ActorGroupMembershipsUrl, actor_group_id)

	res, err1 := c.doRequest("GET", theURL, nil)
	err2 := json.Unmarshal(res, &actorGroupMembershipsRes)
	actorGroupMemberships := actorGroupMembershipsRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return actorGroupMemberships, nil
}

func (c *Client) UpdateActorGroupMemberships(actor_group_id string, members []string) ([]string, error) {
	var actorGroupMembershipsRes ActorGroupMembershipsUpdateResponse
	bod := ActorGroupMembershipsRequest{}
	for _, a := range members {
		bod.Memberships = append(bod.Memberships, ActorID{Actor_id: a})
	}
	j, err1 := json.Marshal(bod)
	theURL := fmt.Sprintf("%s/%s/memberships", ActorGroupMembershipsUrl, actor_group_id)

	res, err2 := c.doRequest("PUT", theURL, bytes.NewBuffer(j))
	err3 := json.Unmarshal(res, &actorGroupMembershipsRes)
	actorGroupMemberships := actorGroupMembershipsRes.Data.Actor_ids

	if err := errors.Join(err1, err2, err3); err != nil {
		return nil, err
	}
	return actorGroupMemberships, nil
}