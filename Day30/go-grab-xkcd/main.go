package main

import (
	"flag"
	"fmt"
	"github.com/canro91/30DaysOfGo/Day30/go-grab-xkcd/client"
	"log"
	"time"
)

func main() {
	comicNumber := flag.Int("n", int(client.LatestComic), "Comic number")
	clientTimeout := flag.Int64("t", int64(client.DefaultTimeout.Seconds()), "Client timeout in seconds")
	saveImage := flag.Bool("s", false, "Save image to current directory")
	outputType := flag.String("o", "text", "Print output in format: text/json")
	flag.Parse()

	xkcdClient := client.NewXKCDClient()
	xkcdClient.SetTimeout(time.Duration(*clientTimeout) * time.Second)

	comic, err := xkcdClient.Fetch(client.ComicNumber(*comicNumber), *saveImage)
	if err != nil {
		log.Println(err)
	}

	if *outputType == "json" {
		fmt.Println(comic.ToJSON())
	} else {
		fmt.Println(comic.PrettyPrint())
	}
}
