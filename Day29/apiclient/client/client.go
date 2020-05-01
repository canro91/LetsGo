package client

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	Client *http.Client
}

type Book struct {
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
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
