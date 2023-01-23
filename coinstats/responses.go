package coinstats

import "fmt"

type coinsResponse struct {
	Coins []Coin `json:"coins"`
}

type Coin struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Rank   int     `json:"rank"`
	Price  float32 `json:"price"`
}

func (c Coin) Info() string {
	return fmt.Sprintf("[Rank %d] %s | ID: %s | Name: %s | Price: %f",
		c.Rank, c.Symbol, c.ID, c.Name, c.Price)
}
