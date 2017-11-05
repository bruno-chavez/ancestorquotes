package main

import (
	"github.com/bruno-chavez/ancestorquotes/random"
	"github.com/bruno-chavez/ancestorquotes/slices"
	"github.com/robfig/cron"
	"github.com/urfave/cli"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	app := cli.NewApp()
	app.Name = "ancestorquotes"
	app.Author = "bruno-chavez"
	app.Usage = "brings quotes from the darkest of dungeons!"
	app.Version = "0.1"
	app.Commands = []cli.Command{
		{
			Name:    "persistent",
			Usage:   "schedules the app to run every x amount of minutes",
			Aliases: []string{"p"},
			Subcommands: []cli.Command{
				{
					Name:    "minutes",
					Usage:   "sets the interval for the scheduled run",
					Aliases: slices.MinutesDay(),
					Action: func(c *cli.Context) error {
						schedule := cron.New()
						schedule.AddFunc("1 * * * * *", random.RandomQuote)
						schedule.Run()
						return nil
					},
				},
				{
					Name:    "stop",
					Usage:   "stops the current scheduled run",
					Aliases: []string{"s"},
					Action: func(c *cli.Context) error {
						schedule := cron.New()
						schedule.Stop()
						return nil
					},
				},
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		random.RandomQuote()
		return nil
	}

	app.Run(os.Args)
}
