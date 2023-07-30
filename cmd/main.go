package main

import (
	"github.com/spintoaguero/meli-challenge/internal/controllers"
	"github.com/spintoaguero/meli-challenge/pkg/mongodb"
)

func main() {
	// get Client, Context, CancelFunc and err from connect method.
	dbClient, err := mongodb.NewDbConnection("mongodb://admin:admin@mongodb:27017/?maxPoolSize=20&w=majority")
	if err != nil {
		panic(err)
	}

	// Release resource when main function is returned.
	defer mongodb.CloseDbConnection(dbClient)

	controllers.SetupRoutes(dbClient)
}
