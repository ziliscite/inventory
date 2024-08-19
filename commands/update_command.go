package commands

import "inventory/config"

func updateCommand(i *config.Inventory, p []string) error {
	return nil
}

type Attribute interface {
	~string | ~int
}

func updateAttribute[T Attribute](i *config.Inventory, a T, n T) error {
	return nil
}

var uc = Command{
	name:        "increase",
	description: "increase the quantity of an item (in the inventory) by n-input amount\n\tdefault increment value is 1 if not specified\n\texample: increase [item's name] [amount]",
	command:     updateCommand,
}
