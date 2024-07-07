package nats

// OrderDTO is a market buy or sell action to be processed.
type OrderDTO struct {
	// Symbol is the product to be processed.
	Symbol string `json:"symbol"`
	// Price is the amount of currency that the product will be bought or sold.
	Price uint32 `json:"price"`
	// Type is the action to be processed, Buy or Sell.
	Type string `json:"type"`
}
