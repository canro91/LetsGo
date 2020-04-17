package main

import (
	"strconv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
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

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	request, err := http.NewRequest("GET", "http://www.reddit.com/r/wallpaper/top/.json?sort=new&limit=1", nil)
	request.Header.Add("User-Agent", "Wallie/1.0")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Fprintln(os.Stderr, "We can't retrieve top posts from /r/wallpaper. You have internet connection?")
		return err
	}

	body, _ := ioutil.ReadAll(response.Body)
	var topResponse TopResponse
	json.Unmarshal(body, &topResponse)

	post := topResponse.Data.Children[0].Data
	fmt.Println(post.Title)
	fmt.Println(post.URL)
	fmt.Println(post.Ups)

	response, err = http.Get(post.URL)
	if err != nil {
		fmt.Fprintln(os.Stderr, "We can't retrieve the wallpaper. You have internet connection?")
	}
	defer response.Body.Close()

	var re = regexp.MustCompile(`http(s)?://` + post.Domain + `/`)
	filename := re.ReplaceAllString(post.URL, "")

	path := "img/" + filename
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	dir, err := os.Getwd()

	command := exec.Command("osascript", "-e", `tell application "System Events" to tell every desktop to set picture to `+strconv.Quote(dir+"/"+path))
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()

	return nil
}
