package commands

import (
	"fmt"
	"inventory/config"
)

func helpCommand(i *config.Inventory, params []string) error {
	fmt.Println("If you have a parameter that is longer than 1 word, use \"[param]\" ")
	for _, c := range Commands {
		fmt.Println("-", c.name, "\t |", c.description)
	}
	return nil
}

var hc = Command{
	name:        "help",
	description: "displays all the commands available",
	command:     helpCommand,
}
