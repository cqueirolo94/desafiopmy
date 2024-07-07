package application

import (
	"orderprocessor/internal/order/domain"
)

type OrderMatchAndPersistService interface {
	Do(*OrderDTO) error
}

type DefaultOrderMatchAndPersistService struct {
	uc domain.OrderMatchAndPersistUsecase
}

func (sv *DefaultOrderMatchAndPersistService) Do(o *OrderDTO) error {
	return sv.uc.Do((*domain.Order)(o))
}

func NewDefaultOrderMatchAndPersistService(uc domain.OrderMatchAndPersistUsecase) OrderMatchAndPersistService {
	return &DefaultOrderMatchAndPersistService{
		uc: uc,
	}
}
