`ancestorquotes` is a fun, little command-line tool written in go,
 prints a random quote used by the Ancestor or also known as
the Narrator from the Darkest Dungeon videogame.

## Installation


### Linux

#### From source code

Requires Go to be installed on your machine. You can install Go from
https://golang.org/doc/install

Once installed, and with a correctly configured GOPATH type:

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

## Sub commands

help: Shows general info about the proyect.

```
$ ancestorquotes help
NAME:
   ancestorquotes - brings quotes from the darkest of dungeons!

USAGE:
   ancestorquotes [global options] command [command options] [arguments...]

VERSION:
   0.1

AUTHOR:
   bruno-chavez

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version


```

## Notes

This is a pretty small and niche project, created mainly to have fun, so do that!

Heavely inspired by [motivate](https://github.com/mubaris/motivate).
