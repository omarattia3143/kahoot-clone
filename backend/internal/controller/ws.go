package controller

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/omarattia3143/quiz/internal/service"
	"log"
)

type WsController struct {
	netService *service.NetService
}

func NewWsController(netService *service.NetService) *WsController {
	return &WsController{netService: netService}
}

func (w *WsController) InitWebSocket(c *websocket.Conn) {
	// c.Locals are added to the *websocket.Conn
	log.Println(c.Locals("allowed"))  // true
	log.Println(c.Params("id"))       // 123
	log.Println(c.Query("v"))         // 1.0
	log.Println(c.Cookies("session")) // ""

	// websocket.Conn bindings https://pkg.go.dev/github.com/fasthttp/websocket?tab=doc#pkg-index
	var (
		mt  int
		msg []byte
		err error
	)
	welcomeMessage := []byte("Client Connected")
	if err = c.WriteMessage(websocket.TextMessage, welcomeMessage); err != nil {
		log.Println("write:", err)
		return
	}

	for {
		if mt, msg, err = c.ReadMessage(); err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", msg)

		w.netService.OnIncomingMessage(c, mt, msg)
	}
}
