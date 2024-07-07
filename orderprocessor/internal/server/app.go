package server

import (
	"fmt"
	oApp "orderprocessor/internal/order/application"
	oDomain "orderprocessor/internal/order/domain"
	oHttp "orderprocessor/internal/order/infrastructure/http"
	oNats "orderprocessor/internal/order/infrastructure/nats"
	oRepo "orderprocessor/internal/order/infrastructure/repository"
	"orderprocessor/internal/server/config"

	"github.com/gofiber/fiber/v2"
)

// Server contains all the dependencies of the app, and has the logic to start listening for requests / incoming orders.
type Server struct {
	config                       *config.Config
	orderReceiver                oNats.OrderReceiver
	obtainMatchOrderCountCtrl    *oHttp.ObtainMatchOrderCountController
	obtainMaxMinPendingOrderCtrl *oHttp.ObtainMaxMinPendingOrderController
}

func (sv *Server) Run() {
	go sv.orderReceiver.Do()
	app := fiber.New()
	app.Get("/order/match/", sv.obtainMatchOrderCountCtrl.Handle)
	app.Get("/order/pending/prices", sv.obtainMaxMinPendingOrderCtrl.Handle)
	app.Listen(fmt.Sprintf(":%s", sv.config.HttpPort))
}

// NewServer returns a new instance of Server, with all dependencies created.
func NewServer() *Server {
	config := config.NewConfig()
	orderRepo := oRepo.NewInMemoryOrderRepository()

	obtainMaxMinPendingOrderUC := oDomain.NewDefaultObtainMaxMinPendingOrderUsecase(orderRepo)
	obtainMatchOrderCountUC := oDomain.NewDefaultObtainMatchOrderCountUsecase(orderRepo)

	orderMatchAndPersistUC := oDomain.NewDefaultOrderMatchAndPersistUsecase(orderRepo)

	orderMatchAndPersistSVC := oApp.NewDefaultOrderMatchAndPersistService(orderMatchAndPersistUC)
	obtainMatchOrderCountSVC := oApp.NewDefaultObtainMatchOrderCountService(obtainMatchOrderCountUC)
	obtainMaxMinPendingOrderSVC := oApp.NewDefaultObtainMaxMinPendingOrderService(obtainMaxMinPendingOrderUC)

	orderReceiver := oNats.NewNatsOrderReceiver(config, orderMatchAndPersistSVC)
	obtainMatchOrderCountCtrl := oHttp.NewObtainMatchOrderCountController(obtainMatchOrderCountSVC)
	obtainMaxMinPendingOrderCtrl := oHttp.NewObtainMaxMinPendingOrderController(obtainMaxMinPendingOrderSVC)

	return &Server{
		config:                       config,
		orderReceiver:                orderReceiver,
		obtainMatchOrderCountCtrl:    obtainMatchOrderCountCtrl,
		obtainMaxMinPendingOrderCtrl: obtainMaxMinPendingOrderCtrl,
	}
}
