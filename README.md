# New York Times Game Solvers

[![Go](https://github.com/kavfixnel/NYTimes-games/actions/workflows/go.yaml/badge.svg)](https://github.com/kavfixnel/NYTimes-games/actions/workflows/go.yaml)

Collection of solvers for the NY Times games.

## How to use

All solvers can be run with the Go cli. each game has it's own subcommand with arguments
corresponding to the game input.

```bash
~/nytimes-games ❯ go run main.go                  
nytimes-solver is a small commandlet that allows you to input
game parameters and present you with possible solutions.

Usage:
  nytimes-solver [command]

Available Commands:
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  spelling-bee Gives solutions to the spelling-bee game
  wordle       Gives solutions to the world game

Flags:
  -h, --help   help for nytimes-solver

Use "nytimes-solver [command] --help" for more information about a command.
```

## Spelling Bee

The offical game can be found [here](https://www.nytimes.com/puzzles/spelling-bee)

The solver requires 2 flag arguments that denot the input of the game (`--extra`/`-e` and `--required`/`-r`)
The value of the flags are a string of characters denoting the required character(s) (denoted by the yellow
hexagon(s)), as well any extra characters (denoted by the gray squares)

An example input could look something like the following:
```bash
~/nytimes-games ❯ go run main.go sp -r f -e etycap
affa
affy
atef
face
fact
facy
faff
fate
feat
taft
teff
yaff
Taffy
aface
caffa
facet
facty
faffy
fatty
featy
taffy
affect
efface
effect
effete
facete
patefy
tepefy
catface
taffeta
taffety
caffeate
affectate
```

The solver will then spit out a list of possible words sorted by points and alphabetically (when applicable)

> Note that the accuracy of the solver depends heavily on which words are accepted by the game
> as well as the words list on the system this is being run on (see https://github.com/kavfixnel/words)