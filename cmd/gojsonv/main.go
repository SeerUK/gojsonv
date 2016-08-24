package main

import (
	"os"

	"github.com/eidolon/gonsole"
	"github.com/eidolon/wordwrap"
	"github.com/SeerUK/gojsonv/command"
)

func main() {
	app := gonsole.NewApplication("gojsonv", "0.0.1")

	wrapper := wordwrap.Wrapper(78, true)

	app.Help = wrapper(`
		Go-based CLI JSON schema validator, with support for local and remote schema, or strings.
	`)

	app.AddCommands([]gonsole.Command{
		command.ValidateCommand(),
	})

	app.Run(os.Args[1:])
}
