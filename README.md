[![GoDoc](https://godoc.org/github.com/bruno-chavez/ancestorquotes?status.svg)](https://godoc.org/github.com/bruno-chavez/ancestorquotes)

`ancestorquotes` is a fun, little command-line app written in Go,
which prints a random quote used by the Ancestor or also known as
the Narrator from the Darkest Dungeon videogame.


## Installation


### Linux

#### From source code:

Requires Go to be installed on your machine. You can install Go from
https://golang.org/doc/install

Once installed, and with a correctly configured GOPATH, type:

```
$ go get -u github.com/bruno-chavez/ancestorquotes
```


Then go to:

```
GO_DIR/src/github.com/bruno-chavez/ancetorquotes
```

On a terminal type:

```
$ go install
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
   0.2

AUTHOR:
   bruno-chavez

COMMANDS:
     persistent, p  makes the app execute itself on a regular basis
     help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
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

### Sub command

minutes: allows `persistant` run evey "x" minutes, where "x" is a number between 1 and 59.

```
$ ancestorquotes persistent minutes 1
```


seconds: allows `persistant` run evey "x" seconds, where "x" is a number between 1 and 59.

```
$ ancestorquotes persistent seconds 30
```


## Notes

This is a pretty small and niche project, created mainly to have fun,
so do that!

Only tested on Linux.

Quotes are parsed from quotes.go which contains a json
top level array.

Complete list of quotes can be found
[here](https://darkestdungeon.gamepedia.com/Narrator).

Heavely inspired by [motivate](https://github.com/mubaris/motivate) and 
[this reddit bot](https://www.reddit.com/r/darkestdungeon/comments/7877vx/darkest_dungeon_ancestor_quote_bot/)
