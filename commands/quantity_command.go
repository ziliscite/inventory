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

func incrementMethod(i *config.Inventory, p []string) error {
	// in this one we should check if the item exists in our inventory

	name, quantity, err := validateParams(p)
	if err != nil {
		return err
	}

	it := config.Item{
		Name: name,
	}

	err = i.Increment(it, quantity)
	if err != nil {
		return err
	}

	return nil
}

func decrementMethod(i *config.Inventory, p []string) error {
	// and in this one, check if the quantity is not lower than 1 post-decrease

	name, quantity, err := validateParams(p)
	if err != nil {
		return err
	}

	it := config.Item{
		Name: name,
	}

	err = i.Increment(it, -quantity)
	if err != nil {
		return err
	}

	return nil
}

var im = Command{
	name:        "increase",
	description: "increase the quantity of an item (in the inventory) by n-input amount\n\tdefault increment value is 1 if not specified\n\texample: increase [item's name] [amount]",
	command:     incrementMethod,
}

var dm = Command{
	name:        "decrease",
	description: "decrease the quantity of an item (in the inventory) by n-input amount\n\tdefault decrement value is 1 if not specified\n\texample: decrease [item's name] [amount]",
	command:     decrementMethod,
}
