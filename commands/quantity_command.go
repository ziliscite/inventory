package commands

import (
	"fmt"
	"inventory/config"
	"strconv"
)

func validateParams(p []string) (string, int, error) {
	switch len(p) {
	case 0:
		return "", 0, fmt.Errorf("invalid command line arguments: require the item's name as an argument")
	case 1:
		p = append(p, "1")
	}

	name := p[0]
	quantity, aErr := strconv.Atoi(p[1])
	if aErr != nil {
		return "", 0, aErr
	}

	if quantity < 1 {
		return "", 0, fmt.Errorf("invalid command line arguments: quantity argument must be greater than zero")
	}

	return name, quantity, nil
}

func quantityMethod(i *config.Inventory, n string, q int) error {
	it := config.Item{
		Name: n,
	}

	err := i.Increment(it, q)
	if err != nil {
		return err
	}

	return nil
}

func incrementCommand(i *config.Inventory, p []string) error {
	name, quantity, err := validateParams(p)
	if err != nil {
		return err
	}

	err = quantityMethod(i, name, quantity)
	if err != nil {
		return err
	}

	return nil
}

func decrementCommand(i *config.Inventory, p []string) error {
	name, quantity, err := validateParams(p)
	if err != nil {
		return err
	}

	err = quantityMethod(i, name, -quantity)
	if err != nil {
		return err
	}

	return nil
}

var ic = Command{
	name:        "increase",
	description: "increase the quantity of an item in the inventory",
	example:     "increase [item's name] [quantity]",
	command:     incrementCommand,
}

var dc = Command{
	name:        "decrease",
	description: "decrease the quantity of an item in the inventory",
	example:     "decrease [item's name] [quantity]",
	command:     decrementCommand,
}
