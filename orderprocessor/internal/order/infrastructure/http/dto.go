package http

type matchOrderCountResponse struct {
	MatchedOrders uint32 `json:"matchedOrders"`
}

type pendingOrderPricesResponse struct {
	BuyPrices  *OrderPricesDTO `json:"buyPrices"`
	SellPrices *OrderPricesDTO `json:"sellPrices"`
}

type OrderPricesDTO struct {
	Max uint32 `json:"max"`
	Min uint32 `json:"min"`
}
