package main

import (
	"fmt"
	"github.com/canro91/30DaysOfGo/Day29/apiclient/client"
	"log"
	"strings"
)

func main() {
	fmt.Println("Creating book")

	// Notice you can override the http.Client used internally
	// myClient, _ := client.NewClient("XYZ", client.WithHttpClient(&http.Client{}))
	myClient, _ := client.NewClient("XYZ")

	_, err := myClient.Books.Post("The Art of War", "Sun Tzu", 5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Querying all the books")
	books, err := myClient.Books.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, book := range books {
		fmt.Printf("%q by %s %s\n", book.Title, book.Author, strings.Repeat("*", book.Rating))
	}
}
