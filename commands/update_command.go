package commands

import "inventory/config"

func updateCommand(i *config.Inventory, p []string) error {
	err := updateAttribute(i, p[1], 2)
	if err != nil {
		return err
	}
	return nil
}

type Attribute interface {
	~string | ~int
}

func updateAttribute[T Attribute](i *config.Inventory, a string, n T) error {
	return nil
}

var uc = Command{
	name:        "update",
	description: "updates an item's attributes. e.g., name, quantity, and price",
	example:     "update [item's name] [attribute's name] [new attribute]",
	command:     updateCommand,
}
