package repository_test

import (
	"orderprocessor/internal/order/infrastructure/repository"
	"orderprocessor/pkg"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Given: a nil order.
// When: performs any repository method.
// Then: returns pkg.ErrNilPointer.
func TestInMemoryOrderRepository_with_nil_order(t *testing.T) {
	// Arrange
	r := setup()
	t.Run("PersistBuy", func(t *testing.T) {
		// Act
		err := r.PersistBuy(nil)

		// Assert
		assert.ErrorIs(t, pkg.ErrNilPointer, err)
	})
	t.Run("PersistSell", func(t *testing.T) {
		// Act
		err := r.PersistSell(nil)

		// Assert
		assert.ErrorIs(t, pkg.ErrNilPointer, err)
	})
	t.Run("MatchBuy", func(t *testing.T) {
		// Act
		ok, err := r.MatchBuy(nil)

		// Assert
		assert.False(t, ok)
		assert.ErrorIs(t, pkg.ErrNilPointer, err)
	})
	t.Run("matchSell", func(t *testing.T) {
		// Act
		ok, err := r.MatchSell(nil)

		// Assert
		assert.False(t, ok)
		assert.ErrorIs(t, pkg.ErrNilPointer, err)
	})
}

// Given: a valid order.
// When: performs any repository method, and the order has the wrong type for the method.
// Then: returns pkg.ErrWrongOrderType.
func TestInMemoryOrderRepository_with_wrong_order_type(t *testing.T) {
	// Arrange
	r := setup()
	o := &repository.OrderDTO{
		Symbol: "testsymbol",
		Price:  2,
		Type:   "WrongType",
	}
	t.Run("PersistBuy", func(t *testing.T) {
		// Act
		err := r.PersistBuy(o)

		// Assert
		assert.ErrorIs(t, pkg.ErrWrongOrderType, err)
	})
	t.Run("PersistSell", func(t *testing.T) {
		// Act
		err := r.PersistSell(o)

		// Assert
		assert.ErrorIs(t, pkg.ErrWrongOrderType, err)
	})
	t.Run("MatchBuy", func(t *testing.T) {
		// Act
		ok, err := r.MatchBuy(o)

		// Assert
		assert.False(t, ok)
		assert.ErrorIs(t, pkg.ErrWrongOrderType, err)
	})
	t.Run("matchSell", func(t *testing.T) {
		// Act
		ok, err := r.MatchSell(o)

		// Assert
		assert.False(t, ok)
		assert.ErrorIs(t, pkg.ErrWrongOrderType, err)
	})
}

// Given: a symbol.
// When: get amount of unmatched orders.
// Then: returns the sum of unmatched buy and sell orders.
func TestInMemoryOrderRepository_GetUnmatchedOrderAmount_with_successful_amount_calculation(t *testing.T) {
	// Prearrange
	symbol := "testsymbol"
	buyOrder := &repository.OrderDTO{
		Symbol: symbol,
		Price:  2,
		Type:   "buy",
	}
	sellOrder := &repository.OrderDTO{
		Symbol: symbol,
		Price:  2,
		Type:   "sell",
	}

	t.Run("no orders persisted", func(t *testing.T) {
		// Arrange
		r := setup()

		// Act
		amt := r.GetUnmatchedOrderAmount()

		// Assert
		assert.Equal(t, uint32(0), amt)
	})

	t.Run("one buy order persisted", func(t *testing.T) {
		// Arrange
		r := setup()
		r.PersistBuy(buyOrder)

		// Act
		amt := r.GetUnmatchedOrderAmount()

		// Assert
		assert.Equal(t, uint32(1), amt)
	})

	t.Run("one sell order persisted", func(t *testing.T) {
		// Arrange
		r := setup()
		r.PersistSell(sellOrder)

		// Act
		amt := r.GetUnmatchedOrderAmount()

		// Assert
		assert.Equal(t, uint32(1), amt)
	})

	t.Run("one buy and one sell orders persisted", func(t *testing.T) {
		// Arrange
		r := setup()
		r.PersistBuy(buyOrder)
		r.PersistSell(sellOrder)

		// Act
		amt := r.GetUnmatchedOrderAmount()

		// Assert
		assert.Equal(t, uint32(2), amt)
	})
}

// Given: a valid buy order.
// When: persists the order.
// Then: returns no error.
func TestInMemoryOrderRepository_PersistBuy_with_successful_persist(t *testing.T) {
	// Prearrange
	symbol := "testsymbol"
	o := &repository.OrderDTO{
		Symbol: symbol,
		Price:  2,
		Type:   "buy",
	}
	t.Run("with no matching symbol/price order persisted", func(t *testing.T) {
		// Arrange
		r := setup()

		// Act
		err := r.PersistBuy(o)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, uint32(1), r.GetUnmatchedOrderAmount())
	})

	t.Run("with matching symbol/price order persisted", func(t *testing.T) {
		// Arrange
		r := setup()
		r.PersistBuy(o)

		// Act
		err := r.PersistBuy(o)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, uint32(2), r.GetUnmatchedOrderAmount())
	})
}

