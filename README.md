[![GoDoc](https://godoc.org/github.com/bruno-chavez/ancestorquotes?status.svg)](https://godoc.org/github.com/bruno-chavez/ancestorquotes)
[![Go Report Card](https://goreportcard.com/badge/github.com/bruno-chavez/ancestorquotes)](https://goreportcard.com/report/github.com/bruno-chavez/ancestorquotes)
[![Build Status](https://travis-ci.org/bruno-chavez/ancestorquotes.svg?branch=master)](https://travis-ci.org/bruno-chavez/ancestorquotes)

`ancestorquotes` is a fun, little command-line app written in Go,
shows quotes used by the Ancestor,
the narrator of the videogame Darkest Dungeon.

## Download & Installation

### Executables:

`ancestorquotes` supports Linux, Mac and Windows! you can 
download it from 
[here](https://github.com/bruno-chavez/ancestorquotes/releases), 
simply click on the file that has your OS and architecture 
on its name.

After it is downloaded, run the executable to start using 
`ancestorquotes`!

### From source code:

Requires Go to be installed on your machine. 
You can install Go [here](https://golang.org/doc/install).

Once installed, on a terminal type:

```
$ git clone https://github.com/bruno-chavez/ancestorquotes
```

Enter the directory:

```
$ cd ancestorquotes
```

Install dependencies:

```
$ go list
```

Install `ancestorquotes`:

```
$ go install
```

## Usage

Type `ancestorquotes` and a random quote from the Ancestor will be displayed:

```
$ ancestorquotes

Remind yourself that overconfidence is a slow and insidious killer.
```

## Commands

#### Help:
Shows general info about `ancestorquotes`.

```
$ ancestorquotes help
NAME:
   ancestorquotes - Brings quotes from the darkest of dungeons!

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:

   2.0.0

AUTHOR:
   bruno-chavez

COMMANDS:
     persistent, p  Makes the Ancestor say a quote every certain amount of time
     all, a         Shows all quotes the Ancestor has to offer
     chat, c        The Ancestor talks with himself in a maddening fashion
     talkback, t    You can talk to the Ancestor and the Ancestor replies back in a crazy manner
     search, s      Searches all quotes the Ancestor has ever said with the word searched in them
     help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

#### Persistent:

Makes `ancestorquotes` run every certain amount of time.

```
$ ancestorquotes persistent
NAME:
   ancestorquotes persistent - Makes the Ancestor say a quote every certain amount of time

USAGE:
   ancestorquotes persistent [global options] command [command options] [arguments...]

VERSION:
   1.2

COMMANDS:
     minute, m, minutes  Intervals in minutes between every quote
     second, s, seconds  Intervals in seconds between every quote

GLOBAL OPTIONS:
   --help, -h  show help
```

#### All:

Shows all quotes the Ancestor has to offer:

```
$ ancestorquotes all
There is a place, beneath those ancient ruins, in the moor, that calls out to the boldest among them...

'We are the Flame!', they cry, 'And Darkness fears us!'
.
.
.
Until the stars align in their inexorable formation and what sleeps is aroused once more. To hatch from this fragile shell of earth and rock, and bring our inescapable end.

So, seek solace in a manner befitting your lineage, and take up your nugatory vigil, haunted forever by that sickening prose, echoing through the infinite blackness of space and time

```


#### Chat:

Prints a random quote that ends in a "?" followed by another one that finishes in a "."

```
$ ancestorquotes chat
How many rats will it take to gnaw through a tonne of putrid flesh?
Ringing ears, blurred vision - the end approaches...

```

#### Talkback:

The user talks to the Ancestor and the Ancestor replies back in a crazy manner <br>
Enter your name first. Keep on chatting with the Ancestor! Enter `stop` to end the chat.

```
$ ancestorquotes talkback
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

#### Search:

Searches for all the quotes with the input word on it:
```
$ ancestorquotes search glory
More arrive foolishly seeking fortune and glory in this domain... Of the damned.

Where there is no peril in the task, there can be no glory in its accomplishment.

The bigger the beast, the greater the glory.

Another life wasted in the pursuit of glory and gold.
```

#### Number:

Shows a fixed number of random quotes the Ancestor has to offer:

```
$ ancestorquotes number 3
Their cursed champion falls!

The bigger the beast, the greater the glory.

There will be no sleep tonight - the wild shouts and frantic drumming will see to that.

```

### Sub-commands

#### Persistent:
For `persistent` to work, you must use one of the following subcommands:

minutes: allows `persistent` run evey "x" minutes, where "x" is a number between 1 and 59.

```
$ ancestorquotes persistent minutes 1
```

seconds: allows `persistent` run evey "x" seconds, where "x" is a number between 1 and 59.

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

It has not been tested on Mac or Windows. If you run into any problems while trying to install or use it, feel free to create an [issue](https://github.com/bruno-chavez/ancestorquotes/issues).

Heavely inspired by [motivate](https://github.com/mubaris/motivate) and
[this reddit bot](https://www.reddit.com/r/darkestdungeon/comments/7877vx/darkest_dungeon_ancestor_quote_bot/).

Sister project of [restedancestor](https://github.com/bruno-chavez/restedancestor).

Check the amazing game this app was inspired in, [Darkest Dungeon](https://store.steampowered.com/app/262060/Darkest_Dungeon/) available in Steam for every platform!

## Contribute

Found a bug or an error? Post it in the [issue tracker](https://github.com/bruno-chavez/ancestorquotes/issues).

Want to add an awesome new feature? [Fork](https://github.com/bruno-chavez/ancestorquotes/fork) this repository and add your feature, then send a pull request.

## License
The MIT License (MIT)
Copyright (c) 2019 Bruno Chavez
