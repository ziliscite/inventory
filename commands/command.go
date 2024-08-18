package commands

import (
	"fmt"
	"inventory/config"
)

type Command struct {
	name, description string
	command           func(inventory *config.Inventory, param []string) error
}

// Mom, can we have an abstraction?
// We have abstraction at home
// The abstraction in question:

func Execute(cmd Command, inventory *config.Inventory, param []string) error {
	err := cmd.command(inventory, param)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

var Commands map[string]Command

func init() {
	Commands = map[string]Command{
		hc.name: hc,
		lc.name: lc,
		ac.name: ac,
	}
}

/*
help commands → displays all the commands available

list commands → listed all the things inside the inventory
	example: list

add commands → add a new item to the inventory. quantity is automatically 1 if not specified
	example: add [item's name] [price] [quantity]

// we probably needed a pointer to the "inventory" - 1st parameter
	// perhaps play with csv first before an actual database
	// update: lmao, played with json instead

// second parameter would probably be ..., well, the callback parameter(s)

// the commands below can have an optional description as the last parameter

remove commands → remove an existing item from the inventory
	example: remove [item's name] [descriptions]

increase/decrease commands → increase/decrease the quantity of an item (in the inventory) by n-input number
	example: increase [item's name] [amount] [descriptions]

update commands → updates an item's attributes. e.g., name, quantity, and price
	example: update [item's name] [attribute] [updated attribute] [descriptions]
	// a nice generics practice by the way
*/
