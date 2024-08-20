package config

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// maybe use this to reinforce the generics
type Name string
type Quantity int

type Item struct {
	Name     string   `json:"name"`
	Quantity int      `json:"quantity"`
	Price    Currency `json:"price"`
}

func (i *Item) Hash() string {
	h := sha256.New()
	h.Write([]byte(strings.ToLower(i.Name)))
	bs := h.Sum(nil)
	hexString := fmt.Sprintf("%x", bs)
	return hexString
}

func (i *Item) Increment(q int) error {
	if v := i.Quantity + q; v < 0 {
		return fmt.Errorf("item's quantity cannot be less than 0")
	}
	i.Quantity += q
	return nil
}

type currencySymbol string

const Symbol currencySymbol = "malizith"

type Currency struct {
	Amount int `json:"amount"`
	Symbol currencySymbol
}
