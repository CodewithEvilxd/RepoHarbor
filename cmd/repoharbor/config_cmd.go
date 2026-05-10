package main

import (
	"fmt"

	"github.com/codewithevilxd/repoharbor/internal/config"
)

type configCommand struct {
	Args []string
}

func parseConfigArgs(args []string) (configCommand, error) {
	if len(args) == 0 {
		return configCommand{}, fmt.Errorf("usage: repoharbor config [list|set|unset|profile|favorite|recent]")
	}
	return configCommand{Args: args}, nil
}

func handleConfig(command configCommand) error {
	if len(command.Args) == 0 {
		return fmt.Errorf("usage: repoharbor config [list|set|unset|profile|favorite|recent]")
	}

	switch command.Args[0] {
	case "list":
		cfg, err := config.Load()
		if err != nil {
			return err
		}
		printConfig(cfg)
		return nil
	default:
		return fmt.Errorf("usage: repoharbor config [list|set|unset|profile|favorite|recent]")
	}
}

func printConfig(cfg config.Config) {
	fmt.Println("--- RepoHarbor config ---")
	fmt.Printf("Active profile: %s\n", cfg.ActiveProfileName())
}
