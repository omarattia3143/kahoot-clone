package internal

import (
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
	a.setupDB()
	a.setupServices()
	a.setupHttp()
	log.Fatal(a.httpServer.Listen(":3000"))
}

func (a *App) setupDB() {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://admin:yourpassword@localhost:27017"))
	if err != nil {
		log.Fatal("can not connect to database")
	}

	// Defer closing the database connection
	//defer func() {
	//	if err := client.Disconnect(context.Background()); err != nil {
	//		log.Printf("Error disconnecting from MongoDB: %v", err)
	//	}
	//}()

	a.database = client.Database("quiz")
}

func (a *App) setupHttp() {
	app := fiber.New()
	app.Use(cors.New())

	app.Get("/", a.index)

	quizController := controller.NewQuizController(a.quizService)
	app.Get("/api/getquizzes", quizController.GetQuizzes)

	wsController := controller.NewWsController()
	app.Get("/ws", websocket.New(wsController.InitWebSocket))

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}

	a.httpServer = app
}

func (a *App) setupServices() {
	a.quizService = service.NewQuizService(collection.NewQuizCollection(a.database.Collection("quizzes")))
}

func (a *App) index(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
