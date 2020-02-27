package command

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"wikible/config"
)

// GetCommand is a Get the wiki tree from the text
type GetCommand struct {
}

//Run function is for excuting Get
//Return the wiki tree with numbering
func (c *GetCommand) Run(args []string) int {
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
func (c *GetCommand) RunContext(buildCtx context.Context, args []string) int {
	var cfg config.Config

	//parse Get Args
	parseFlag("Get", args, &cfg)


	return 0
}

//Help function is for usage of Get
func (c *GetCommand) Help() string {
	helpText := `
Usage :
	wikible get [options]

	Get the wiki pages to code.

Options:
	-i   parent id
	-a   wiki address
`
	return strings.TrimSpace(helpText)
}

//Synopsis is the single line description of this command
func (c *GetCommand) Synopsis() string {
	return "Getting the wiki tree to code"
}
