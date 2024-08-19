package config

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

type Item struct {
	Name     string   `json:"name"`
	Quantity int      `json:"quantity"`
	Price    Currency `json:"price"`
}

func (i Item) Hash() string {
	h := sha256.New()
	h.Write([]byte(strings.ToLower(i.Name)))
	bs := h.Sum(nil)
	hexString := fmt.Sprintf("%x", bs)
	return hexString
}

// I've tried countless number of times to erase the white spaces
// between these things below, but IntelliJ seems to not like it.
// They always add it back in, so I figured, why bother?

type currencySymbol string

const Symbol currencySymbol = "malizith"

type Currency struct {
	Amount int `json:"amount"`
	Symbol currencySymbol
}
