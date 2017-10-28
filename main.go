package main

import (
	"fmt"
	"math/rand"
	"os"
	"github.com/urfave/cli"
	"encoding/json"
	"time"
	//"github.com/bruno-chavez/ancestorquotes"
	"ancestorquotes/quotes"
	)

type Quotes struct {
	Quote string `json:"quote"`
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	quoteSlice := make([]Quotes, 0)

	app := cli.NewApp()
	app.Name = "ancestorquotes"
	app.Author = "bruno-chavez"
	app.Usage = "brings quotes from the darkest of dungeons!"
	app.Version = "0.1"
	app.Action = func(c *cli.Context) error {

		json.Unmarshal(quotes.Q(), &quoteSlice)
		selectedQuote := quoteSlice[rand.Intn(len(quoteSlice))]
		fmt.Printf("%v", selectedQuote.Quote)
		return nil
	}
	app.Run(os.Args)
}
