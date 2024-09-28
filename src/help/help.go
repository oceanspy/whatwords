package help

import (
	"whatwords/src/color"
	"whatwords/src/message"
)

func Print() {
	message.Info("Usage: ee <command>")
	message.Title("    Tools:")
	message.Text("    ", color.Get("red"), "psw ", color.Get("magenta"), "", color.Reset, "                                              ", "Generate password")
}

func CommandNotFound() {
	message.Error("Command not found")
}

func CommandNotFoundAndHelp() {
	message.Error("Command not found")
	message.Info("whatwords -h/--help to show the help")
}
