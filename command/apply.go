package command

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"strings"
	base64 "encoding/base64"

	"wikible/config"
	w "wikible/wiki"
	u "wikible/util"
)

// ApplyCommand is for applying the wiki creation
type ApplyCommand struct {
}

// Run Command for Apply.
// Create the wiki tree
func (c *ApplyCommand) Run(args []string) int {
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

//RunContext is for ApplyCommand
func (c *ApplyCommand) RunContext(buildCtx context.Context, args []string) int {
	var cfg config.Config

	//parse Plan Args
	parseFlag("apply", args, &cfg)

	//username password
	username, password := u.GetCredentials()
	cred := "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password))

	if cfg.PageID == "" {
		cfg.PageID = u.GetPromptString("Enter PageID:")
	}

	if cfg.WikiAddress == "" {
		cfg.WikiAddress = u.GetPromptString("Enter WikiAddress:")
	}

	if cfg.StructedTextFile == "" {
		log.Fatalln("No StructedTextfile")
		return 1
	}

	wiki, err := w.CreateWikiAPI(cfg.WikiAddress, cred)

	if err != nil {
		fmt.Println("CreateWikiAPI Error:")
		log.Fatalln(err)
		return 1
	}

	bindingVars := u.ParseProjectConfig(cfg.StructedTextFile)
	wikiNode := w.GenerateWikiNodeFromString(bindingVars)

	fmt.Println("------------Result of Plan------------")
	w.PrintWikiNode(wikiNode)
	fmt.Println("--------------------------------------")



	//Are you apply?
	if strings.ToLower(u.GetPromptString("Do you apply?(y/n)")) == "y" {
		pageContent, err := wiki.GetPageContent(cfg.PageID)

		if err != nil {
			fmt.Println("GetPageContent Error:")
			log.Fatalln(err)
			return 1
		}

		var wg sync.WaitGroup
		wg.Add(1)
		go w.CreateWikiNode(&wg, wiki, pageContent.Space, cfg.PageID, wikiNode)
		wg.Wait()

		fmt.Println("Apply Done !")
	}

	return 0
}

// Help Command is for Help
// How to use the Apply
func (c *ApplyCommand) Help() string {
	helpText := `
Usage :
	wikible apply [options]

	Apply the wiki tree creation and numbering the tree.

Options:
   -p   plan text file
   -s	tructured text file
   -v	variable file
   -i   parent id
`
	return strings.TrimSpace(helpText)
}

//Synopsis is the single line description of this command
func (c *ApplyCommand) Synopsis() string {
	return "Apply the wiki tree creation"
}

