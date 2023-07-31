package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spintoaguero/meli-challenge/internal/models"
	"github.com/spintoaguero/meli-challenge/pkg/mongodb"
	"github.com/spintoaguero/meli-challenge/pkg/utils"
)

func (mh *AdminHandler) CreateShortUrl(w http.ResponseWriter, req *http.Request) {
	// Parse body request
	var link models.Link
	if err := json.NewDecoder(req.Body).Decode(&link); err != nil {
		utils.ErrorResponse(w, req, "fail", err)
		return
	}

	err := link.Create(mh.Database, context.Background())
	if err != nil {
		utils.ErrorResponse(w, req, "error", err)
		return
	}

	data := map[string]interface{}{
		"link": LinkResponse{
			Url:       link.Url,
			ShortUrl:  link.ShortUrl,
			CreatedAt: link.CreatedAt,
		},
	}

	utils.SuccessResponse(w, req, data)
}

func (mh *AdminHandler) FindUrl(w http.ResponseWriter, req *http.Request) {
	// Parse body request
	var linkReq LinkRequest
	fmt.Println(req.Body)
	if err := json.NewDecoder(req.Body).Decode(&linkReq); err != nil {
		utils.ErrorResponse(w, req, "fail", err)
		return
	}

	var link models.Link

	err := link.Find(mh.Database, context.Background(), mongodb.CreateFilter("short_url", linkReq.ShortUrl))
	if err != nil {
		utils.ErrorResponse(w, req, "error", err)
		return
	}

	data := map[string]interface{}{
		"link": LinkResponse{
			Url:       link.Url,
			ShortUrl:  link.ShortUrl,
			CreatedAt: link.CreatedAt,
		},
	}

	utils.SuccessResponse(w, req, data)
}

func (mh *AdminHandler) DeleteShortUrl(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	shortUrl := vars["key"]

	var link models.Link
	err := link.Find(mh.Database, context.Background(), mongodb.CreateFilter("short_url", shortUrl))
	if err != nil {
		utils.ErrorResponse(w, req, "error", err)
		return
	}

	if err := link.Delete(mh.Database, context.Background()); err != nil {
		utils.ErrorResponse(w, req, "error", err)
		return
	}

	data := map[string]interface{}{
		"message": "link deleted",
	}

	utils.SuccessResponse(w, req, data)
}
