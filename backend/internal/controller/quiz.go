package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/omarattia3143/quiz/internal/service"
)

type QuizController struct {
	quizService *service.QuizService
}

func NewQuizController(q *service.QuizService) *QuizController {
	return &QuizController{
		quizService: q,
	}
}

func (c QuizController) GetQuizzes(ctx *fiber.Ctx) error {
	result, err := c.quizService.GetQuizzes()
	if err != nil {
		return err
	}
	return ctx.JSON(result)
}
