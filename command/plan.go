package command

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"wikible/config"
	util "wikible/util"
	w "wikible/wiki"
)

// PlanCommand is a plan the wiki tree from the text
type PlanCommand struct {
}

//Run function is for excuting Plan
//Return the wiki tree with numbering
func (c *PlanCommand) Run(args []string) int {
	buildCtx, cancelBuildCtx := context.WithCancel(context.Background())
	// Handle interrupts for this build
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	defer func() {
		cancelBuildCtx()
		signal.Stop(sigCh)
		close(sigCh)
	}()
	go func() {
		select {
		case sig := <-sigCh:
			if sig == nil {
				// context got cancelled and this closed chan probably
				// triggered first
				return
			}
			cancelBuildCtx()
		case <-buildCtx.Done():
		}
	}()

	return c.RunContext(buildCtx, args)
}

//RunContext is for
func (c *PlanCommand) RunContext(buildCtx context.Context, args []string) int {
	var cfg config.Config

	//parse Plan Args
	parseFlag("plan", args, &cfg)

	bindingVars := util.ParseProjectConfig(cfg.StructedTextFile)
	fmt.Println("------------Result of Plan------------")
	w.PrintWikiNode(w.GenerateWikiNodeFromString(bindingVars))
	fmt.Println("--------------------------------------")

	return 0
}

//Help function is for usage of plan
func (c *PlanCommand) Help() string {
	helpText := `
Usage :
	wikible plan [options]

	Planning the wiki tree creation and numbering the tree.

Options:
   -s	structured text file
   -v	variable file
`
	return strings.TrimSpace(helpText)
}

//Synopsis is the single line description of this command
func (c *PlanCommand) Synopsis() string {
	return "Planning the wiki tree creation"
}
