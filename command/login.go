package command

import (
	"strings"
	"fmt"
)

// LoginCommand is for login the wiki api
type LoginCommand struct {
	Meta
}

func (c *LoginCommand) Run(args []string) int {
	fmt.Printf("Login Run")

	return 1
}

func (c *LoginCommand) Help() string {
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
func (c *LoginCommand) Synopsis() string {
	return "Planning the wiki tree creation"
}

