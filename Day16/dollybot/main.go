package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/shomali11/slacker"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func downloadAndParsePage(quantity int, code string) (string, error) {

	var pageName string

	switch code {
	case "USD", "usd":
		pageName = "dolar"

	case "EURO", "euro":
		pageName = "euro"

	case "CHF", "chf":
		pageName = "franco-suizo"
	}

	if len(pageName) == 0 {
		return "", errors.New("Currency code not supported")
	}

	url := fmt.Sprintf("https://dolar.wilkinsonpc.com.co/divisas/%s.html", pageName)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}

	converted := doc.Find("#indicador_vigente .caja .valor .numero").Text()
	return converted, nil
}

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

	convert := &slacker.CommandDefinition{
		Description: "Convert from any currency to COP",
		Example:     "conv 10 USD",
		Handler: func(request slacker.Request, response slacker.ResponseWriter) {
			// If quantity is missing, use 1
			// If currency code is missing, use USD
			quantity := request.IntegerParam("quantity", 1)
			code := request.StringParam("code", "USD")

			// Show dollybot is typing...
			response.Typing()

			converted, error := downloadAndParsePage(quantity, code)
			if error != nil {
				response.ReportError(errors.New("Oops! Something went wrong. I can't find a conversion"))
			}

			v := strings.Replace(converted, ",", "", -1)
			if s, err := strconv.ParseFloat(v, 64); err == nil {
				value := s * float64(quantity)
				response.Reply(fmt.Sprintf("%d %s = %.2f COP", quantity, code, value))
			} else {
				response.Reply(fmt.Sprintf("1 %s = %s COP", code, converted))
			}
		},
	}
	bot.Command("conv <quantity> <code>", convert)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