// Given: a valid sell order.
// When: persists the order.
// Then: returns no error.
func TestInMemoryOrderRepository_PersistSell_with_successful_persist(t *testing.T) {
	// Prearrange
	symbol := "testsymbol"
	o := &repository.OrderDTO{
		Symbol: symbol,
		Price:  2,
		Type:   "sell",
	}
	t.Run("with no matching symbol/price order persisted", func(t *testing.T) {
		// Arrange
		r := setup()

		// Act
		err := r.PersistSell(o)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, uint32(1), r.GetUnmatchedOrderAmount())
	})

	t.Run("with matching symbol/price order persisted", func(t *testing.T) {
		// Arrange
		r := setup()
		r.PersistSell(o)

		// Act
		err := r.PersistSell(o)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, uint32(2), r.GetUnmatchedOrderAmount())
	})
}

// Given: a valid buy order.
// When: matches the order and there is no sell order for that price.
// Then: returns false and no error.
func TestInMemoryOrderRepository_MatchBuy_with_unsuccessful_match(t *testing.T) {
	t.Run("with no orders for that symbol", func(t *testing.T) {
		// Arrange
		symbol := "testsymbol"
		buyOrder := &repository.OrderDTO{
			Symbol: symbol,
			Price:  2,
			Type:   "buy",
		}
		r := setup()

		// Preassert
		assert.Equal(t, uint32(0), r.GetUnmatchedOrderAmount())

		// Act
		ok, err := r.MatchBuy(buyOrder)

		// Assert
		assert.False(t, ok)
		assert.NoError(t, err)
		assert.Equal(t, uint32(0), r.GetUnmatchedOrderAmount())
	})

}

// Given: a valid buy order.
// When: matches the order.
// Then: returns true and no error.
func TestInMemoryOrderRepository_MatchBuy_with_successful_match(t *testing.T) {
	// Arrange
	symbol := "testsymbol"
	buyOrder := &repository.OrderDTO{
		Symbol: symbol,
		Price:  2,
		Type:   "buy",
	}
	sellOrder := &repository.OrderDTO{
		Symbol: symbol,
		Price:  2,
		Type:   "sell",
	}
	r := setup()
	r.PersistSell(sellOrder)
	r.PersistSell(sellOrder)

	// Preassert
	assert.Equal(t, uint32(2), r.GetUnmatchedOrderAmount())
	assert.Equal(t, uint32(0), r.GetMatchedOrderAmount())

	// Act
	ok, err := r.MatchBuy(buyOrder)

	// Assert
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), r.GetUnmatchedOrderAmount())
	assert.Equal(t, uint32(1), r.GetMatchedOrderAmount())
}

