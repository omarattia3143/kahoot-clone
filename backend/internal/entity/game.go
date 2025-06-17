package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Game struct {
	Id              primitive.ObjectID
	Quiz            Quiz
	CurrentQuestion int
	Code            string
}
