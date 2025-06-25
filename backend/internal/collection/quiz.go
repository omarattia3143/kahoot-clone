package collection

import (
	"context"
	"github.com/omarattia3143/quiz/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type QuizCollection struct {
	collection *mongo.Collection
}

func NewQuizCollection(collection *mongo.Collection) *QuizCollection {
	return &QuizCollection{
		collection: collection,
	}
}

func (c QuizCollection) InsertQuiz(quiz entity.Quiz) error {
	_, err := c.collection.InsertOne(context.Background(), quiz)
	return err
}

func (c QuizCollection) GetQuizById(id primitive.ObjectID) (*entity.Quiz, error) {
	quiz := entity.Quiz{}
	result := c.collection.FindOne(context.Background(), bson.M{"_id": id})
	err := result.Decode(&quiz)
	return &quiz, err
}

func (c QuizCollection) GetQuizzes() ([]entity.Quiz, error) {
	var quizzes []entity.Quiz
	cursor, err := c.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return quizzes, err
	}
	err = cursor.All(context.Background(), &quizzes)
	return quizzes, err
}
