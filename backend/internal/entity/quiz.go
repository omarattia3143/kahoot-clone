package entity

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Quiz struct {
	Id        primitive.ObjectID
	Name      string
	Questions []QuizQuestion
}

type QuizQuestion struct {
	Id          uuid.UUID
	Name        string
	QuizChoices []QuizChoice
}

type QuizChoice struct {
	Id      uuid.UUID
	Name    string
	Correct bool
}
