package service

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"strings"
	"time"
)

type NetService struct {
	quizService *QuizService
	host        *websocket.Conn
	tick        int
}

func NewNetService(quizService *QuizService) *NetService {
	return &NetService{
		quizService: quizService,
	}
}
func (c *NetService) OnIncomingMessage(con *websocket.Conn, mt int, msg []byte) {
	str := string(msg)
	parts := strings.Split(str, ":")
	cmd := parts[0]
	argument := parts[1]

	switch cmd {
	case "host":
		{
			fmt.Println("host quiz:", argument)
			c.host = con
			c.tick = 100
			go func() {
				for {
					c.tick--
					err := c.host.WriteMessage(mt, []byte(fmt.Sprintf("tick: %d", c.tick)))
					if err != nil {
						fmt.Println("error sending tick")
					}
					time.Sleep(time.Second)
				}
			}()
			break
		}
	case "join":
		{
			fmt.Println("join quiz:", argument)
			err := c.host.WriteMessage(mt, []byte("A player has joined"))
			if err != nil {
				fmt.Println("Player error joining")
			}
			break
		}
	}
}
func (c *NetService) PacketToBytes(packet any) ([]byte, error) {
	var packetId uint = 0
	bytes, err := json.Marshal(packet)
	if err != nil {
		return nil, err
	}
	final := append([]byte{byte(packetId)}, bytes...)
	return final, nil
}

func (c *NetService) SendPacket(connection *websocket.Conn, packet any) error {
	bytes, err := c.PacketToBytes(packet)
	if err != nil {
		return err
	}
	return connection.WriteMessage(websocket.BinaryMessage, bytes)
}
