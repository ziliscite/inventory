package commands

import "inventory/config"

func listCommand(i *config.Inventory, p []string) error {
	err := i.Display()
	if err != nil {
		return err
	}

	return nil
}

var lc = Command{
	name:        "list",
	description: "list every items in the inventory",
	command:     listCommand,
}
