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
func TestDefaultObtainMatchOrderCountUsecase_Do_with_successful_process(t *testing.T) {
	// Arrange
	uc, r := setupObtainMatchOrderCountUsecaseTest(t)
	r.On("GetMatchedOrderAmount").Return(uint32(2)).Once()

	// Act
	count := uc.Do()

	// Assert
	assert.Equal(t, uint32(2), count)
}

func setupObtainMatchOrderCountUsecaseTest(t *testing.T) (domain.ObtainMatchOrderCountUsecase, *testdata.OrderRepository) {
	r := testdata.NewOrderRepository(t)
	uc := domain.NewDefaultObtainMatchOrderCountUsecase(r)

	return uc, r
}
