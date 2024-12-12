package firezone

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

const ActorsUrl string = "/actors"

type ActorType string

const (
	ADMIN ActorType = "account_admin_user"
	USER ActorType = "account_user"
	SERVICE ActorType = "api_client"
)

type Actor struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Type ActorType  `json:"type"`
}

type ActorsResponse struct {
	Data []Actor `json:"data"`
}

type ActorResponse struct {
	Data Actor `json:"data"`
}

type ActorRequest struct {
	Actor struct {
		Name string `json:"name"`
		Type ActorType `json:"type"`
	} `json:"actor"`
}

func (c *Client) ReadActors() ([]Actor, error) {
	var actorsRes ActorsResponse

	res, err1 := c.doRequest("GET", ActorsUrl, nil)
	err2 := json.Unmarshal(res, &actorsRes)
	actors := actorsRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return actors, nil
}

func (c *Client) CreateActor(a *Actor) (*Actor, error) {
	var actorRes ActorResponse
	bod := ActorRequest{}
	bod.Actor.Name = a.Name
	bod.Actor.Type = a.Type
	j, err1 := json.Marshal(bod)

	res, err2 := c.doRequest("POST", ActorsUrl, bytes.NewBuffer(j))
	err3 := json.Unmarshal(res, &actorRes)
	actor := actorRes.Data

	if err := errors.Join(err1, err2, err3); err != nil {
		return nil, err
	}
	return &actor, nil
}

func (c *Client) ReadActor(id string) (*Actor, error) {
	var actorRes ActorResponse
	theURL := fmt.Sprintf("%s/%s", ActorsUrl, id)

	res, err1 := c.doRequest("GET", theURL, nil)
	err2 := json.Unmarshal(res, &actorRes)
	actor := actorRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &actor, nil
}

func (c *Client) UpdateActor(a *Actor) (*Actor, error) {
	var actorRes ActorResponse
	bod := ActorRequest{}
	bod.Actor.Name = a.Name
	bod.Actor.Type = a.Type
	j, err1 := json.Marshal(bod)
	theURL := fmt.Sprintf("%s/%s", ActorsUrl, a.Id)

	res, err2 := c.doRequest("PUT", theURL, bytes.NewBuffer(j))
	err3 := json.Unmarshal(res, &actorRes)
	actor := actorRes.Data

	if err := errors.Join(err1, err2, err3); err != nil {
		return nil, err
	}
	return &actor, nil
}

func (c *Client) DeleteActor(id string) (*Actor, error) {
	var actorRes ActorResponse
	theURL := fmt.Sprintf("%s/%s", ActorsUrl, id)

	res, err1 := c.doRequest("DELETE", theURL, nil)
	err2 := json.Unmarshal(res, &actorRes)
	actor := actorRes.Data

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}
	return &actor, nil
}
