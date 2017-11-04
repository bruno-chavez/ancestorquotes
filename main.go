package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
	"github.com/urfave/cli"
	_"github.com/robfig/cron"
	"github.com/bruno-chavez/ancestorquotes/quotes"
	"github.com/bruno-chavez/ancestorquotes/slices"
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
	app.Commands = []cli.Command{
		{
			Name:    "persistent",
			Usage:   "makes the app execute itself on a regular basis",
			Aliases: []string{"p"},
			Subcommands: []cli.Command{
				{
					Action: func(c *cli.Context) {

					},
					},
					},
					},
	}
	app.Action = func(c *cli.Context) error {

		json.Unmarshal(quotes.Q(), &quoteSlice)
		selectedQuote := quoteSlice[rand.Intn(len(quoteSlice))]
		fmt.Printf("%v", selectedQuote.Quote + "\n")
		return nil
	}
	app.Run(os.Args)
}
