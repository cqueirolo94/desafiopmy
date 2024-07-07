package domain

// Order is a market buy or sell action to be processed.
type Order struct {
	// Symbol is the product to be processed.
	Symbol string
	// Price is the amount of currency that the product will be bought or sold.
	Price uint32
	// Type is the action to be processed, Buy or Sell.
	Type string
}

type UnmatchedOrderPrices struct {
	BuyPrices  *OrderPrices
	SellPrices *OrderPrices
}

type OrderPrices struct {
	Max uint32
	Min uint32
}
