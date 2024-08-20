package commands

import (
	"fmt"
	"inventory/config"
	"strconv"
)

func updateCommand(i *config.Inventory, p []string) error {
	na, err := strconv.Atoi(p[2])
	if err != nil {
		err = updateAttribute(i, p[0], p[1], p[2])
	} else {
		err = updateAttribute(i, p[0], p[1], na)
	}

	if err != nil {
		return err
	}
	return nil
}

func updateAttribute[T int | string](i *config.Inventory, n, a string, na T) error {
	it, err := i.GetItem(n)
	if err != nil {
		return err
	}

	itd := it.Hash()

	switch a {
	case "name":
		if name, ok := any(na).(string); ok {
			it.Name = name
		} else {
			return fmt.Errorf("expected a string for name")
		}
	case "quantity":
		if quantity, ok := any(na).(int); ok {
			it.Quantity = quantity
		} else {
			return fmt.Errorf("expected an int for quantity")
		}
	case "price":
		if price, ok := any(na).(int); ok {
			it.Quantity = price
		} else {
			return fmt.Errorf("expected an int for price")
		}
	default:
		return fmt.Errorf("unknown attribute: %s", a)
	}

	err = i.Update(itd, it)
	if err != nil {
		return err
	}

	return nil
}

var uc = Command{
	name:        "update",
	description: "updates an item's attributes. e.g., name, quantity, and price",
	example:     "update [item's name] [attribute's name] [new attribute]",
	command:     updateCommand,
}
