package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		router(os.Args[1]) // Ignore the first argument and accept the next X number of arguments. Ignore any after that.
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
	// In order to get the user's location, use the client IP address
	req, err := http.NewRequest(http.MethodGet, "http://api.open-notify.org/astros.json", nil)
	if err != nil {
		fmt.Println("This is the error block.")
		log.Fatal(err)
	}
	fmt.Println(req)
}

func handleHelp() {
	fmt.Println("This is the help command function.")
}

func handleUnrecognisedCommand() {
	fmt.Println("Please input a recognised command.")
}
