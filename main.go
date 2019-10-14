package main

import (
  "fmt"
  "os"
  "log"

  "./version"
  "github.com/mitchellh/cli"
)

func init() {
	fmt.Printf("start")
}

func main() {
	log.SetPrefix("")
	os.Exit(realMain())
}


func realMain() int {	
	cli := &cli.CLI{
		Args:         os.Args[1:],
		Autocomplete: true,
		Commands:     Commands,
		HelpFunc:     excludeHelpFunc(Commands, []string{"plugin"}),
		HelpWriter:   os.Stdout,
		Name:         "wikible",
		Version:      version.Version,
	}

	exitStatus, err := cli.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)

	return 0
}

// excludeHelpFunc filters commands we don't want to show from the list of
// commands displayed in packer's help text.
func excludeHelpFunc(commands map[string]cli.CommandFactory, exclude []string) cli.HelpFunc {
	// Make search slice into a map so we can use use the `if found` idiom
	// instead of a nested loop.
	var excludes = make(map[string]interface{}, len(exclude))
	for _, item := range exclude {
		excludes[item] = nil
	}

	// Create filtered list of commands
	helpCommands := []string{}
	for command := range commands {
		if _, found := excludes[command]; !found {
			helpCommands = append(helpCommands, command)
		}
	}

	return cli.FilteredHelpFunc(helpCommands, cli.BasicHelpFunc("wikible"))
}