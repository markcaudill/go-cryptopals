package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/genuinetools/pkg/cli"
	"github.com/markcaudill/go-cryptopals/version"
)

var (
	challenge uint
	set       uint
)

func main() {
	p := cli.NewProgram()
	p.Name = "cryptopals"
	p.Description = "Go implementation of cryptopals solutions"
	p.GitCommit = version.GITCOMMIT
	p.Version = version.VERSION
	p.FlagSet = flag.NewFlagSet("global", flag.ExitOnError)
	p.FlagSet.UintVar(&challenge, "challenge", 1, "the challenge number")
	p.FlagSet.UintVar(&set, "set", 1, "the set number")
	p.Action = func(ctx context.Context, args []string) error {
		// On ^C, or SIGTERM handle exit.
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		signal.Notify(c, syscall.SIGTERM)
		go func() {
			for sig := range c {
				log.Printf("Received %s, exiting.", sig.String())
				os.Exit(0)
			}
		}()

		// Determine which challenge to run and run it

		return fmt.Errorf("challenge does not exist (s%d c%d)", challenge, set)
	}

	p.Run()
}
