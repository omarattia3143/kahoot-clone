package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omarattia3143/quiz/internal/service"
)

type QuizController struct {
	quizService *service.QuizService
}

func NewQuizController(quizService *service.QuizService) *QuizController {
	return &QuizController{
		quizService: quizService,
	}
}

func (c QuizController) GetQuizzes(ctx *fiber.Ctx) error {
	result, err := c.quizService.GetQuizzes()
	if err != nil {
		return err
	}
	return ctx.JSON(result)
}
