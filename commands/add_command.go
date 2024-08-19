package commands

import (
	"fmt"
	"inventory/config"
	"strconv"
)

func addCommand(i *config.Inventory, p []string) error {
	switch len(p) {
	case 0:
		return fmt.Errorf("invalid command line arguments: require the item's name as an argument")
	case 1:
		return fmt.Errorf("invalid command line arguments: require the item's price as an argument")
	case 2:
		p = append(p, "1")
	}

	name := p[0]
	amount, aErr := strconv.Atoi(p[1])
	if aErr != nil {
		return aErr
	}

	if amount < 1 {
		return fmt.Errorf("invalid command line arguments: price must be greater than zero")
	}

	quantity, qErr := strconv.Atoi(p[2])
	if qErr != nil {
		return qErr
	}

	if quantity < 1 {
		return fmt.Errorf("invalid command line arguments: quantity must be greater than zero")
	}

	// object creation is expensive, yknow
	price := config.Currency{
		Amount: amount,
		Symbol: config.Symbol,
	}

	it := config.Item{
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}

	iErr := i.Add(it)
	if iErr != nil {
		return iErr
	}
	return nil
}

var ac = Command{
	name:        "add",
	description: "add a new item to the inventory",
	example:     "add [item's name] [price] [quantity]",
	command:     addCommand,
}
