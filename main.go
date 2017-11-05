package main

import (
	"github.com/bruno-chavez/ancestorquotes/random"
	"github.com/bruno-chavez/ancestorquotes/slices"
	"github.com/robfig/cron"
	"github.com/urfave/cli"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "ancestorquotes"
	app.Author = "bruno-chavez"
	app.Usage = "brings quotes from the darkest of dungeons!"
	app.Version = "0.1"
	app.Commands = []cli.Command{
		{
			Name:    "persistent",
			Usage:   "schedules the app to run every x amount of time",
			Aliases: []string{"p"},
			Subcommands: []cli.Command{
				{
					Name:    "minutes",
					Usage:   "minutes between every quote, value accepted if between 1 and 1439",
					Aliases: []string{"m"},
					Action: func(c *cli.Context) error {

						return nil
					},
					Subcommands: []cli.Command{
						{
							Aliases: slices.MinutesDay(),
							Hidden:  true,
							Action: func(c *cli.Context) error {
								schedule := cron.New()
								schedule.AddFunc("", random.RandomQuote)
								schedule.Run()
								return nil
							},
						},
					},
				},
				{
					Name:    "seconds",
					Usage:   "seconds between every quote, value accepted if between 1 and 59",
					Aliases: []string{"m"},
					Action: func(c *cli.Context) error {
						return nil
					},
					Subcommands: []cli.Command{
						{
							Aliases: slices.SecondsMinute(),
							Hidden:  true,
							Action: func(c *cli.Context) error {
								schedule := cron.New()
								schedule.AddFunc("", random.RandomQuote)
								schedule.Run()
								return nil
							},
						},
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
