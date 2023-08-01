package main

import (
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

	// Release resource when main function is returned.
	defer mongodb.CloseDbConnection(dbClient)

	controllers.SetupRoutes(dbClient)
}
