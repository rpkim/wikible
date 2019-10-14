package command

import (
	"strings"
	"fmt"
)

// PlanCommand is a plan the wiki tree from the text
type PlanCommand struct {
	Meta
}

func (c *PlanCommand) Run(args []string) int {
	fmt.Printf("Plan Run")

	return 1
}

func (c *PlanCommand) Help() string {
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
func (c *PlanCommand) Synopsis() string {
	return "Planning the wiki tree creation"
}

