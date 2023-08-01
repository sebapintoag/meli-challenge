package admin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spintoaguero/meli-challenge/internal/models"
	"github.com/spintoaguero/meli-challenge/pkg/mongodb"
	"github.com/spintoaguero/meli-challenge/pkg/utils"
)

// Create short URL
func (ah *AdminHandler) CreateShortUrl(w http.ResponseWriter, req *http.Request) {
	// Parse body request
	var link models.Link
	if err := json.NewDecoder(req.Body).Decode(&link); err != nil {
		utils.ErrorResponse(w, req, "fail", http.StatusUnprocessableEntity, err)
		return
	}

	// Remove redirect URL from link.ShortURL
	link.RemoveUrlPrefix()

	// Return existing link if already present in DB
	if err := link.Find(ah.Database, context.Background(), mongodb.CreateFilter("url", link.Url)); err == nil {
		utils.SuccessResponse(w, req, ah.newLinkResponse(link), http.StatusOK)
		return
	}

	// Create new link for URL
	err := link.Create(ah.Database, context.Background())
	if err != nil {
		utils.ErrorResponse(w, req, "error", http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(w, req, ah.newLinkResponse(link), http.StatusCreated)
}

// Find short URL
func (ah *AdminHandler) FindUrl(w http.ResponseWriter, req *http.Request) {
	// Parse body request
	var link models.Link
	if err := json.NewDecoder(req.Body).Decode(&link); err != nil {
		utils.ErrorResponse(w, req, "fail", http.StatusUnprocessableEntity, err)
		return
	}

	// Remove redirect URL from link.ShortURL
	link.RemoveUrlPrefix()

	err := link.Find(ah.Database, context.Background(), mongodb.CreateFilter("short_url", link.ShortUrl))
	if err != nil {
		utils.ErrorResponse(w, req, "error", http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(w, req, ah.newLinkResponse(link), http.StatusOK)
}

// Delete short URL
func (ah *AdminHandler) DeleteShortUrl(w http.ResponseWriter, req *http.Request) {
	// Parse body request
	var link models.Link
	if err := json.NewDecoder(req.Body).Decode(&link); err != nil {
		utils.ErrorResponse(w, req, "fail", http.StatusUnprocessableEntity, err)
		return
	}

	// Remove redirect URL from link.ShortURL
	link.RemoveUrlPrefix()

	err := link.Find(ah.Database, context.Background(), mongodb.CreateFilter("short_url", link.ShortUrl))
	if err != nil {
		utils.ErrorResponse(w, req, "error", http.StatusInternalServerError, err)
		return
	}

	if err := link.Delete(ah.Database, context.Background()); err != nil {
		utils.ErrorResponse(w, req, "error", http.StatusInternalServerError, err)
		return
	}

	data := map[string]interface{}{
		"message": "link deleted",
	}

	utils.SuccessResponse(w, req, data, http.StatusCreated)
}

// Creates a new LinkResponse object
func (ah *AdminHandler) newLinkResponse(link models.Link) map[string]interface{} {
	return map[string]interface{}{
		"link": LinkResponse{
			Url:       link.Url,
			ShortUrl:  fmt.Sprintf("%s/%s", os.Getenv("MELI_REDIRECT_URL"), link.ShortUrl),
			CreatedAt: link.CreatedAt,
		},
	}
}
