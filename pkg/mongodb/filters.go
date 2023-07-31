package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateFilter(name, value string) primitive.M {
	return bson.M{name: value}
}
