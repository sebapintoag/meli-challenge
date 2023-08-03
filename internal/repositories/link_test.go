package repositories_test

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/spintoaguero/meli-challenge/internal/models"
	"github.com/spintoaguero/meli-challenge/internal/repositories"
	"github.com/spintoaguero/meli-challenge/pkg/mongodb"
)

func initRepositoryWithTestDB() *repositories.LinkRepository {
	dbClient, _ := mongodb.NewDbConnection(os.Getenv("MELI_MONGODB_URI"), os.Getenv("MELI_DB_TEST_NAME"))

	return &repositories.LinkRepository{
		Database:   dbClient,
		Collection: "links_test",
	}
}

var r *repositories.LinkRepository

func TestMain(m *testing.M) {
	r = initRepositoryWithTestDB()
	defer mongodb.CloseDbConnection(r.Database)

	code := m.Run()
	// Drop test DB after tests are executed
	r.Database.Client.Database(r.Database.Name).Drop(context.TODO())
	os.Exit(code)
}

func TestNewLinkRepository(t *testing.T) {
	result := repositories.NewLinkRepository(nil)
	resultType := reflect.TypeOf(result).String()
	if resultType != "*repositories.LinkRepository" {
		t.Errorf("got %v type, expected *repositories.LinkRepository type", resultType)
	}
}

func TestFind_ByUrl_Present(t *testing.T) {
	// Insert one record
	expectedLink := &models.Link{
		ShortUrl: "XXYYZZ",
		Url:      "https://www.mercado.libre/item-12345",
	}
	collection := r.Database.Client.Database(r.Database.Name).Collection(r.Collection)
	collection.InsertOne(context.TODO(), expectedLink)

	// Find record
	link, _ := r.Find(context.TODO(), mongodb.CreateFilter("url", expectedLink.Url))

	if !reflect.DeepEqual(expectedLink.ShortUrl, link.ShortUrl) {
		t.Errorf("got expectedLink.ShortUrl != link.ShortUrl, expected expectedLink.ShortUrl == link.ShortUrl")
	}
}

func TestFind_ByUrl_NotPresent(t *testing.T) {
	// Find record
	link, _ := r.Find(context.TODO(), mongodb.CreateFilter("url", "https://www.mercado.libre/item-54321"))

	if link != nil {
		t.Errorf("got link, expected nil")
	}
}

func TestFind_ByShortUrl_Present(t *testing.T) {
	// Insert one record
	expectedLink := &models.Link{
		ShortUrl: "AABBCC",
		Url:      "https://www.mercado.libre/item-67890",
	}
	collection := r.Database.Client.Database(r.Database.Name).Collection(r.Collection)
	collection.InsertOne(context.TODO(), expectedLink)

	// Find record
	link, _ := r.Find(context.TODO(), mongodb.CreateFilter("short_url", expectedLink.ShortUrl))

	if !reflect.DeepEqual(expectedLink.Url, link.Url) {
		t.Errorf("got expectedLink.Url != link.Url, expected expectedLink.ShortUrl == link.ShortUrl")
	}
}

func TestFind_ByShortUrl_NotPresent(t *testing.T) {
	// Find record
	link, _ := r.Find(context.TODO(), mongodb.CreateFilter("url", "https://www.mercado.libre/item-54321"))

	if link != nil {
		t.Errorf("got link, expected nil")
	}
}

func TestCreate_NewUrl_Succeeded(t *testing.T) {
	newLink := &models.Link{
		ShortUrl: "123456",
		Url:      "https://www.mercado.libre/item-00001",
	}

	_, err := r.Create(context.TODO(), newLink)
	if err != nil {
		t.Errorf("got error, expected nil")
	}
}

func TestCreate_OldUrl_Failed(t *testing.T) {
	link := &models.Link{
		ShortUrl: "567890",
		Url:      "https://www.mercado.libre/item-00003",
	}

	r.Create(context.TODO(), link)

	newLink := &models.Link{
		ShortUrl: "567890",
		Url:      "https://www.mercado.libre/item-00004",
	}

	_, err := r.Create(context.TODO(), newLink)
	if err != nil {
		t.Errorf("got error, expected nil")
	}
}

func TestCreate_Delete_Succeeded(t *testing.T) {
	// Insert one record
	link := &models.Link{
		ShortUrl: "DDDDDD",
		Url:      "https://www.mercado.libre/item-00003",
	}
	collection := r.Database.Client.Database(r.Database.Name).Collection(r.Collection)
	collection.InsertOne(context.TODO(), link)

	err := r.Delete(context.TODO(), link)
	if err != nil {
		t.Errorf("got error, expected nil")
	}

	var newLink *models.Link
	collection.FindOne(context.TODO(), mongodb.CreateFilter("short_url", link.ShortUrl)).Decode(newLink)
	if newLink != nil {
		t.Errorf("got link, expected nil")
	}
}
