package server

import (
	"apigateway/internal/server/config"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// Server contains all the dependencies of the app, and has the logic to start listening for requests / incoming orders.
type Server struct {
	config *config.Config
}

func (sv *Server) Run() {
	app := fiber.New()
	app.Get("/order/match/:symbol", func(c *fiber.Ctx) error {
		symbol := strings.ToUpper(c.Params("symbol"))
		switch symbol {
		case "ALUA":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.AluaUrl))
		case "BBAR":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.BbarUrl))
		case "BMA":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.BmaUrl))
		case "BYMA":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.BymaUrl))
		case "CEPU":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.CepuUrl))
		case "COME":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.ComeUrl))
		case "CRES":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.CresUrl))
		case "EDN":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.EdnUrl))
		case "GGAL":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.GgalUrl))
		case "IRSA":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.IrsaUrl))
		case "LOMA":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.LomaUrl))
		case "MIRG":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.MirgUrl))
		case "PAMP":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.PampUrl))
		case "SUPV":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.SupvUrl))
		case "TECO2":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.Teco2Url))
		case "TGNO4":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.Tgno4Url))
		case "TGSU2":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.Tgsu2Url))
		case "TRAN":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.TranUrl))
		case "TXAR":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.TxarUrl))
		case "VALO":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.ValoUrl))
		case "YPFD":
			sendToAddress(c, fmt.Sprintf("%s/order/match", sv.config.YpfdUrl))
		default:
			c.Status(fiber.StatusInternalServerError).SendString("could not match symbol")
		}

		return nil
	})
	app.Get("/order/pending/:symbol/prices", func(c *fiber.Ctx) error {
		symbol := strings.ToUpper(c.Params("symbol"))
		switch symbol {
		case "ALUA":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.AluaUrl))
		case "BBAR":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.BbarUrl))
		case "BMA":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.BmaUrl))
		case "BYMA":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.BymaUrl))
		case "CEPU":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.CepuUrl))
		case "COME":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.ComeUrl))
		case "CRES":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.CresUrl))
		case "EDN":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.EdnUrl))
		case "GGAL":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.GgalUrl))
		case "IRSA":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.IrsaUrl))
		case "LOMA":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.LomaUrl))
		case "MIRG":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.MirgUrl))
		case "PAMP":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.PampUrl))
		case "SUPV":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.SupvUrl))
		case "TECO2":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.Teco2Url))
		case "TGNO4":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.Tgno4Url))
		case "TGSU2":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.Tgsu2Url))
		case "TRAN":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.TranUrl))
		case "TXAR":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.TxarUrl))
		case "VALO":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.ValoUrl))
		case "YPFD":
			sendToAddress(c, fmt.Sprintf("%s/order/pending/prices", sv.config.YpfdUrl))
		default:
			c.Status(fiber.StatusInternalServerError).SendString("could not match symbol")
		}

		return nil
	})
	app.Listen(fmt.Sprintf(":%s", sv.config.HttpPort))
}

func sendToAddress(c *fiber.Ctx, address string) {
	res, err := http.Get(address)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendStream(res.Body)
	}
	c.Status(res.StatusCode).SendStream(res.Body)
}

// NewServer returns a new instance of Server, with all dependencies created.
func NewServer() *Server {
	config := config.NewConfig()

	return &Server{
		config: config,
	}
}
