package internal

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/omarattia3143/quiz/internal/collection"
	"github.com/omarattia3143/quiz/internal/controller"
	"github.com/omarattia3143/quiz/internal/service"
	"log"

	"github.com/gofiber/contrib/websocket"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type JsonArray []map[string]any

type App struct {
	httpServer *fiber.App
	database   *mongo.Database

	quizService *service.QuizService
}

func (a *App) Init() {

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

		serverMsg := "Server: " + string(msg)
		if err = c.WriteMessage(mt, []byte(serverMsg)); err != nil {
			log.Println("write:", err)
			break
		}
	}

}
func (a *App) setupDB() {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://admin:yourpassword@localhost:27017"))
	if err != nil {
		log.Fatal("can not connect to database")
	}

	// Defer closing the database connection
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	a.database = client.Database("quiz")
}

func (a *App) setupHttp() {
	app := fiber.New()
	app.Use(cors.New())

	quizController := controller.NewQuizController(a.quizService)

	app.Get("/", a.index)
	app.Get("/api/getquizzes", quizController.GetQuizzes)
	app.Get("/ws", websocket.New(websocketTest))

	log.Fatal(app.Listen(":3000"))
}

func (a *App) setupServices() {
	a.quizService = service.NewQuizService(collection.NewQuizCollection(a.database.Collection("quiz")))
}

func (a *App) index(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
