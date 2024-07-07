package http

import (
	"orderprocessor/internal/order/application"

	"github.com/gofiber/fiber/v2"
)

// ObtainMaxMinPendingOrderController provides a method to handle incoming requests for consulting Max and Min price of pending orders by symbol.
type ObtainMaxMinPendingOrderController struct {
	svc application.ObtainMaxMinPendingOrderService
}

func (c *ObtainMaxMinPendingOrderController) Handle(ctx *fiber.Ctx) error {
	prices := c.svc.Do()

	return ctx.Status(fiber.StatusOK).JSON(&pendingOrderPricesResponse{
		BuyPrices:  (*OrderPricesDTO)(prices.BuyPrices),
		SellPrices: (*OrderPricesDTO)(prices.SellPrices),
	})
}

// NewObtainMaxMinPendingOrderController returns a new instance of ObtainMaxMinPendingOrderController.
func NewObtainMaxMinPendingOrderController(svc application.ObtainMaxMinPendingOrderService) *ObtainMaxMinPendingOrderController {
	return &ObtainMaxMinPendingOrderController{
		svc: svc,
	}
}
