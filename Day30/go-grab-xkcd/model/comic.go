package model

import (
	"encoding/json"
	"fmt"
)

type ComicResponse struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}
type Comic struct {
	Title       string `json:"title"`
	Number      int    `json:"number"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func (r ComicResponse) FormattedDate() string {
	return fmt.Sprintf("%s-%s-%s", r.Year, r.Month, r.Day)
}

func (r ComicResponse) MapToComic() Comic {
	return Comic{
		Title:       r.Title,
		Number:      r.Num,
		Date:        r.FormattedDate(),
		Description: r.Alt,
		Image:       r.Img,
	}
}

func (c Comic) PrettyPrint() string {
	return fmt.Sprintf(
		"Title: %s\nComic No: %d\nDate: %s\nDescription: %s\nImage: %s\n",
		c.Title, c.Number, c.Date, c.Description, c.Image)
}

func (c Comic) ToJSON() string {
	json, err := json.Marshal(c)
	if err != nil {
		return ""
	}

	return string(json)
}
