package admin

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spintoaguero/meli-challenge/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (mh *AdminHandler) GenerateShortUrl(w http.ResponseWriter, req *http.Request) {
	link := models.Link{
		FullUrl: "https://mercadolibre.cl/dsadsad",
	}
	err := link.Create(mh.Database, context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Link created: %v\n", link)

	link.Delete(mh.Database, context.Background())

	fmt.Printf("Link deleted: %v\n", link)

	objID, _ := primitive.ObjectIDFromHex("64c6c5b2de5ccbd917af34fc")

	newLink := models.Link{
		ID: objID,
	}

	newLink.Read(mh.Database, context.Background())

	fmt.Printf("Link readed: %v\n", newLink)

	fmt.Fprintf(w, link.ShortUrl)
}
