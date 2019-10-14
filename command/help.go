package command

import (
	"strings"
	"fmt"
)

// HelpCommand is for listing the command
type HelpCommand struct {
	Meta
}

//Run Command
func (c *HelpCommand) Run(args []string) int {
	fmt.Printf("Help Run")

	return 1
}

//Hep Command
func (c *HelpCommand) Help() string {
	helpText := `
Usage :
	wikible plan [options] 
	
	Planning the wiki tree creation and numbering the tree.

Options:
   -s	tructured text file
   -v	variable file
`
	return strings.TrimSpace(helpText)
}
func (c *HelpCommand) Synopsis() string {
	return "Planning the wiki tree creation"
}

