package client

import (
	"time"
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	ApiKey string
	BaseUrl string
	Client *http.Client
}

type Book struct {
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
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
	return client, nil
}

func WithHttpClient(client *http.Client) (func(*Client) error) {
	return func(c *Client) error {
		c.Client = client
		return nil
	}
}

func (c *Client) CreateBook(title, author string, rating int) (*Book, error) {
	input := Book{Title: title, Author: author, Rating: rating}
	data, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Post("http://localhost:3000/api/v1/Book", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var book Book
	err = json.NewDecoder(resp.Body).Decode(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
