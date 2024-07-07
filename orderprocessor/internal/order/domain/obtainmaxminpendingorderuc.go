package domain

import (
	oRepo "orderprocessor/internal/order/infrastructure/repository"
)

type ObtainMaxMinPendingOrderUsecase interface {
	Do() *UnmatchedOrderPrices
}

type DefaultObtainMaxMinPendingOrderUsecase struct {
	repo oRepo.OrderRepository
}

func (uc *DefaultObtainMaxMinPendingOrderUsecase) Do() *UnmatchedOrderPrices {
	buyMax, sellMax, buyMin, sellMin := uc.repo.GetUnmatcheOrderPrices()

	return &UnmatchedOrderPrices{
		BuyPrices: &OrderPrices{
			Max: buyMax,
			Min: buyMin,
		},
		SellPrices: &OrderPrices{
			Max: sellMax,
			Min: sellMin,
		},
	}
}

func NewDefaultObtainMaxMinPendingOrderUsecase(repo oRepo.OrderRepository) ObtainMaxMinPendingOrderUsecase {
	return &DefaultObtainMaxMinPendingOrderUsecase{
		repo: repo,
	}
}
