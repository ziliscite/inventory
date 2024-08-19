package commands

import (
	"fmt"
	"inventory/config"
	"os"
	"text/tabwriter"
)

func listCommand(i *config.Inventory, p []string) error {
	fmt.Println("Items: ")

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 4, 0, '\t', 4)

	_, _ = fmt.Fprintf(w, "Id\t | Name \t| Quantity \t| Price\n")
	for k, v := range i.Items {
		_, _ = fmt.Fprintf(
			w,
			"%v\t | %s \t| %d \t| %d %s\n",
			k,
			v.Name,
			v.Quantity,
			v.Price.Amount,
			v.Price.Symbol,
		)
	}

	err := w.Flush()
	if err != nil {
		return err
	}

	return nil
}

var lc = Command{
	name:        "list",
	description: "list every items in the inventory",
	example:     "list",
	command:     listCommand,
}
