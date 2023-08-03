package handlers_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/spintoaguero/meli-challenge/internal/controllers/handlers"
	"github.com/spintoaguero/meli-challenge/internal/models"
	"github.com/spintoaguero/meli-challenge/internal/repositories"
	"github.com/spintoaguero/meli-challenge/pkg/mongodb"
)

func TestNewLinkHandler(t *testing.T) {
	result := handlers.NewLinkHandler(nil)
	resultType := reflect.TypeOf(result).String()
	if resultType != "*handlers.LinkHandler" {
		t.Errorf("got %v type, expected *handlers.LinkHandler type", resultType)
	}
}

func initRepositoryWithTestDB() *repositories.LinkRepository {
	dbClient, _ := mongodb.NewDbConnection(os.Getenv("MELI_MONGODB_URI"), os.Getenv("MELI_DB_TEST_NAME"))

	return &repositories.LinkRepository{
		Database:   dbClient,
		Collection: "links_test",
	}
}

var h *handlers.LinkHandler

func TestMain(m *testing.M) {
	r := initRepositoryWithTestDB()

	h = &handlers.LinkHandler{
		LinkRepository: r,
	}
	defer mongodb.CloseDbConnection(r.Database)

	code := m.Run()
	// Drop test DB after tests are executed
	r.Database.Client.Database(r.Database.Name).Drop(context.TODO())
	os.Exit(code)
}

func TestRedirect_ExistingUrl_Status303(t *testing.T) {
	// Create previous link
	link := &models.Link{
		ShortUrl: "ZZZZZZ",
		Url:      "https://www.mercadolibre.cl",
	}
	h.LinkRepository.Create(context.TODO(), link)

	r := mux.NewRouter()
	r.HandleFunc("/{key}", h.Redirect).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/ZZZZZZ")

	if err != nil {
		t.Errorf("got %s, expected nil", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("got %d, expected %d", res.StatusCode, http.StatusSeeOther)
	}
	if strings.Compare(res.Request.URL.String(), link.Url) != 0 {
		t.Errorf("got %s, expected %s", res.Request.URL, link.Url)
	}

	defer res.Body.Close()
}

func TestRedirect_NonExistingUrl_Status404(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/{key}", h.Redirect).Methods(http.MethodGet)
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/ZZZAAA")

	if err != nil {
		t.Errorf("got %s, expected nil", err.Error())
	}
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("got %d, expected %d", res.StatusCode, http.StatusNotFound)
	}

	defer res.Body.Close()
}

func TestCreateShortUrl_NewUrl_Status201(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/create", h.CreateShortUrl).Methods(http.MethodPost)
	ts := httptest.NewServer(r)
	defer ts.Close()

	body := []byte(`{
		"url": "https://mercado.libre/item-00001"
	}`)

	res, err := http.Post(ts.URL+"/api/v1/create", "application/json", bytes.NewBuffer(body))

	if err != nil {
		t.Errorf("got %s, expected nil", err.Error())
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("got %d, expected %d", res.StatusCode, http.StatusCreated)
	}

	defer res.Body.Close()
}

func TestCreateShortUrl_OldUrl_Status200(t *testing.T) {
	// Create previous link
	link := &models.Link{
		ShortUrl: "123456",
		Url:      "https://www.mercado.libre/item-00002",
	}
	h.LinkRepository.Create(context.TODO(), link)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/create", h.CreateShortUrl).Methods(http.MethodPost)
	ts := httptest.NewServer(r)
	defer ts.Close()

	body := []byte(`{
		"url": "https://www.mercado.libre/item-00002"
	}`)

	res, err := http.Post(ts.URL+"/api/v1/create", "application/json", bytes.NewBuffer(body))

	if err != nil {
		t.Errorf("got %s, expected nil", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("got %d, expected %d", res.StatusCode, http.StatusOK)
	}

	defer res.Body.Close()
}

func TestFindUrl_ExistingCompleteShortUrl_Status200(t *testing.T) {
	// Create previous link
	link := &models.Link{
		ShortUrl: "234567",
		Url:      "https://www.mercado.libre/item-00003",
	}
	h.LinkRepository.Create(context.TODO(), link)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/find", h.FindUrl).Methods(http.MethodPost)
	ts := httptest.NewServer(r)
	defer ts.Close()

	body := []byte(`{
		"short_url": "http://me.li/234567"
	}`)

	res, err := http.Post(ts.URL+"/api/v1/find", "application/json", bytes.NewBuffer(body))

	if err != nil {
		t.Errorf("got %s, expected nil", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("got %d, expected %d", res.StatusCode, http.StatusOK)
	}

	defer res.Body.Close()
}

func TestFindUrl_ExistingShortUrlKey_Status200(t *testing.T) {
	// Create previous link
	link := &models.Link{
		ShortUrl: "234567",
		Url:      "https://www.mercado.libre/item-00004",
	}
	h.LinkRepository.Create(context.TODO(), link)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/find", h.FindUrl).Methods(http.MethodPost)
	ts := httptest.NewServer(r)
	defer ts.Close()

	body := []byte(`{
		"short_url": "234567"
	}`)

	res, err := http.Post(ts.URL+"/api/v1/find", "application/json", bytes.NewBuffer(body))

	if err != nil {
		t.Errorf("got %s, expected nil", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("got %d, expected %d", res.StatusCode, http.StatusOK)
	}

	defer res.Body.Close()
}

func TestDeleteShortUrl_ExistingCompleteShortUrl_Status201(t *testing.T) {
	// Create previous link
	link := &models.Link{
		ShortUrl: "345678",
		Url:      "https://www.mercado.libre/item-00005",
	}
	h.LinkRepository.Create(context.TODO(), link)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/delete", h.DeleteShortUrl).Methods(http.MethodDelete)
	ts := httptest.NewServer(r)
	defer ts.Close()

	body := []byte(`{
		"short_url": "http://me.li/345678"
	}`)

	client := &http.Client{}

	req, _ := http.NewRequest("DELETE", ts.URL+"/api/v1/delete", bytes.NewBuffer(body))
	res, err := client.Do(req)

	if err != nil {
		t.Errorf("got %s, expected nil", err.Error())
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("got %d, expected %d", res.StatusCode, http.StatusCreated)
	}

	defer res.Body.Close()
}

func TestDeleteShortUrl_ExistingShortUrlKey_Status201(t *testing.T) {
	// Create previous link
	link := &models.Link{
		ShortUrl: "4567890",
		Url:      "https://www.mercado.libre/item-00006",
	}
	h.LinkRepository.Create(context.TODO(), link)

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/delete", h.DeleteShortUrl).Methods(http.MethodDelete)
	ts := httptest.NewServer(r)
	defer ts.Close()

	body := []byte(`{
		"short_url": "4567890"
	}`)

	client := &http.Client{}

	req, _ := http.NewRequest("DELETE", ts.URL+"/api/v1/delete", bytes.NewBuffer(body))
	res, err := client.Do(req)

	if err != nil {
		t.Errorf("got %s, expected nil", err.Error())
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("got %d, expected %d", res.StatusCode, http.StatusCreated)
	}

	defer res.Body.Close()
}
