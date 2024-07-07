package repository

import (
	"orderprocessor/pkg"
	"sync"
	"time"
)

// OrderRepository provides methods to persist and match orders.
type OrderRepository interface {
	// MatchBuy returns a boolean if the given buy order matches a sell order.
	MatchBuy(*OrderDTO) (bool, error)
	// MatchBuy returns a boolean if the given sell order matches a buy order.
	MatchSell(*OrderDTO) (bool, error)
	// PersistBuy persists a buy order. Returns an error if something unexpected happens.
	PersistBuy(*OrderDTO) error
	// PersistSell persists a sell order. Returns an error if something unexpected happens.
	PersistSell(*OrderDTO) error
	// GetMatchedOrderAmount returns the number of matched orders.
	GetMatchedOrderAmount() uint32
	// GetMinUnmatcheOrderPrice returns the max and min prices of unmatched orders (buy and sell).
	GetUnmatcheOrderPrices() (uint32, uint32, uint32, uint32)
}

// InMemoryOrderRepository implements OrderRepository.
type InMemoryOrderRepository struct {
	buymutex      sync.Mutex
	sellmutex     sync.Mutex
	buyOrderMap   map[uint32]uint32
	sellOrderMap  map[uint32]uint32
	umOrderAmount uint32
	mOrderAmount  uint32
}

func (r *InMemoryOrderRepository) PersistBuy(o *OrderDTO) error {
	if o == nil {
		return pkg.ErrNilPointer
	}
	if o.Type != "buy" {
		return pkg.ErrWrongOrderType
	}

	r.buymutex.Lock()
	r.buyOrderMap[o.Price]++
	r.buymutex.Unlock()

	r.umOrderAmount++

	return nil
}

func (r *InMemoryOrderRepository) PersistSell(o *OrderDTO) error {
	if o == nil {
		return pkg.ErrNilPointer
	}
	if o.Type != "sell" {
		return pkg.ErrWrongOrderType
	}

	r.sellmutex.Lock()
	r.sellOrderMap[o.Price]++
	r.sellmutex.Unlock()

	r.umOrderAmount++

	return nil
}

func (r *InMemoryOrderRepository) MatchBuy(o *OrderDTO) (bool, error) {
	if o == nil {
		return false, pkg.ErrNilPointer
	}
	if o.Type != "buy" {
		return false, pkg.ErrWrongOrderType
	}

	r.sellmutex.Lock()
	if r.sellOrderMap[o.Price] != 0 {
		r.sellOrderMap[o.Price]--
		r.sellmutex.Unlock()
		r.umOrderAmount--
		r.mOrderAmount++
		return true, nil
	}
	r.sellmutex.Unlock()

	return false, nil
}

func (r *InMemoryOrderRepository) MatchSell(o *OrderDTO) (bool, error) {
	if o == nil {
		return false, pkg.ErrNilPointer
	}
	if o.Type != "sell" {
		return false, pkg.ErrWrongOrderType
	}

	r.buymutex.Lock()
	if r.buyOrderMap[o.Price] != 0 {
		r.buyOrderMap[o.Price]--
		r.buymutex.Unlock()
		r.umOrderAmount--
		r.mOrderAmount++
		return true, nil
	}
	r.buymutex.Unlock()

	return false, nil
}

func (r *InMemoryOrderRepository) GetUnmatchedOrderAmount() uint32 {
	return r.umOrderAmount
}

func (r *InMemoryOrderRepository) GetMatchedOrderAmount() uint32 {
	return r.mOrderAmount
}

func (r *InMemoryOrderRepository) GetUnmatcheOrderPrices() (uint32, uint32, uint32, uint32) {
	var buyMax, sellMax, buyMin, sellMin uint32
	buyMin--
	sellMin--

	r.buymutex.Lock()
	for price := range r.buyOrderMap {
		if price > buyMax {
			buyMax = price
		}
		if price < buyMin {
			buyMin = price
		}
	}
	r.buymutex.Unlock()

	r.sellmutex.Lock()
	for price := range r.sellOrderMap {
		if price > sellMax {
			sellMax = price
		}
		if price < sellMin {
			sellMin = price
		}
	}
	r.sellmutex.Unlock()

	return buyMax, sellMax, buyMin, sellMin
}

func (r *InMemoryOrderRepository) cleanUp() {
	for range time.Tick(time.Second * 1) {
		r.buymutex.Lock()
		for price, qtty := range r.buyOrderMap {
			if qtty == 0 {
				delete(r.buyOrderMap, price)
			}
		}
		r.buymutex.Unlock()

		r.sellmutex.Lock()
		for price, qtty := range r.sellOrderMap {
			if qtty == 0 {
				delete(r.sellOrderMap, price)
			}
		}
		r.sellmutex.Unlock()
	}
}

func (r *InMemoryOrderRepository) GetStoredSize() int {
	var buysize, sellsize int

	r.buymutex.Lock()
	buysize = len(r.buyOrderMap)
	r.buymutex.Unlock()

	r.sellmutex.Lock()
	sellsize = len(r.sellOrderMap)
	r.sellmutex.Unlock()

	return buysize + sellsize
}

// NewInMemoryOrderRepository returns a new in-memory implementation of OrderRepository.
func NewInMemoryOrderRepository() OrderRepository {
	r := &InMemoryOrderRepository{
		buyOrderMap:  make(map[uint32]uint32),
		sellOrderMap: make(map[uint32]uint32),
	}
	go r.cleanUp()

	return r
}
