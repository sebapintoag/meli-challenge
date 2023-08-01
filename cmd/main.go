package main

import (
	"context"
	"fmt"
	"os"

	"github.com/spintoaguero/meli-challenge/internal/controllers"
	"github.com/spintoaguero/meli-challenge/pkg/mongodb"
)

func main() {
	// get Client, Context, CancelFunc and err from connect method.
	dbClient, err := mongodb.NewDbConnection(os.Getenv("MELI_MONGODB_URI"))
	if err != nil {
		panic(err)
	}

	// Create short_url index in db.links
	err = mongodb.CreateUniqueIndex(dbClient.Client, context.Background(), "meli-db", "links", "short_url")
	if err != nil {
		fmt.Println(err)
	}

	// Release resource when main function is returned.
	defer mongodb.CloseDbConnection(dbClient)

	controllers.SetupRoutes(dbClient)
}
