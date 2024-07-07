package domain

import (
	"fmt"
	oRepo "orderprocessor/internal/order/infrastructure/repository"
	"orderprocessor/pkg"
)

type OrderMatchAndPersistUsecase interface {
	Do(*Order) error
}

type DefaultOrderMatchAndPersistUsecase struct {
	repo oRepo.OrderRepository
}

func (uc *DefaultOrderMatchAndPersistUsecase) Do(o *Order) error {
	if o == nil {
		return pkg.ErrNilPointer
	}

	if o.Type != "buy" && o.Type != "sell" {
		return pkg.ErrWrongOrderType
	}

	if o.Type == "buy" {
		return uc.process((*oRepo.OrderDTO)(o), uc.repo.PersistBuy, uc.repo.MatchBuy)
	} else {
		return uc.process((*oRepo.OrderDTO)(o), uc.repo.PersistSell, uc.repo.MatchSell)
	}
}

func (uc *DefaultOrderMatchAndPersistUsecase) process(order *oRepo.OrderDTO, persistF func(*oRepo.OrderDTO) error, matchF func(*oRepo.OrderDTO) (bool, error)) error {
	isMatch, mErr := matchF(order)
	if mErr != nil {
		return mErr
	}
	if isMatch {
		fmt.Printf("se produjo un match para el producto %s con precio %d\n", order.Symbol, order.Price)
		return nil
	}

	return persistF(order)
}

func NewDefaultOrderMatchAndPersistUsecase(repo oRepo.OrderRepository) OrderMatchAndPersistUsecase {
	return &DefaultOrderMatchAndPersistUsecase{
		repo: repo,
	}
}
