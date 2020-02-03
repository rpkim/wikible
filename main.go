package main

import (
	"github.com/mitchellh/cli"
	"log"
	"os"
	"runtime"
	"wikible/command"
)

const APP_NAME = "wikible"
const APP_VERSION = "0.0.1"

func main() {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	//code from https://github.com/hashicorp/terraform/blob/master/main.go#L40
	os.Exit(realMain())
}

func realMain() int {
	// CLI Creation
	c := cli.NewCLI(APP_NAME, APP_VERSION)
	c.Args = os.Args[1:]
	c.HelpWriter = os.Stdout

	// Commands Registration
	c.Commands = map[string]cli.CommandFactory{
		"plan": func() (cli.Command, error) {
			return &command.PlanCommand{}, nil
		},
		"apply": func() (cli.Command, error) {
			return &command.ApplyCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Fatalln(err)
	}

	os.Exit(exitStatus)

	return 0
}

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}
