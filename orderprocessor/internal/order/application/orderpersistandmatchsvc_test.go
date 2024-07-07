package application_test

import (
	"orderprocessor/internal/order/application"
	"orderprocessor/internal/order/domain"
	"orderprocessor/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Given: a valid order.
// When: matches and persists the order.
// Then: returns no error.
func TestDefaultOrderMatchAndPersistService_Do_with_succesful_process(t *testing.T) {
	// Arrange
	o := &application.OrderDTO{
		Symbol: "some_symbol",
		Price:  1,
		Type:   "some_type",
	}
	svc, uc := setupOrderMatchAndPersistServiceTest(t)
	uc.On("Do", (*domain.Order)(o)).Return(nil).Once()

	// Act
	err := svc.Do(o)

	// Assert
	assert.NoError(t, err)
}

// Given: a valid order.
// When: matches and persists the order, and there is a usecase failure.
// Then: returns some error.
func TestDefaultOrderMatchAndPersistService_Do_with_usecase_failure(t *testing.T) {
	// Arrange
	o := &application.OrderDTO{
		Symbol: "some_symbol",
		Price:  1,
		Type:   "some_type",
	}
	svc, uc := setupOrderMatchAndPersistServiceTest(t)
	uc.On("Do", (*domain.Order)(o)).Return(testdata.SomeErr).Once()

	// Act
	err := svc.Do(o)

	// Assert
	assert.ErrorIs(t, err, testdata.SomeErr)
}

func setupOrderMatchAndPersistServiceTest(t *testing.T) (application.OrderMatchAndPersistService, *testdata.OrderMatchAndPersistUsecase) {
	uc := testdata.NewOrderMatchAndPersistUsecase(t)
	svc := application.NewDefaultOrderMatchAndPersistService(uc)

	return svc, uc
}
