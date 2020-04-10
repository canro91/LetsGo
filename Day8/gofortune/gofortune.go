package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"
)

type Response struct {
	Quote Quote `json:"quote"`
}

type Quote struct {
	Body string `json:"body"`
}

func main(){
	response, err := http.Get("https://favqs.com/api/qotd")
	if err != nil {
		fmt.Println("We can't retrieve your quote. You have internet connection?")
	} else {
		rawQuotes, _ := ioutil.ReadAll(response.Body)
		var quote Response
		json.Unmarshal(rawQuotes, &quote)
		fmt.Println(quote.Quote.Body)
	}
}