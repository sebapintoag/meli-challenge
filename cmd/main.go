package main

import (
	"github.com/spintoaguero/meli-challenge/configs"
)

func main() {
	// get Client, Context, CancelFunc and err from connect method.
	client, ctx, cancel, err := configs.NewDbConnection("mongodb://admin:admin@mongodb:27017/?maxPoolSize=20&w=majority")
	if err != nil {
		panic(err)
	}

	// Release resource when main function is returned.
	defer configs.CloseDbConnection(client, ctx, cancel)

	configs.SetupRoutes()
}
