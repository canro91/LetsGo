package main

import (
	"os"
	"log"
	"context"
	"github.com/shomali11/slacker"
)

func main() {
	bot := slacker.NewClient(os.Getenv("SLACK_API_TOKEN"))

	pingPong := &slacker.CommandDefinition{
		Description: "Say Pong back!",
		Example:     "Ping",
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Pong!")
		},
	}
	bot.Command("ping", pingPong)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
