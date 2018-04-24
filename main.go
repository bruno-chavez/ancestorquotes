package main

import (
	"bufio"
	"fmt"
	"github.com/bruno-chavez/ancestorquotes/commands"
	"github.com/bruno-chavez/ancestorquotes/quotes"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"time"
)

//Contains all the quotes in a QuoteSlice type slice, used as a global variable to avoid having multiples of the same slice, reducing load times and memory usage.
var quoteSLice = quotes.Parse()

func main() {

	app := cli.NewApp()
	app.Name = "ancestorquotes"
	app.Author = "bruno-chavez"
	app.Usage = "brings quotes from the darkest of dungeons!"
	app.Version = "1.0"
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
							Aliases: commands.SliceSecondsMinutes(),
							Hidden:  true,
							Action: func(c *cli.Context) error {

								go func() {
									timer, _ := strconv.Atoi(os.Args[3])
									ticking := time.Tick(time.Duration(timer) * time.Minute)
									for range ticking {
										fmt.Println(quoteSLice.RandomQuote())
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
							Aliases: commands.SliceSecondsMinutes(),
							Hidden:  true,
							Action: func(c *cli.Context) error {
								go func() {
									timer, _ := strconv.Atoi(os.Args[3])
									ticking := time.Tick(time.Duration(timer) * time.Second)
									for range ticking {
										fmt.Println(quoteSLice.RandomQuote())
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
				commands.AllQuotes(quoteSLice)
				return nil
			},
		},
		{
			Name:    "chat",
			Usage:   "The Ancestor talks with himself in a maddening fashion.",
			Aliases: []string{"c"},
			Action: func(c *cli.Context) error {
				commands.Chat(quoteSLice)
				return nil
			},
		},
		{
			Name:    "talkback",
			Usage:   "You can talk to the Ancestor and the Ancestor replies back in a crazy manner.",
			Aliases: []string{"t"},
			Action: func(c *cli.Context) error {
				commands.TalkBack(quoteSLice)
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(c.Args()) > 0 {
			fmt.Printf("%v is not a valid command.\nEnter ancestorquotes --help to see list of valid commands", c.Args()[0])
		} else {
			fmt.Println(quoteSLice.RandomQuote())
		}
		return nil
	}
	app.Run(os.Args)
	return
}
