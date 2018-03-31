package main

import (
	"bufio"
	"github.com/bruno-chavez/ancestorquotes/slices"
	"github.com/bruno-chavez/ancestorquotes/structs"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"time"
)

func main() {
	var search cli.StringSlice
	filters := func() []structs.Filter {
		var result []structs.Filter
		for _, s := range search {
			result = append(result, structs.Contains(s))
		}
		return result
	}

	app := cli.NewApp()
	app.Name = "ancestorquotes"
	app.Author = "bruno-chavez"
	app.Usage = "brings quotes from the darkest of dungeons!"
	app.Version = "0.5"
	app.Flags = []cli.Flag{
		cli.StringSliceFlag{
			Name:  "search, s",
			Value: &search,
			Usage: "search term that the quote must contain",
		},
	}
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
										structs.Quotes.Filter(filters()...).Random().Print()
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
										structs.Quotes.Filter(filters()...).Random().Print()
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
		{
			Name:    "chat",
			Usage:   "The Ancestor talks with himself in a maddenly fashion.",
			Aliases: []string{"c"},
			Action: func(c *cli.Context) error {
				structs.Chat()
				return nil
			},
		},
                {
			Name: "talkback",
                        Usage: "You can talk to the Ancestor and the Ancestor replies back in a crazy manner.",
                        Aliases: []string{"t"},
                        Action: func(c *cli.Context) error {
                                structs.TalkBack()
                                return nil
                        },
                },
	}

	app.Action = func(c *cli.Context) error {
		structs.Quotes.Filter(filters()...).Random().Print()
		return nil
	}
	app.Run(os.Args)
	return
}
