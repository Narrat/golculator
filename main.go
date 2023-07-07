package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golculator/compute"

	"github.com/containerd/console"
	"golang.org/x/term"
)

// Stores the state of the terminal before making it raw
var regularState console.Console

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

	// Get capabilities of current console
	regularState = console.Current()

	if err := regularState.SetRaw(); err != nil {
		panic(err)
	}

	terminal := term.NewTerminal(regularState, "> ")
	//terminal.AutoCompleteCallback = handleKey
	for {
		text, err := terminal.ReadLine()
		if err != nil {
			if err == io.EOF {
				// Quit without error on Ctrl^D on empty line or Ctrl^C
				exit()
			}
			panic(err)
		}

		text = strings.Replace(text, " ", "", -1)
		if text == "exit" || text == "quit" {
			terminal.Write([]byte(fmt.Sprintf("Quitting")))
			exit()
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
	// Is without function; io.EOF did catch ^D & ^C
	if key == '\x03' {
		// Quit without error on Ctrl^C
		exit()
	}
	return "", 0, false
}

func exit() {
	regularState.Reset()
	fmt.Println()
	os.Exit(0)
}
