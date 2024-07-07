package application

import (
	"orderprocessor/internal/order/domain"
)

type ObtainMaxMinPendingOrderService interface {
	Do() *UnmatchedOrderPricesDTO
}

type DefaultObtainMaxMinPendingOrderService struct {
	uc domain.ObtainMaxMinPendingOrderUsecase
}

func (sv *DefaultObtainMaxMinPendingOrderService) Do() *UnmatchedOrderPricesDTO {
	prices := sv.uc.Do()
	return &UnmatchedOrderPricesDTO{
		BuyPrices:  (*OrderPricesDTO)(prices.BuyPrices),
		SellPrices: (*OrderPricesDTO)(prices.SellPrices),
	}
}

func NewDefaultObtainMaxMinPendingOrderService(uc domain.ObtainMaxMinPendingOrderUsecase) ObtainMaxMinPendingOrderService {
	return &DefaultObtainMaxMinPendingOrderService{
		uc: uc,
	}
}
