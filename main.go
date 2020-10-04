package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"ancestorquotes/commands"

	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "ancestorquotes"
	app.Author = "bruno-chavez"
	app.Usage = "Brings quotes from the darkest of dungeons!"
	app.Version = "2.1.0"
	app.Commands = []cli.Command{
		{
			Name:    "persistent",
			Usage:   "Makes the Ancestor say a quote every certain amount of time",
			Aliases: []string{"p"},
			Subcommands: []cli.Command{
				{
					Name:    "duration",
					Usage:   `Intervals between every quote in the form of a duration string (ex: "5m" or "1h30m10s")`,
					Aliases: []string{"d", "duration"},
					Action: func(c *cli.Context) error {

						//os.Args[3] is the interval of time between quotes.
						if len(os.Args) < 4 {
							fmt.Println("Incorrect use of the persistent commnad," +
								" type 'ancestorquotes persistent help' for more information")
							return nil
						}

						duration, err := time.ParseDuration(os.Args[3])
						if err != nil {
							fmt.Println("Incorrect use of the persistent commnad," +
								" type 'ancestorquotes persistent help' for more information")
							return nil
						}

						commands.Persistent(duration)

						return nil
					},
				},
				{
					Name:    "minute",
					Usage:   "Intervals in minutes between every quote",
					Aliases: []string{"m", "minutes"},
					Action: func(c *cli.Context) error {

						//os.Args[3] is the interval of time between quotes.
						if len(os.Args) < 4 {
							fmt.Println("Incorrect use of the persistent commnad," +
								" type 'ancestorquotes persistent help' for more information")
							return nil
						}

						timer, err := strconv.Atoi(os.Args[3])
						if err != nil {
							fmt.Println("Incorrect use of the persistent commnad," +
								" type 'ancestorquotes persistent help' for more information")
							return nil
						}

						// When using strconv on a non numeral string it gets converted 0
						if timer == 0 {
							fmt.Println("Incorrect use of the persistent commnad," +
								" type 'ancestorquotes persistent help' for more information")
							return nil
						}

						commands.Persistent(time.Duration(timer) * time.Minute)

						return nil
					},
				},
				{
					Name:    "second",
					Usage:   "Intervals in seconds between every quote",
					Aliases: []string{"s", "seconds"},
					Action: func(c *cli.Context) error {

						//os.Args[3] is the interval of time between quotes.
						if len(os.Args) < 4 {
							fmt.Println("Incorrect use of the persistent commnad," +
								" type 'ancestorquotes persistent help' for more information")
							return nil
						}

						timer, err := strconv.Atoi(os.Args[3])
						if err != nil {
							fmt.Println("Incorrect use of the persistent commnad," +
								" type 'ancestorquotes persistent help' for more information")
							return nil
						}

						// When using strconv on a non numeral string it gets converted 0
						if timer == 0 {
							fmt.Println("Incorrect use of the persistent commnad," +
								" type 'ancestorquotes persistent help' for more information")
							return nil
						}

						commands.Persistent(time.Duration(timer) * time.Second)

						return nil
					},
				},
			},
		},
		{
			Name:    "all",
			Usage:   "Shows all quotes the Ancestor has to offer",
			Aliases: []string{"a"},
			Action: func(c *cli.Context) error {
				commands.AllQuotes()
				return nil
			},
		},
		{
			Name:    "chat",
			Usage:   "The Ancestor talks with himself in a maddening fashion",
			Aliases: []string{"c"},
			Action: func(c *cli.Context) error {
				commands.Chat()
				return nil
			},
		},
		{
			Name:    "talkback",
			Usage:   "You can talk to the Ancestor and the Ancestor replies back in a crazy manner",
			Aliases: []string{"t"},
			Action: func(c *cli.Context) error {
				commands.TalkBack()
				return nil
			},
		},
		{
			Name:    "search",
			Usage:   "Searches all quotes the Ancestor has ever said with the word searched in them",
			Aliases: []string{"s"},
			Action: func(c *cli.Context) error {
				commands.Search(c.Args().First())
				return nil
			},
		},
		{
			Name:    "number",
			Usage:   "Shows a given number of random quotes the Ancestor has to offer",
			Aliases: []string{"n", "numbers"},
			Action: func(c *cli.Context) error {
				if len(os.Args) != 3 {
					fmt.Println("Incorrect use of the number command, example of correct usage: ancestorquotes number 3")
					return nil
				}
				number, err := strconv.Atoi(os.Args[2])
				if err != nil {
					fmt.Println("input parameter must be integer, example of correct usage: ancestorquotes number 3")
					return nil
				}
				if number < 1 {
					fmt.Println("input parameter must be greater than zero, example of correct usage: ancestorquotes number 3")
					return nil
				}
				commands.NumberOfQuotes(number)
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {

		if len(c.Args()) > 0 {
			fmt.Printf("%v is not a valid command.\n"+
				"Enter ancestorquotes --help to see the list of valid commands.\n", c.Args()[0])
		} else {
			fmt.Println(commands.RandomQuote())
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
