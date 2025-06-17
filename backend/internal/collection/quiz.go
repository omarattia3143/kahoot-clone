package collection

import (
	"context"
	"github.com/omarattia3143/quiz/internal/entity"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type QuizCollection struct {
	collection *mongo.Collection
}

func (c QuizCollection) InsertQuiz(quiz entity.Quiz) error {
	_, err := c.collection.InsertOne(context.Background(), quiz)
	return err
}
