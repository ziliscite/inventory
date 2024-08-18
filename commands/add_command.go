package commands

import (
	"fmt"
	"inventory/config"
	"strconv"
)

func addCommand(i *config.Inventory, param []string) error {
	switch len(param) {
	case 0:
		return fmt.Errorf("invalid command line arguments: require the item's name as an argument")
	case 1:
		return fmt.Errorf("invalid command line arguments: require the item's price as an argument")
	case 2:
		param = append(param, "1")
	}

	name := param[0]
	amount, aErr := strconv.Atoi(param[1])
	if aErr != nil {
		return aErr
	}

	price := config.Currency{
		Amount: amount,
		Symbol: config.Symbol,
	}

	quantity, qErr := strconv.Atoi(param[2])
	if qErr != nil {
		return qErr
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
	description: "add a new item to the inventory\n\tquantity is automatically 1 if not specified\n\texample: add [item's name] [price] [quantity]",
	command:     addCommand,
}
