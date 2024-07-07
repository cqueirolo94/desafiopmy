package domain_test

import (
	"orderprocessor/internal/order/domain"
	"orderprocessor/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Given: an valid symbol.
// When: obtains count of matched orders.
// Then: returns no error.
func TestDefaultObtainMaxMinPendingOrderUsecase_Do_with_successful_process(t *testing.T) {
	// Arrange
	uc, r := setupObtainMaxMinPendingOrderUsecaseTest(t)
	r.On("GetUnmatcheOrderPrices").Return(uint32(4), uint32(3), uint32(2), uint32(1)).Once()

	// Act
	prices := uc.Do()

	// Assert
	assert.Equal(t, uint32(4), prices.BuyPrices.Max)
	assert.Equal(t, uint32(2), prices.BuyPrices.Min)
	assert.Equal(t, uint32(3), prices.SellPrices.Max)
	assert.Equal(t, uint32(1), prices.SellPrices.Min)
}

func setupObtainMaxMinPendingOrderUsecaseTest(t *testing.T) (domain.ObtainMaxMinPendingOrderUsecase, *testdata.OrderRepository) {
	r := testdata.NewOrderRepository(t)
	uc := domain.NewDefaultObtainMaxMinPendingOrderUsecase(r)

	return uc, r
}
