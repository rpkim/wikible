package command

import (
	"strings"
	"fmt"
)

// ApplyCommand is for applying the wiki creation
type ApplyCommand struct {
	Meta
}

func (c *ApplyCommand) Run(args []string) int {
	fmt.Printf("Apply Run")

	return 1
}

func (c *ApplyCommand) Help() string {
	helpText := `
Usage :
	wikible apply [options] 
	
	Apply the wiki tree creation and numbering the tree.

Options:
   -p   plan text file
   -s	tructured text file
   -v	variable file
`
	return strings.TrimSpace(helpText)
}
func (c *ApplyCommand) Synopsis() string {
	return "Planning the wiki tree creation"
}

