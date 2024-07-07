package http

import (
	"orderprocessor/internal/order/application"

	"github.com/gofiber/fiber/v2"
)

// ObtainMatchOrderCountController provides a method to handle incoming requests for consulting the number of matching orders by symbol.
type ObtainMatchOrderCountController struct {
	svc application.ObtainMatchOrderCountService
}

func (c *ObtainMatchOrderCountController) Handle(ctx *fiber.Ctx) error {
	count := c.svc.Do()

	return ctx.Status(fiber.StatusOK).JSON(&matchOrderCountResponse{
		MatchedOrders: count,
	})
}

// NewObtainMatchOrderCountController returns a new instance of ObtainMatchOrderCountController.
func NewObtainMatchOrderCountController(svc application.ObtainMatchOrderCountService) *ObtainMatchOrderCountController {
	return &ObtainMatchOrderCountController{
		svc: svc,
	}
}
