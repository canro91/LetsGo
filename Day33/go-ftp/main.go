package main

import (
	"fmt"
	"github.com/canro91/30DaysOfGo/Day33/go-ftp/client"
	"log"
)

func main() {
	client := client.NewClient("mirror.us.leaseweb.net", 21)
	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	fmt.Println("*** Connected")

	err = client.AnonymousLogin()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("*** LoggedIn")

	fmt.Printf("*** Listing files:\n")
	message, err := client.List()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", message)
}
