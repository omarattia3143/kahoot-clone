package entity

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Quiz struct {
	Id        bson.ObjectID  `bson:"_id,omitempty" json:"id"`
	Name      string         `bson:"name" json:"name"`
	Questions []QuizQuestion `bson:"questions" json:"questions"`
}

type QuizQuestion struct {
	Id          string       `bson:"id" json:"id"`
	Name        string       `bson:"name" json:"name"`
	QuizChoices []QuizChoice `bson:"quiz_choices" json:"quizChoices"`
}

type QuizChoice struct {
	Id      string `bson:"id" json:"id"`
	Name    string `bson:"name" json:"name"`
	Correct bool   `bson:"correct" json:"correct"`
}
