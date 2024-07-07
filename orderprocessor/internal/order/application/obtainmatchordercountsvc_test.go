package application_test

import (
	"orderprocessor/internal/order/application"
	"orderprocessor/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Given: a valid symbol.
// When: matches and persists the order.
// Then: returns no error.
func TestDefaultObtainMatchOrderCountService_Do_with_succesful_process(t *testing.T) {
	// Arrange

	svc, uc := setupObtainMatchOrderCountServiceTest(t)
	uc.On("Do").Return(uint32(2), nil).Once()

	// Act
	count := svc.Do()

	// Assert
	assert.Equal(t, uint32(2), count)
}

func setupObtainMatchOrderCountServiceTest(t *testing.T) (application.ObtainMatchOrderCountService, *testdata.ObtainMatchOrderCountUsecase) {
	uc := testdata.NewObtainMatchOrderCountUsecase(t)
	svc := application.NewDefaultObtainMatchOrderCountService(uc)

	return svc, uc
}
