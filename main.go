package main

import (
	"fmt"
	"github.com/bruno-chavez/ancestorquotes/commands"
	"github.com/bruno-chavez/ancestorquotes/quotes"
	"github.com/urfave/cli"
	"os"
	"strconv"
)

//Contains all the quotes in a QuoteSlice type slice.
//Used as a global variable to avoid having multiples of the same slice, reducing load times and memory usage.
var quoteSLice = quotes.Parse()

func main() {

	app := cli.NewApp()
	app.Name = "ancestorquotes"
	app.Author = "bruno-chavez"
	app.Usage = "Brings quotes from the darkest of dungeons!"
	app.Version = "1.2"
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
						} else {
							timer, _ := strconv.Atoi(os.Args[3])
							if timer == 0 {
								fmt.Println("Incorrect use of the persistent commnad," +
									" type 'ancestorquotes persistent help' for more information")
								return nil
							} else {
								commands.Persistent(quoteSLice, timer, "minute")
							}
						}
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
						} else {
							timer, _ := strconv.Atoi(os.Args[3])
							if timer == 0 {
								fmt.Println("Incorrect use of the persistent commnad," +
									" type 'ancestorquotes persistent help' for more information")
								return nil
							} else {
								commands.Persistent(quoteSLice, timer, "second")
							}
						}
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
				commands.AllQuotes(quoteSLice)
				return nil
			},
		},
		{
			Name:    "chat",
			Usage:   "The Ancestor talks with himself in a maddening fashion",
			Aliases: []string{"c"},
			Action: func(c *cli.Context) error {
				commands.Chat(quoteSLice)
				return nil
			},
		},
		{
			Name:    "talkback",
			Usage:   "You can talk to the Ancestor and the Ancestor replies back in a crazy manner",
			Aliases: []string{"t"},
			Action: func(c *cli.Context) error {
				commands.TalkBack(quoteSLice)
				return nil
			},
		},
		{
			Name:    "search",
			Usage:   "Searches all quotes the Ancestor has ever said with the word wanted to be searched on it",
			Aliases: []string{"s"},
			Action: func(c *cli.Context) error {
				commands.Search(quoteSLice, c.Args().First())
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(c.Args()) > 0 {
			fmt.Printf("%v is not a valid command.\nEnter ancestorquotes --help to see the list of valid commands.\n", c.Args()[0])
		} else {
			fmt.Println(quoteSLice.RandomQuote())
		}
		return nil
	}
	app.Run(os.Args)
	return
}
