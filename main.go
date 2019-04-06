package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bruno-chavez/ancestorquotes/commands"
	"github.com/urfave/cli"
)

func main() {

	app := cli.NewApp()
	app.Name = "ancestorquotes"
	app.Author = "bruno-chavez"
	app.Usage = "Brings quotes from the darkest of dungeons!"
	app.Version = "2.0.0"
	app.Commands = []cli.Command{
		{
			Name:    "persistent",
			Usage:   "Makes the Ancestor say a quote every certain amount of time",
			Aliases: []string{"p"},
			Subcommands: []cli.Command{
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
						timer, _ := strconv.Atoi(os.Args[3])
						// When using strconv on a non numeral string it gets converted 0
						if timer == 0 {
							fmt.Println("Incorrect use of the persistent commnad," +
								" type 'ancestorquotes persistent help' for more information")
							return nil
						}
						commands.Persistent(timer, "minute")

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
						timer, _ := strconv.Atoi(os.Args[3])
						// When using strconv on a non numeral string it gets converted 0
						if timer == 0 {
							fmt.Println("Incorrect use of the persistent commnad," +
								" type 'ancestorquotes persistent help' for more information")
							return nil
						}
						commands.Persistent(timer, "second")

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
