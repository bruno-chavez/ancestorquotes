[![GoDoc](https://godoc.org/github.com/bruno-chavez/ancestorquotes?status.svg)](https://godoc.org/github.com/bruno-chavez/ancestorquotes)
[![Go Report Card](https://goreportcard.com/badge/github.com/bruno-chavez/ancestorquotes)](https://goreportcard.com/report/github.com/bruno-chavez/ancestorquotes)

`ancestorquotes` is a fun, little command-line app written in Go,
which prints a random quote used by the Ancestor or also known as
the Narrator from the Darkest Dungeon videogame.

## Installation

### Linux

#### From source code:

Requires Go to be installed on your machine. You can install Go from
https://golang.org/doc/install

Once installed, and with a correctly configured GOPATH, on a terminal type:

```
$ go get github.com/bruno-chavez/ancestorquotes
```

Then go to:

```
$GOPATH/src/github.com/bruno-chavez/ancetorquotes
```

And last, on a terminal type:

```
$ go install
```

## Updating

To update `ancestorquotes` simply type:

```
$ go get -u github.com/bruno-chavez/ancestorquotes
```

You can always check check your current version by typing:

```
$ ancestorquotes --version
ancestorquotes version 0.5
```

## Usage

Simply type `ancestorquotes`

```
$ ancestorquotes

Remind yourself that overconfidence is a slow and insidious killer.
```

## Commands

help: Shows general info about `ancestorquotes`.

```
$ ancestorquotes help
NAME:
   ancestorquotes - brings quotes from the darkest of dungeons!

USAGE:
   ancestorquotes [global options] command [command options] [arguments...]

VERSION:
   0.5

AUTHOR:
   bruno-chavez

COMMANDS:
     persistent, p  makes the app execute itself on a regular basis
     all, a         Prints all quotes from the darkest of dungeons!
     chat, c        The Ancestor talks with himself in a maddenly fashion
     talkback, t    The user talks to the Ancestor and the Ancestor replies back in a crazy manner
     help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --search value, -s value  search term that the quote must contain
   --help, -h     show help
   --version, -v  print the version
```

persistent: allows `ancestorquotes` to run every "x" amount of time.

```
$ ancestorquotes persistent
NAME:
   ancestorquotes persistent - schedules the app to run every x amount of time

USAGE:
   ancestorquotes persistent command [command options] [arguments...]

COMMANDS:
     minutes, m  minutes between every quote, value accepted if between 1 and 59
     seconds, s  seconds between every quote, value accepted if between 1 and 59

OPTIONS:
   --help, -h  show help
```

all: Displays all the quotes

```
$ ancestorquotes a -h
NAME:
   ancestorquotes all - Prints all quotes from the darkest of dungeons!

USAGE:
   ancestorquotes all [arguments...]
```


chat: prints a random quote that ends in a "?" followed by another one that finishes in a "."

```
$ ancestorquotes c
How many rats will it take to gnaw through a tonne of putrid flesh?
Ringing ears, blurred vision - the end approaches...

```

talkback: The user talks to the Ancestor and the Ancestor replies back in a crazy manner <br>
Enter your name first. Keep on chatting with the Ancestor! Enter `stop` to end the chat.

```
$ ancestorquotes t
Enter your name:
user
Hi user
What do you wanna say?
What a wonderful day!
Ancestor says: To those with the keen eye, gold gleams like a dagger's point.
What do you wanna say?
stop
GoodBye user
Bear in mind my last quote
Ancestor says: Perched at the very precipice of oblivion...


NAME:
   ancestorquotes talkback - User talks to the Ancestor and the Ancestor replies back in a crazy manner

USAGE:
   ancestorquotes talkback
```

### Sub command

For `persistant` to work, you must use one of the following subcommands:

minutes: allows `persistant` run evey "x" minutes, where "x" is a number between 1 and 59.

```
$ ancestorquotes persistent minutes 1
```

seconds: allows `persistant` run evey "x" seconds, where "x" is a number between 1 and 59.

```
$ ancestorquotes persistent seconds 30
```

You can type `stop` at any time during execution to stop the program.

```
$ ancestorquotes persistent second 1
Word is travelling. Ambition is stirring in distant cities. We can use this.
Perched at the very precipice of oblivion...
An increasing stockpile of curious trinkets, gathered from forbidden places.
stop
$
```

## Notes

This is a pretty small and niche project, created mainly to have fun,
so do that!

Only tested on Linux.

Heavely inspired by [motivate](https://github.com/mubaris/motivate) and
[this reddit bot](https://www.reddit.com/r/darkestdungeon/comments/7877vx/darkest_dungeon_ancestor_quote_bot/)

## Contribute

Found an bug or an error? Post it in the [issue tracker](https://github.com/bruno-chavez/ancestorquotes/issues).

Want to add an awesome new feature? [Fork](https://github.com/bruno-chavez/ancestorquotes/fork) this repository and add your feature, then send a pull request.

## License
The MIT License (MIT)
Copyright (c) 2017 Bruno Chavez
