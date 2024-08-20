package commands

import (
	"cmp"
	"fmt"
	"inventory/config"
	"os"
	"slices"
	"text/tabwriter"
)

func getItems(i *config.Inventory) []config.Item {
	ordered := make([]config.Item, 0)
	for _, v := range i.Items {
		ordered = append(ordered, v)
	}

	return ordered
}

func orderByPrice(i *config.Inventory, o string) []config.Item {
	ordered := getItems(i)

	slices.SortFunc(ordered, func(a, b config.Item) int {
		if o == "ascending" {
			return cmp.Compare(a.Price.Amount, b.Price.Amount)
		}
		return cmp.Compare(b.Price.Amount, a.Price.Amount)
	})

	return ordered
}

func orderByQuantity(i *config.Inventory, o string) []config.Item {
	ordered := getItems(i)

	slices.SortFunc(ordered, func(a, b config.Item) int {
		if o == "descending" {
			return cmp.Compare(a.Quantity, b.Quantity)
		}
		return cmp.Compare(b.Quantity, a.Quantity)
	})

	return ordered
}

func validateListCommand(p []string) (string, string, error) {
	if len(p) == 0 {
		p = append(p, "price")
	}
	if len(p) == 1 {
		p = append(p, "ascending")
	}

	if p[0] != "price" && p[0] != "quantity" {
		return "", "", fmt.Errorf("invalid attribute\nvalid attributes are 'price' or 'quantity'")
	}

	if p[1] != "ascending" && p[1] != "descending" {
		return "", "", fmt.Errorf("invalid sort order\nvalid sort orders are 'ascending' or 'descending'")
	}

	return p[0], p[1], nil
}

func listCommand(i *config.Inventory, p []string) error {
	a, o, err := validateListCommand(p)
	if err != nil {
		return err
	}

	fmt.Println("Items: ")

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 4, 0, '\t', 4)

	var it []config.Item
	if a == "price" {
		it = orderByPrice(i, o)
	} else if a == "quantity" {
		it = orderByQuantity(i, o)
	}

	_, _ = fmt.Fprintf(w, "Id\t | Name \t| Quantity \t| Price\n")
	for k, v := range it {
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

	err = w.Flush()
	if err != nil {
		return err
	}

	return nil
}

var lc = Command{
	name:        "list",
	description: "list every items in the inventory. can sort by item's attribute",
	example:     "list [price] [ascending / descending]",
	command:     listCommand,
}
