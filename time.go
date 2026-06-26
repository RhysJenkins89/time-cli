package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		router(os.Args[1])
	} else {
		fmt.Println("Please input a command.")
	}
}

func router(command string) {
	switch command {
	case "now":
		handleNow()
	case "help":
		handleHelp()
	default:
		handleUnrecognisedCommand()
	}
}

func handleNow() {
	fmt.Println("This is the now command function.")
}

func handleHelp() {
	fmt.Println("This is the help command function.")
}

func handleUnrecognisedCommand() {
	fmt.Println("Please input a recognised command.")
}
