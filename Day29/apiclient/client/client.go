package client

import (
	"time"
	"net/http"
)

type Client struct {
	ApiKey string
	BaseUrl string
	Client *http.Client

	Books *BookService
}

func NewClient(apiKey string, opts ...func(*Client) error) (*Client, error) {
	client := &Client{
		ApiKey: apiKey,
		BaseUrl: "localhost:3000",
		Client: &http.Client{ Timeout: 20*time.Second},
	}

	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, err
		}
	}

	client.Books = &BookService{client: client}

	return client, nil
}

func WithHttpClient(client *http.Client) (func(*Client) error) {
	return func(c *Client) error {
		c.Client = client
		return nil
	}
}