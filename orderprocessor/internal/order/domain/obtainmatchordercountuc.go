package domain

import (
	oRepo "orderprocessor/internal/order/infrastructure/repository"
)

type ObtainMatchOrderCountUsecase interface {
	Do() uint32
}

type DefaultObtainMatchOrderCountUsecase struct {
	repo oRepo.OrderRepository
}

func (uc *DefaultObtainMatchOrderCountUsecase) Do() uint32 {
	return uc.repo.GetMatchedOrderAmount()
}

func NewDefaultObtainMatchOrderCountUsecase(repo oRepo.OrderRepository) ObtainMatchOrderCountUsecase {
	return &DefaultObtainMatchOrderCountUsecase{
		repo: repo,
	}
}
