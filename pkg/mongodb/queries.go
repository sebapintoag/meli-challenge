package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func FindOne(client *mongo.Client, ctx context.Context, databaseName, collectionName string, filter interface{}, result interface{}) error {
	collection := client.Database(databaseName).Collection(collectionName)

	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func InsertOne(client *mongo.Client, ctx context.Context, databaseName, collectionName string, document interface{}) (*mongo.InsertOneResult, error) {
	// select database and collection ith Client.Database method and Database.Collection method
	collection := client.Database(databaseName).Collection(collectionName)

	// InsertOne accept two argument of type Context and of empty interface
	result, err := collection.InsertOne(ctx, document)
	return result, err
}

func DeleteOne(client *mongo.Client, ctx context.Context, databaseName, collectionName string, query interface{}) (*mongo.DeleteResult, error) {

	// select document and collection
	collection := client.Database(databaseName).Collection(collectionName)

	// query is used to match a document  from the collection.
	result, err := collection.DeleteOne(ctx, query)
	return result, err
}
