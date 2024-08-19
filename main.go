package main

import (
	"bufio"
	"fmt"
	"inventory/commands"
	"inventory/config"
	"os"
	"regexp"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	i, iErr := config.NewInventory()
	if iErr != nil {
		panic(iErr)
	}

	for {
		fmt.Printf("inventory > ")
		sc.Scan()
		c, p := format(sc.Text())
		if _, ok := commands.Commands[c]; !ok {
			fmt.Println("Unknown command: ", c)
			continue
		}

		err := commands.Commands[c].Execute(i, p)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func format(inp string) (string, []string) {
	if len(inp) == 0 {
		return "", []string{}
	}

	// IDK, I asked ChatGPT on this one.
	// I want to make sure that something that is surrounded with ""
	// is regarded as 1 element within the array
	re := regexp.MustCompile(`"([^"]*)"|(\S+)`)
	matches := re.FindAllStringSubmatch(inp, -1)

	var result []string
	for _, match := range matches {
		// Only one of the groups will have a match, pick the first non-empty group
		for _, group := range match[1:] {
			if group != "" {
				result = append(result, group)
				break
			}
		}
	}

	command := strings.ToLower(result[0])
	return command, result[1:]
}
