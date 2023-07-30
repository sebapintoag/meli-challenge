package configs

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client  *mongo.Client
	Context context.Context
	Cancel  context.CancelFunc
}

//func NewDbConnection(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
func NewDbConnection(uri string) (*Database, error) {

	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return nil, err
	}

	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	if err != nil {
		return nil, err
	}

	return &Database{
		Client:  client,
		Context: ctx,
		Cancel:  cancel,
	}, nil
}

func CloseDbConnection(database *Database) {

	defer database.Cancel()

	defer func() {
		if err := database.Client.Disconnect(database.Context); err != nil {
			panic(err)
		}
	}()
}
