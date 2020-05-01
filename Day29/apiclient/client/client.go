package client

import (
	"fmt"
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

	// Notice you would have to set up a ApiKey header.
	// For example:
	// req, err := http.NewRequest("GET", "http://example.com", nil)
	// req.Header.Add("X-API-KEY", c.ApiKey)
	// resp, err := c.Client.Do(req)

	url := fmt.Sprintf("http://%s/api/v1/Book", c.BaseUrl)
	resp, err := c.Client.Post(url, "application/json", bytes.NewBuffer(data))
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

func (c *Client) GetAllBooks() ([]Book, error) {
	url := fmt.Sprintf("http://%s/api/v1/Book", c.BaseUrl)
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var books []Book
	err = json.NewDecoder(resp.Body).Decode(&books)
	if err != nil {
		return nil, err
	}

	return books, nil
}
