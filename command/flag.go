package command

import (
	"flag"
	"fmt"
	"os"
	"wikible/config"
)

func parseFlag(cmdName string, args []string, cfg *config.Config) {

	cmd := flag.NewFlagSet(cmdName, flag.ExitOnError)

	if cmdName == "plan" {
		cmd.StringVar(&cfg.StructedTextFile, "p", "", "Project Template File Path")
		cmd.StringVar(&cfg.WikiAddress, "a", "", "Address of Wiki")
	} else if cmdName == "apply" {
		cmd.StringVar(&cfg.StructedTextFile, "p", "", "Project Template File Path")
		cmd.StringVar(&cfg.PageID, "i", "", "Page ID(parent ID)")
		cmd.StringVar(&cfg.WikiAddress, "a", "", "Address of Wiki")
	}

	cmd.Parse(args)

	if cfg.StructedTextFile == "" {
		fmt.Println("-p Project Template File Path is not setup")
		os.Exit(-1)
	}

}
