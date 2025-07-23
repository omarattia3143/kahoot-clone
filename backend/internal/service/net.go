package service

import (
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"strings"
)

type NetService struct {
	quizService *QuizService
}

func NewNetService(quizService *QuizService) *NetService {
	return &NetService{
		quizService: quizService,
	}
}
func (c *NetService) OnIncomingMessage(con *websocket.Conn, mt int, msg []byte) {
	str := string(msg)
	parts := strings.Split(":", str)
	cmd := parts[0]
	argument := parts[1]

	switch cmd {
	case "host":
		{
			fmt.Println("host quiz: ", argument)
		}
	case "join":
		{
			fmt.Println("join quiz: ", argument)
		}
	}
}
