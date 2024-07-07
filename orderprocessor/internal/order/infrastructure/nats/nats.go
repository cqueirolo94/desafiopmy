package nats

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	oApp "orderprocessor/internal/order/application"

	"orderprocessor/internal/server/config"

	"github.com/nats-io/nats.go"
)

type semaGo = struct{}

// OrderReceiver provides a method to receive orders to be processed.
type OrderReceiver interface {
	Do()
}

// NatsOrderReceiver implements OrderReceiver.
type NatsOrderReceiver struct {
	config *config.Config
	svc    oApp.OrderMatchAndPersistService
}

func (nor *NatsOrderReceiver) Do() {
	nc, connErr := nats.Connect(nor.config.NatsUrl)
	if connErr != nil {
		log.Fatal(connErr)
	}
	sem := make(chan struct{}, nor.config.MaxGoroutines)

	sus, susErr := nc.SubscribeSync(fmt.Sprintf("pmy.order.%s", nor.config.Subject))
	if susErr != nil {
		log.Fatal(susErr)
	}

	for {
		msg, nextErr := sus.NextMsg(10 * time.Second)
		if nextErr != nil {
			log.Print(nextErr)
		}

		order := new(OrderDTO)
		json.Unmarshal(msg.Data, order)
		sem <- semaGo{}
		go func(order *OrderDTO, sem chan semaGo) {
			svcErr := nor.svc.Do((*oApp.OrderDTO)(order))
			if svcErr != nil {
				fmt.Println(svcErr.Error())
			}
			<-sem
		}(order, sem)
	}

}

// NewNatsOrderReceiver returns a new instance of NatsOrderReceiver.
func NewNatsOrderReceiver(cfg *config.Config, svc oApp.OrderMatchAndPersistService) OrderReceiver {
	return &NatsOrderReceiver{
		config: cfg,
		svc:    svc,
	}
}
