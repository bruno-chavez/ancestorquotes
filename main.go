package main

import (
	"bufio"
	"os"
	"strconv"
	"time"

	"github.com/bruno-chavez/ancestorquotes/slices"
	"github.com/bruno-chavez/ancestorquotes/structs"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "ancestorquotes"
	app.Author = "bruno-chavez"
	app.Usage = "brings quotes from the darkest of dungeons!"
	app.Version = "0.2"
	app.Commands = []cli.Command{
		{
			Name:    "persistent",
			Usage:   "schedules the app to run every x amount of time",
			Aliases: []string{"p"},
			Subcommands: []cli.Command{
				{
					Name:    "minutes",
					Usage:   "minutes between every quote, value accepted if between 1 and 59",
					Aliases: []string{"m"},
					Action: func(c *cli.Context) error {

						return nil
					},
					Subcommands: []cli.Command{
						{
							Aliases: slices.SecondsMinutes(),
							Hidden:  true,
							Action: func(c *cli.Context) error {

								go func() {
									timer, _ := strconv.Atoi(os.Args[3])
									ticking := time.Tick(time.Duration(timer) * time.Minute)
									for range ticking {
										structs.RandomQuote()
									}
								}()

								scanner := bufio.NewScanner(os.Stdin)
								for scanner.Scan() {
									if scanner.Text() == "stop" {
										return nil
									}
								}
								return nil
							},
						},
					},
				},
				{
					Name:    "seconds",
					Usage:   "seconds between every quote, value accepted if between 1 and 59",
					Aliases: []string{"s"},
					Action: func(c *cli.Context) error {
						return nil
					},
					Subcommands: []cli.Command{
						{
							Aliases: slices.SecondsMinutes(),
							Hidden:  true,
							Action: func(c *cli.Context) error {
								go func() {
									timer, _ := strconv.Atoi(os.Args[3])
									ticking := time.Tick(time.Duration(timer) * time.Second)
									for range ticking {
										structs.RandomQuote()
									}
								}()

								scanner := bufio.NewScanner(os.Stdin)
								for scanner.Scan() {
									if scanner.Text() == "stop" {
										return nil
									}
								}
								return nil
							},
						},
					},
				},
			},
		},
		{
			Name:    "all",
			Usage:   "Prints all quotes from the darkest of dungeons!",
			Aliases: []string{"a"},
			Action: func(c *cli.Context) error {
				structs.AllQuotes()
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		structs.RandomQuote()
		return nil
	}
	app.Run(os.Args)
	return
}
