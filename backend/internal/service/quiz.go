package service

import (
	"github.com/omarattia3143/quiz/internal/collection"
	"github.com/omarattia3143/quiz/internal/entity"
)

type QuizService struct {
	quizCollection *collection.QuizCollection
}

func NewQuizService(quizCollection *collection.QuizCollection) *QuizService {
	return &QuizService{
		quizCollection: quizCollection,
	}
}

func (s QuizService) GetQuizzes() ([]entity.Quiz, error) {
	return s.quizCollection.GetQuizzes()
}
