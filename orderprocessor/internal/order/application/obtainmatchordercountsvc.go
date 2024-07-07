package application

import (
	"orderprocessor/internal/order/domain"
)

type ObtainMatchOrderCountService interface {
	Do() uint32
}

type DefaultObtainMatchOrderCountService struct {
	uc domain.ObtainMatchOrderCountUsecase
}

func (sv *DefaultObtainMatchOrderCountService) Do() uint32 {
	return sv.uc.Do()
}

func NewDefaultObtainMatchOrderCountService(uc domain.ObtainMatchOrderCountUsecase) ObtainMatchOrderCountService {
	return &DefaultObtainMatchOrderCountService{
		uc: uc,
	}
}
