package main

import (
	"fmt"
	"github.com/canro91/30DaysOfGo/Day29/client/client"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Creating book")

	myClient := client.Client{
		Client: &http.Client{},
	}
	_, err := myClient.CreateBook("The Art of War", "Sun Tzu", 5)
	if err != nil {
		log.Fatal(err)
	}
}
