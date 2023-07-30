package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOne(client *mongo.Client, ctx context.Context, database, col string, document interface{}) (*mongo.InsertOneResult, error) {

	// select database and collection ith Client.Database method and Database.Collection method
	collection := client.Database(database).Collection(col)

	// InsertOne accept two argument of type Context and of empty interface
	result, err := collection.InsertOne(ctx, document)
	return result, err
}
