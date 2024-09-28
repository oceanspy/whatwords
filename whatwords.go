package main

import (
	"os"
	"whatwords/src/help"
	"whatwords/src/message"
)

func main() {
	if len(os.Args) == 1 {
		help.CommandNotFoundAndHelp()
		os.Exit(0)
	}

	var wordList string
	var filePath string
	var url string
	var action string

	for i, e := range os.Args {
		if i == 0 {
			continue
		}

		switch e {
		case "-w":
			wordList = os.Args[i+1]
			continue
		case "-f":
			filePath = os.Args[i+1]
			action = "file"
			continue
		case "-u":
			url = os.Args[i+1]
			action = "url"
			continue
		default:
		}
	}

	switch action {
	case "file":
		message.Info("Parsing file")
	case "url":
		message.Info("Parsing URL")
	default:
		help.CommandNotFoundAndHelp()
	}
}
