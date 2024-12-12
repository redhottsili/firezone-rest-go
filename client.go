package firezone

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

const DefaultRestUrl string = "https://api.firezone.dev"

type Client struct {
	HttpClient *http.Client
	ApiKey     string
	Host       string
	Base       string
}

type GenericData interface {
	GetData() any
}

func (c *Client) doRequest(method string, path string, data io.Reader) ([]byte, error) {
	url := DefaultRestUrl + path
	request, err1 := http.NewRequest(method, url, data)
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))
	if method == "PUT" || method == "POST" {
		request.Header.Set("Content-Type", "application/json")
	}
	res, err2 := c.HttpClient.Do(request)

	defer res.Body.Close()
	body, err3 := io.ReadAll(res.Body)

	if err := errors.Join(err1, err2, err3); err != nil {
		return nil, err
	}

	if (res.StatusCode >= 200 && res.StatusCode < 300) || res.StatusCode == http.StatusNoContent {
		return body, nil
	} else {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
}

func NewClient(apiKey string) *Client {
	return &Client{
		HttpClient: http.DefaultClient,
		ApiKey:     apiKey,
	}
}