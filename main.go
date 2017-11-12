package main

import (
	"bufio"
	_ "github.com/bruno-chavez/ancestorquotes/converter"
	"github.com/bruno-chavez/ancestorquotes/slices"
	"github.com/bruno-chavez/ancestorquotes/structs"
	"github.com/urfave/cli"
	"os"
	"time"
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
									c := time.Tick(1 * time.Second)
									for range c {
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
									c := time.Tick(1 * time.Second)
									for range c {
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
	}

	app.Action = func(c *cli.Context) error {
		structs.RandomQuote()
		return nil
	}
	app.Run(os.Args)
	return
}