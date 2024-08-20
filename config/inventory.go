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

func (i *Inventory) Update(itd string, nit Item) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	id := nit.Hash()
	// if id == itd, it means that we update the quantity or price
	// so, it is allowed to be already in the inventory
	if v, ok := i.Items[id]; ok && id != itd {
		return fmt.Errorf("%s is already in the inventory", v.Name)
	}

	pit := i.Items[itd]
	delete(i.Items, itd)

	i.Items[id] = nit

	err := helper.SaveToJSON(i.path, i)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	fmt.Println("successfully updated", pit, "into", nit)
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

// GetItem takes an item's name and returns a copy of the item if it exists
func (i *Inventory) GetItem(n string) (Item, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	it := Item{Name: n}

	id := it.Hash()
	var item Item

	if _, ok := i.Items[id]; !ok {
		return it, fmt.Errorf("%s is not in the inventory", it.Name)
	} else {
		item = i.Items[id]
	}

	return item, nil
}
