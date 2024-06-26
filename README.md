## golculator

[![GoDoc](https://godoc.org/github.com/alfredxing/calc?status.svg)](https://godoc.org/github.com/alfredxing/calc) [![Build Status](https://travis-ci.org/alfredxing/calc.svg?branch=master)](https://travis-ci.org/alfredxing/calc)

A simple, fast, and intuitive command-line calculator written in Go.  
Fork of [calc by alfredxing](https://github.com/alfredxing/calc).
Reason for the fork, because the original repo is a puplic archive and the name was to conflicting.

### Install
The original can be installed as any other Go program (with Go 1.17+):
```
go install github.com/alfredxing/calc@latest
```
In the future the same may work for golculator, if I decide to add the hostname to the `go.mod` .   
Instead you could clone this repo and and run either `go build` or `go install` in the source directory.

### Usage
You can use golculator in two ways: shell mode and command.

#### Shell mode
This is probably the mode you'll want to use. It's like the `python` shell or `irb`.  
The shell mode uses the combination of [`containerd/console`](https://pkg.go.dev/github.com/containerd/console) and [`term`](https://pkg.go.dev/golang.org/x/term).
Whereas `term` provides the new shell and support for some console features (like history, pasting, and the `exit` command). But out of the box it only works good enough on `unix`/`linux`-based consoles. For a better cross-platform support is some more work necessary.  
This is were `containderd/console` comes into play as it allows to recognize the features of the current console and sets them active. Greatly enhancing the featureset on `Windows`.
```shell
> 1+1
2
> 3(5/(3-4))
-15
> 3*pi^2
29.608813203268074
> @+1
30.608813203268074
> @@@*2
-30
> ln(-1)
NaN
```

#### Command
You can also use golculator to evaluate an expression with just a single command (i.e. without opening the shell). To do this, just use `golculator [expression]`:
```shell
bash$ golculator 1+1
2
bash$
```

### Supported functions, operators, and constants
golculator supports all the standard stuff, and I'm definitely adding more later (also feel free to fork and add your own!)

##### Operators
`+`, `-`, `*`, `/`, `^`, `%`

##### Functions
`sin`, `cos`, `tan`, `cot`, `sec`, `csc`, `asin`, `acos`, `atan`, `acot`, `asec`, `acsc`, `sqrt`, `log`, `lg`, `ln`, `abs`

##### Constants
`e`, `pi`, `π`

##### History
Previous results can be accessed with the `@` symbol. A single `@` returns the result of the last computation, while multiple `@` gets the n<sup>th</sup> last result, where n is the number of `@`s used (for example, `@@` returns the second-last result, `@@@@@` returns the fifth-last result).

### Why not use ...?
- Google
  - Doesn't work without an internet connection
  - Slower
  - Doesn't show previous computations, so you end up with multiple tabs open at once.
- Spotlight (on OS X)
  - No history
  - Switching between Spotlight and other windows isn't too fun
- Python/IRB
  - Requires use of a separate math module for most functions and constants
  - A little bit slower to start up
- `bc`
  - Limited number of built-in functions; these have shortened (not too intuitive) names as well.

The alternatives above are all great, and have their own advantages over calc/golculator. I highly recommend looking into these if you don't like how calc/golculator works.
