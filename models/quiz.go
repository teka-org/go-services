package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quiz struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Question string             `bson:"question" json:"question"`
	Answer   string             `bson:"answer" json:"answer"`
	Option1  string             `bson:"option1" json:"option1"`
	Option2  string             `bson:"option2" json:"option2"`
	Option3  string             `bson:"option3" json:"option3"`
}
