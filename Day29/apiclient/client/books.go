package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type BookService struct {
	client *Client
}

type Book struct {
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func (b *BookService) Post(title, author string, rating int) (*Book, error) {
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

	url := fmt.Sprintf("http://%s/api/v1/Book", b.client.BaseUrl)
	resp, err := b.client.Client.Post(url, "application/json", bytes.NewBuffer(data))
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

func (b *BookService) GetAll() ([]Book, error) {
	url := fmt.Sprintf("http://%s/api/v1/Book", b.client.BaseUrl)
	resp, err := b.client.Client.Get(url)
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