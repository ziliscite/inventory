package commands

import (
	"fmt"
	"inventory/config"
)

func removeCommand(i *config.Inventory, p []string) error {
	if len(p) == 0 {
		return fmt.Errorf("no name identifier is given")
	}

	// whatever, all we ever needed is the name anyway
	name := p[0]
	it := config.Item{
		Name: name,
	}

	iErr := i.Remove(it)
	if iErr != nil {
		return iErr
	}

	return nil
}

var rc = Command{
	name:        "remove",
	description: "remove an new item from the inventory",
	example:     "remove [item's name]",
	command:     removeCommand,
}
