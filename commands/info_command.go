package commands

import (
	"fmt"
	"inventory/config"
	"os"
)

func helpCommand(i *config.Inventory, p []string) error {
	fmt.Println("If you have a parameter that is longer than 1 word, use \"[param]\" ")
	for _, c := range Commands {
		fmt.Println("-", c.name, " - ", c.description)
	}
	return nil
}

var hc = Command{
	name:        "help",
	description: "displays all the commands available",
	command:     helpCommand,
}

func exitCommand(i *config.Inventory, p []string) error {
	os.Exit(0)
	return nil
}

var ec = Command{
	name:        "exit",
	description: "exit the program",
	command:     exitCommand,
}
