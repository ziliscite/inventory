package commands

import (
	"fmt"
	"inventory/config"
	"os"
	"text/tabwriter"
)

func helpCommand(i *config.Inventory, p []string) error {
	fmt.Println("Commands: ")

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 4, 0, '\t', 4)

	_, _ = fmt.Fprintf(w, "Command\t | Description\t | Example\n")
	for _, c := range Commands {
		_, _ = fmt.Fprintf(
			w,
			"%s\t | %s\t | %s\n",
			c.name,
			c.description,
			c.example,
		)
	}

	err := w.Flush()
	if err != nil {
		return err
	}

	fmt.Println("Tip: use \"[param]\" if a parameter is more than 1 word | example: add \"Sword Of The Damned\" 524 1")
	return nil
}

var hc = Command{
	name:        "help",
	description: "displays all the commands available",
	example:     "help",
	command:     helpCommand,
}

func exitCommand(i *config.Inventory, p []string) error {
	os.Exit(0)
	return nil
}

var ec = Command{
	name:        "exit",
	description: "exit the program",
	example:     "exit",
	command:     exitCommand,
}
