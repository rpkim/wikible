package main 

import (
	"os"
	"fmt"

	"./command"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all the available Wikible commands.
var Commands map[string]cli.CommandFactory

// CommandMeta is the Meta to use for the commands. This must be written
// before the CLI is started.

func init() {
	fmt.Printf("init Commands")

	meta := command.Meta{
		Version: "0.0.1",
	}
	
	c := cli.NewCLI("wikible", "0.0.1")
	c.Args = os.Args[1:]
	Commands = map[string]cli.CommandFactory{
		"help": func() (cli.Command, error) {
			return &command.HelpCommand{
				Meta: meta,
			}, nil
		},
		"login": func() (cli.Command, error) {
			return &command.LoginCommand{
				Meta: meta,
			}, nil
		},
		"plan": func() (cli.Command, error) {
			return &command.PlanCommand{
				Meta: meta,
			}, nil
		},
		"apply": func() (cli.Command, error) {
			return &command.ApplyCommand{
				Meta: meta,
			}, nil
		},
	}
}