// Given: a valid sell order.
// When: matches the order and there is no buy order for that symbol and price.
// Then: returns false and no error.
func TestInMemoryOrderRepository_MatchSell_with_unsuccessful_match(t *testing.T) {
	t.Run("with no orders for that symbol", func(t *testing.T) {
		// Arrange
		symbol := "testsymbol"
		sellOrder := &repository.OrderDTO{
			Symbol: symbol,
			Price:  2,
			Type:   "sell",
		}
		r := setup()

		// Preassert
		assert.Equal(t, uint32(0), r.GetUnmatchedOrderAmount())

		// Act
		ok, err := r.MatchSell(sellOrder)

		// Assert
		assert.False(t, ok)
		assert.NoError(t, err)
		assert.Equal(t, uint32(0), r.GetUnmatchedOrderAmount())
	})

}

// Given: a valid sell order.
// When: matches the order.
// Then: returns true and no error.
func TestInMemoryOrderRepository_MatchSell_with_successful_match(t *testing.T) {
	// Arrange
	symbol := "testsymbol"
	buyOrder := &repository.OrderDTO{
		Symbol: symbol,
		Price:  2,
		Type:   "buy",
	}
	sellOrder := &repository.OrderDTO{
		Symbol: symbol,
		Price:  2,
		Type:   "sell",
	}
	r := setup()
	r.PersistBuy(buyOrder)
	r.PersistBuy(buyOrder)

	// Preassert
	assert.Equal(t, uint32(2), r.GetUnmatchedOrderAmount())
	assert.Equal(t, uint32(0), r.GetMatchedOrderAmount())

	// Act
	ok, err := r.MatchSell(sellOrder)

	// Assert
	assert.True(t, ok)
	assert.NoError(t, err)
	assert.Equal(t, uint32(1), r.GetUnmatchedOrderAmount())
	assert.Equal(t, uint32(1), r.GetMatchedOrderAmount())
}

// Given: no input.
// When: gets max and min price of buy and sell orders.
// Then: returns the max and min buy price and max and min sell price.
func TestInMemoryOrderRepository_GetUnmatcheOrderPrices(t *testing.T) {
	// Arrange
	symbol := "testsymbol"
	r := setup()
	buyOrder := &repository.OrderDTO{
		Symbol: symbol,
		Price:  2,
		Type:   "buy",
	}
	r.PersistBuy(buyOrder)

	buyOrder.Price = 5
	r.PersistBuy(buyOrder)

	sellOrder := &repository.OrderDTO{
		Symbol: symbol,
		Price:  2,
		Type:   "sell",
	}
	r.PersistSell(sellOrder)

	sellOrder.Price = 1
	r.PersistSell(sellOrder)

	//Act
	maxBuy, maxSell, minBuy, minSell := r.GetUnmatcheOrderPrices()

	//Assert
	assert.Equal(t, uint32(5), maxBuy)
	assert.Equal(t, uint32(2), maxSell)
	assert.Equal(t, uint32(2), minBuy)
	assert.Equal(t, uint32(1), minSell)
}

/*
func TestInMemoryOrderRepository_cleanUp(t *testing.T) {
	// Arrange
	r := setup()
	buyOrder := &repository.OrderDTO{
		Symbol: "testsymbol",
		Price:  2,
		Type:   "buy",
	}
	sellOrder := &repository.OrderDTO{
		Symbol: "testsymbol",
		Price:  2,
		Type:   "sell",
	}
	r.PersistBuy(buyOrder)
	assert.Equal(t, 1, r.GetStoredSize())
	r.MatchBuy(sellOrder)
	assert.Equal(t, 1, r.GetStoredSize())

	//Act
	time.Sleep(5 * time.Second)

	//Assert
	assert.Equal(t, 0, r.GetStoredSize())
}*/

func setup() *repository.InMemoryOrderRepository {
	return (repository.NewInMemoryOrderRepository()).(*repository.InMemoryOrderRepository)
}
