package config

import (
	"fmt"
	"sync"
)

// This is an extra, you know, like some fantasy store?
// Yeah, I'm going to make that after the inventory repl is finished.
// Or not?

type User struct {
	inventory []Item
	balance   Currency
	mu        sync.Mutex
}

func (u *User) Buy(i Item) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.balance.Amount < i.Price.Amount {
		return fmt.Errorf(
			"your money is %d %s less than the %s's price",
			u.balance.Amount-i.Price.Amount,
			u.balance.Symbol,
			i.Name,
		)
	}
	u.balance.Amount -= i.Price.Amount
	u.inventory = append(u.inventory, i)
	fmt.Println("transaction is successful\nyour remaining balance is", u.balance.Amount, u.balance.Symbol)
	return nil
}

// The second argument would be to the vendor? I think I should make a vendor struct.
// Also applies to the "Buy" function? IDK

func (u *User) Sell(i Item) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	return nil
}
