package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type Inventory struct {
	path  string
	mu    sync.Mutex
	Items map[string]Item `json:"items"`
}

// Decoding the data of a JSON file
func loadFromJSON(fileName string, key interface{}) error {
	in, err := os.Open(fileName)
	if err != nil {
		return err
	}

	// Check if the file is empty
	fileInfo, err := in.Stat()
	if err != nil {
		return err
	}

	if fileInfo.Size() == 0 {
		// If the file is empty, return nil (no error) as there is nothing to load
		return nil
	}

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}

	err = in.Close()
	if err != nil {
		return err
	}

	return nil
}

func NewInventory() (*Inventory, error) {
	i := Inventory{
		path:  "database/inventory.json",
		Items: make(map[string]Item),
	}

	err := loadFromJSON(i.path, &i)
	if err != nil {
		return nil, err
	}

	return &i, nil
}

// Encoding data to JSON file
func saveToJSON(fileName *os.File, key interface{}) error {
	encodeJSON := json.NewEncoder(fileName)
	err := encodeJSON.Encode(key)
	if err != nil {
		return fmt.Errorf("error writing to JSON file: %v", err)
	}

	return nil
}

// Yeah, I copied both the loadFromJSON and saveToJSON method from somewhere,
// but I made my fair share of adjustment so that it would fit in my code.
// So, hey, why not

func (i *Inventory) Add(it Item) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	file, err := os.OpenFile(i.path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("Failed to open file: %v\n", err)
	}

	id := it.Hash()
	if v, ok := i.Items[id]; ok {
		// should've checked it with something like, id?
		// i.Items.id == item.id or something

		// guess what
		return fmt.Errorf("%s is already in the inventory\nuse increase or update command instead", v.Name)
	}

	// I know, "Why the short variable name?", same, man, same
	// However, I don't make the rules, Gophers are deranged
	i.Items[id] = it

	err = saveToJSON(file, i)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	fmt.Println("successfully added " + it.Name + " to the inventory")
	return nil
}

func (i *Inventory) Display() error {
	// honestly, it could be formatted better, but it works for now
	fmt.Println("items: ")
	for _, v := range i.Items {
		fmt.Println("\tname: ", v.Name)
		fmt.Println("\tprice: ", v.Price.Amount, "", v.Price.Symbol)
		fmt.Println("\tquantity", v.Quantity)
		fmt.Println()
	}
	return nil
}
