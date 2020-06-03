package client

import (
	"encoding/json"
	"fmt"
	"github.com/canro91/30DaysOfGo/Day30/go-grab-xkcd/model"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	BaseUrl        string        = "https://xkcd.com"
	DefaultTimeout time.Duration = 30 * time.Second
	LatestComic    ComicNumber   = 0
)

type ComicNumber int

type XKCDClient struct {
	client  *http.Client
	baseUrl string
}

func NewXKCDClient() *XKCDClient {
	return &XKCDClient{
		client: &http.Client{
			Timeout: DefaultTimeout,
		},
		baseUrl: BaseUrl,
	}
}

func (c *XKCDClient) SetTimeout(d time.Duration) {
	c.client.Timeout = d
}

func (c *XKCDClient) Fetch(n ComicNumber, save bool) (model.Comic, error) {
	response, err := c.client.Get(c.BuildUrl(n))
	if err != nil {
		return model.Comic{}, err
	}
	defer response.Body.Close()

	var comicResponse model.ComicResponse
	if err := json.NewDecoder(response.Body).Decode(&comicResponse); err != nil {
		return model.Comic{}, err
	}

	if save {
		if err := c.SaveToDisk(comicResponse.Img, "."); err != nil {
			fmt.Println("Failed to saved image")
		}
	}

	return comicResponse.MapToComic(), nil
}

func (c *XKCDClient) SaveToDisk(url, savePath string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	absolutePath, err := filepath.Abs(savePath)
	filePath := fmt.Sprintf("%s/%s", absolutePath, path.Base(url))
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func (c *XKCDClient) BuildUrl(n ComicNumber) string {
	var url string
	if n == LatestComic {
		url = fmt.Sprintf("%s/info.0.json", BaseUrl)
	} else {
		url = fmt.Sprintf("%s/%d/info.0.json", BaseUrl, n)
	}
	return url
}
