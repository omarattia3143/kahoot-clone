package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/bson"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type JsonArray []map[string]any

var client *mongo.Client
var quizCollection *mongo.Collection

func main() {
	setupDB()
	defer disconnectDB()
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", index)
	app.Get("/api/getquizzes", getQuizzes)
	app.Get("/ws", websocket.New(websocketTest))

	log.Fatal(app.Listen(":3000"))
}

func websocketTest(c *websocket.Conn) {
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
	welcomeMessage := []byte("Welcome to the WebSocket server!")
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

		if err = c.WriteMessage(mt, msg); err != nil {
			log.Println("write:", err)
			break
		}
	}

}
func setupDB() {
	client, _ = mongo.Connect(options.Client().ApplyURI("mongodb://admin:yourpassword@localhost:27017"))
	quizCollection = client.Database("quiz").Collection("quizzes")
}

func disconnectDB() {
	if err := client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}

func index(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func getQuizzes(c *fiber.Ctx) error {

	cursor, err := quizCollection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	var list JsonArray
	err = cursor.All(context.Background(), &list)
	if err != nil {
		panic(err)
	}

	return c.JSON(list)
}
