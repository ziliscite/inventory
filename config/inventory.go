package config

import (
	"fmt"
	"inventory/helper"
	"sync"
)

type Inventory struct {
	path  string
	mu    sync.Mutex
	Items map[string]Item `json:"items"`
}

func NewInventory() (*Inventory, error) {
	i := Inventory{
		path:  "database/inventory.json",
		Items: make(map[string]Item),
	}

	err := helper.LoadFromJSON(i.path, &i)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &i, nil
}

func (i *Inventory) Add(it Item) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	id := it.Hash()
	if v, ok := i.Items[id]; ok {
		return fmt.Errorf("%s is already in the inventory\nuse increase or update command instead", v.Name)
	}

	i.Items[id] = it

	err := helper.SaveToJSON(i.path, i)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	fmt.Println("successfully added " + it.Name + " to the inventory")
	return nil
}

func (i *Inventory) Remove(it Item) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	id := it.Hash()
	if _, ok := i.Items[id]; !ok {
		return fmt.Errorf("%s is not in the inventory", it.Name)
	}

	delete(i.Items, id)

	err := helper.SaveToJSON(i.path, i)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	// listen, I know that this looks like a code repetition,
	// but like, 80% of the codes here are error handling, we would have to do it regardless
	fmt.Println("successfully deleted " + it.Name + " from the inventory")
	return nil
}

func (i *Inventory) Increment(it Item, q int) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	id := it.Hash()
	var item Item

	if _, ok := i.Items[id]; !ok {
		return fmt.Errorf("%s is not in the inventory", it.Name)
	} else {
		item = i.Items[id]
	}

	err := item.Increment(q)
	if err != nil {
		return err
	}

	i.Items[id] = item
	sErr := helper.SaveToJSON(i.path, i)
	if sErr != nil {
		return fmt.Errorf(sErr.Error())
	}

	fmt.Println("successfully incremented", it.Name, "by", q, "to the inventory")
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
