package commands

import (
	"fmt"
	"inventory/config"
)

type Command struct {
	name, description, example string
	command                    func(inventory *config.Inventory, params []string) error
}

func (c Command) Execute(i *config.Inventory, p []string) error {
	err := c.command(i, p)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

var Commands map[string]Command

func init() {
	Commands = map[string]Command{
		hc.name: hc,
		ec.name: ec,
		lc.name: lc,
		ac.name: ac,
		rc.name: rc,
		ic.name: ic,
		dc.name: dc,
		uc.name: uc,
	}
}

/*
update commands → updates an item's attributes. e.g., name, quantity, and price
	example: update [item's name] [attribute] [updated attribute] [descriptions]
	// a nice generics practice by the way
*/
