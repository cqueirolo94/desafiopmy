package domain_test

import (
	"orderprocessor/internal/order/domain"
	"orderprocessor/internal/order/infrastructure/repository"
	"orderprocessor/pkg"
	"orderprocessor/testdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Given: a nil order.
// When: matches and persists the order.
// Then: return pkg.ErrNilPointer.
func TestDefaultOrderMatchAndPersistUsecase_Do_with_nil_order(t *testing.T) {
	// Arrange
	uc, _ := setupOrderMatchAndPersistUsecaseTest(t)

	// Act
	err := uc.Do(nil)

	// Assert
	assert.ErrorIs(t, err, pkg.ErrNilPointer)
}

// Given: a valid order.
// When: matches and persists the order, and the order type is not buy or sell.
// Then: return pkg.ErrWrongOrderType.
func TestDefaultOrderMatchAndPersistUsecase_Do_with_wrong_order_type(t *testing.T) {
	// Arrange
	o := &domain.Order{
		Symbol: "some_symbol",
		Price:  1,
		Type:   "some_type",
	}
	uc, _ := setupOrderMatchAndPersistUsecaseTest(t)

	// Act
	err := uc.Do(o)

	// Assert
	assert.ErrorIs(t, err, pkg.ErrWrongOrderType)
}

// Given: a valid order.
// When: matches and persists the order, and the order type matches.
// Then: return nil.
func TestDefaultOrderMatchAndPersistUsecase_Do_with_matching_order(t *testing.T) {
	// Prearrange
	uc, r := setupOrderMatchAndPersistUsecaseTest(t)
	t.Run("buy order", func(t *testing.T) {
		// Arrange
		o := &domain.Order{
			Symbol: "some_symbol",
			Price:  1,
			Type:   "buy",
		}
		r.On("MatchBuy", (*repository.OrderDTO)(o)).Return(true, nil).Once()

		// Act
		err := uc.Do(o)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("sell order", func(t *testing.T) {
		// Arrange
		o := &domain.Order{
			Symbol: "some_symbol",
			Price:  1,
			Type:   "sell",
		}
		r.On("MatchSell", (*repository.OrderDTO)(o)).Return(true, nil).Once()

		// Act
		err := uc.Do(o)

		// Assert
		assert.NoError(t, err)
	})
}

// Given: a valid order.
// When: matches and persists the order, and the repository returns some matching error.
// Then: return some error.
func TestDefaultOrderMatchAndPersistUsecase_Do_with_matching_error(t *testing.T) {
	// Prearrange
	uc, r := setupOrderMatchAndPersistUsecaseTest(t)
	t.Run("buy order", func(t *testing.T) {
		// Arrange
		o := &domain.Order{
			Symbol: "some_symbol",
			Price:  1,
			Type:   "buy",
		}
		r.On("MatchBuy", (*repository.OrderDTO)(o)).Return(false, testdata.SomeErr).Once()

		// Act
		err := uc.Do(o)

		// Assert
		assert.ErrorIs(t, err, testdata.SomeErr)
	})

	t.Run("sell order", func(t *testing.T) {
		// Arrange
		o := &domain.Order{
			Symbol: "some_symbol",
			Price:  1,
			Type:   "sell",
		}
		r.On("MatchSell", (*repository.OrderDTO)(o)).Return(false, testdata.SomeErr).Once()

		// Act
		err := uc.Do(o)

		// Assert
		assert.ErrorIs(t, err, testdata.SomeErr)
	})
}

// Given: a valid order.
// When: matches and persists the order, with no match, and the repository returns some persistence error.
// Then: return some error.
func TestDefaultOrderMatchAndPersistUsecase_Do_with_persist_error(t *testing.T) {
	// Prearrange
	uc, r := setupOrderMatchAndPersistUsecaseTest(t)
	t.Run("buy order", func(t *testing.T) {
		// Arrange
		o := &domain.Order{
			Symbol: "some_symbol",
			Price:  1,
			Type:   "buy",
		}
		r.On("MatchBuy", (*repository.OrderDTO)(o)).Return(false, nil).Once()
		r.On("PersistBuy", (*repository.OrderDTO)(o)).Return(testdata.SomeErr).Once()

		// Act
		err := uc.Do(o)

		// Assert
		assert.ErrorIs(t, err, testdata.SomeErr)
	})

	t.Run("sell order", func(t *testing.T) {
		// Arrange
		o := &domain.Order{
			Symbol: "some_symbol",
			Price:  1,
			Type:   "sell",
		}
		r.On("MatchSell", (*repository.OrderDTO)(o)).Return(false, nil).Once()
		r.On("PersistSell", (*repository.OrderDTO)(o)).Return(testdata.SomeErr).Once()

		// Act
		err := uc.Do(o)

		// Assert
		assert.ErrorIs(t, err, testdata.SomeErr)
	})
}

// Given: a valid order.
// When: matches and persists the order, with no match, and no persistence error.
// Then: return no error.
func TestDefaultOrderMatchAndPersistUsecase_Do_with_successful_persistence(t *testing.T) {
	// Prearrange
	uc, r := setupOrderMatchAndPersistUsecaseTest(t)
	t.Run("buy order", func(t *testing.T) {
		// Arrange
		o := &domain.Order{
			Symbol: "some_symbol",
			Price:  1,
			Type:   "buy",
		}
		r.On("MatchBuy", (*repository.OrderDTO)(o)).Return(false, nil).Once()
		r.On("PersistBuy", (*repository.OrderDTO)(o)).Return(nil).Once()

		// Act
		err := uc.Do(o)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("sell order", func(t *testing.T) {
		// Arrange
		o := &domain.Order{
			Symbol: "some_symbol",
			Price:  1,
			Type:   "sell",
		}
		r.On("MatchSell", (*repository.OrderDTO)(o)).Return(false, nil).Once()
		r.On("PersistSell", (*repository.OrderDTO)(o)).Return(nil).Once()

		// Act
		err := uc.Do(o)

		// Assert
		assert.NoError(t, err)
	})
}

func setupOrderMatchAndPersistUsecaseTest(t *testing.T) (domain.OrderMatchAndPersistUsecase, *testdata.OrderRepository) {
	r := testdata.NewOrderRepository(t)
	uc := domain.NewDefaultOrderMatchAndPersistUsecase(r)

	return uc, r
}
