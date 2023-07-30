package services

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/spintoaguero/meli-challenge/configs"
	"github.com/spintoaguero/meli-challenge/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

type ManagerHandler struct {
	Database *configs.Database
}

func Headers(w http.ResponseWriter, req *http.Request) {

	for _, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", "aaaa", h)
		}
	}
}

func (mh *ManagerHandler) GenerateShortUrl(w http.ResponseWriter, req *http.Request) {
	b := make([]byte, 3) //equals 6 characters
	rand.Read(b)
	key := hex.EncodeToString(b)

	var document interface{}

	document = bson.D{
		{"rollNo", 175},
		{"maths", 80},
		{"science", 90},
		{"computer", 95},
	}

	// insertOne accepts client , context, database
	// name collection name and an interface that
	// will be inserted into the  collection.
	// insertOne returns an error and a result of
	// insert in a single document into the collection.
	insertOneResult, err := mongodb.InsertOne(mh.Database.Client, mh.Database.Context, "gfg",
		"marks", document)

	// handle the error
	if err != nil {
		panic(err)
	}

	// print the insertion id of the document,
	// if it is inserted.
	fmt.Println("Result of InsertOne")
	fmt.Println(insertOneResult.InsertedID)

	fmt.Fprintf(w, key)
}
