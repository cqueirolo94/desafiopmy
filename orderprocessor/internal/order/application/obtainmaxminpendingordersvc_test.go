package application_test

import (
	"orderprocessor/internal/order/application"
	"orderprocessor/internal/order/domain"
	"orderprocessor/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Given: a valid symbol.
// When: matches and persists the order.
// Then: returns no error.
func TestDefaultObtainMaxMinPendingOrderService_Do_with_succesful_process(t *testing.T) {
	// Arrange
	svc, uc := setupObtainMaxMinPendingOrderServiceTest(t)
	ucPrices := &domain.UnmatchedOrderPrices{
		BuyPrices: &domain.OrderPrices{
			Max: 5,
			Min: 1,
		},
		SellPrices: &domain.OrderPrices{
			Max: 4,
			Min: 2,
		},
	}
	uc.On("Do").Return(ucPrices).Once()

	// Act
	prices := svc.Do()

	// Assert
	assert.Equal(t, ucPrices.BuyPrices.Max, prices.BuyPrices.Max)
	assert.Equal(t, ucPrices.BuyPrices.Min, prices.BuyPrices.Min)
	assert.Equal(t, ucPrices.SellPrices.Max, prices.SellPrices.Max)
	assert.Equal(t, ucPrices.SellPrices.Min, prices.SellPrices.Min)
}

func setupObtainMaxMinPendingOrderServiceTest(t *testing.T) (application.ObtainMaxMinPendingOrderService, *testdata.ObtainMaxMinPendingOrderUsecase) {
	uc := testdata.NewObtainMaxMinPendingOrderUsecase(t)
	svc := application.NewDefaultObtainMaxMinPendingOrderService(uc)

	return svc, uc
}
