package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

type TopResponse struct {
	Data struct {
		Children []struct {
			Data struct {
				Title  string `json:"title"`
				Ups    int    `json:"ups"`
				Domain string `json:"domain"`
				URL    string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

func downloadTopPost() (TopResponse, error) {
	request, err := http.NewRequest("GET", "http://www.reddit.com/r/wallpaper/top/.json?sort=new&limit=1", nil)
	request.Header.Add("User-Agent", "Wallie/1.0")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return TopResponse{}, err
	}

	body, _ := ioutil.ReadAll(response.Body)
	var topResponse TopResponse
	json.Unmarshal(body, &topResponse)
	return topResponse, nil
}

func getFilename(url, domain string) string {
	re := regexp.MustCompile(`http(s)?://` + domain + `/`)
	filename := re.ReplaceAllString(url, "")
	return filename
}

func downloadImage(url, path string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(path)
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

func changeWallpaper(path string) {
	exec.Command("osascript", "-e", `tell application "System Events" to tell every desktop to set picture to `+strconv.Quote(path)).Run()
}

func main() {
	topResponse, err := downloadTopPost()
	if err != nil {
		log.Fatal(err)
	}

	post := topResponse.Data.Children[0].Data
	filename := getFilename(post.URL, post.Domain)
	currentDir, _ := os.Getwd()
	absolutePath := currentDir + "/img/" + filename
	err = downloadImage(post.URL, absolutePath)
	if err != nil {
		log.Fatal(err)
	}

	changeWallpaper(absolutePath)
}
