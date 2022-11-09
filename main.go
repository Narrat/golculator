package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/alfredxing/calc/compute"
	"golang.org/x/term"
)

// Stores the state of the terminal before making it raw
var regularState *term.State
var fd_term int

func main() {
	if len(os.Args) > 1 {
		input := strings.Replace(strings.Join(os.Args[1:], ""), " ", "", -1)
		res, err := compute.Evaluate(input)
		if err != nil {
			return
		}
		fmt.Printf("%s\n", strconv.FormatFloat(res, 'G', -1, 64))
		return
	}

	var err error
	fd_term = int(os.Stdin.Fd())
	regularState, err = term.MakeRaw(fd_term)
	if err != nil {
		panic(err)
	}
	defer term.Restore(fd_term, regularState)

	screen := struct {
		io.Reader
		io.Writer
	}{os.Stdin, os.Stdout}
	terminal := term.NewTerminal(screen, "> ")
	terminal.AutoCompleteCallback = handleKey
	for {
		text, err := terminal.ReadLine()
		if err != nil {
			if err == io.EOF {
				// Quit without error on Ctrl^D
				exit()
			}
			panic(err)
		}

		text = strings.Replace(text, " ", "", -1)
		if text == "exit" || text == "quit" {
			break
		}

		res, err := compute.Evaluate(text)
		if err != nil {
			terminal.Write([]byte(fmt.Sprintln("Error: " + err.Error())))
			continue
		}
		terminal.Write([]byte(fmt.Sprintln(strconv.FormatFloat(res, 'G', -1, 64))))
	}
}

func handleKey(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
	if key == '\x03' {
		// Quit without error on Ctrl^C
		exit()
	}
	return "", 0, false
}

func exit() {
	term.Restore(fd_term, regularState)
	fmt.Println()
	os.Exit(0)
}
