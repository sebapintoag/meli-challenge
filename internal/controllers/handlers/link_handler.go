package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/spintoaguero/meli-challenge/internal/models"
	"github.com/spintoaguero/meli-challenge/pkg/mongodb"
	"github.com/spintoaguero/meli-challenge/pkg/utils"
)

type LinkHandler struct {
	LinkRepository models.ILinkRepository
}

type LinkResponse struct {
	ShortUrl  string    `json:"short_url"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}

// Creates a new link handler object
func NewLinkHandler(linkRepository models.ILinkRepository) *LinkHandler {
	return &LinkHandler{
		LinkRepository: linkRepository,
	}
}

// Creates a new LinkResponse object
func (l *LinkHandler) newLinkResponse(link *models.Link) map[string]interface{} {
	return map[string]interface{}{
		"link": LinkResponse{
			Url:       link.Url,
			ShortUrl:  fmt.Sprintf("%s/%s", os.Getenv("MELI_REDIRECT_URL"), link.ShortUrl),
			CreatedAt: link.CreatedAt,
		},
	}
}

// Performs a redirect from short URL to full URL
func (h *LinkHandler) Redirect(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	shortUrl := vars["key"]

	link, err := h.LinkRepository.Find(context.Background(), mongodb.CreateFilter("short_url", shortUrl))
	if err != nil {
		utils.ErrorResponse(w, req, "fail", http.StatusNotFound, err)
		return
	}

	http.Redirect(w, req, link.Url, http.StatusSeeOther)
}

// Create short URL
func (h *LinkHandler) CreateShortUrl(w http.ResponseWriter, req *http.Request) {
	// Parse body request
	var linkReq models.Link
	if err := json.NewDecoder(req.Body).Decode(&linkReq); err != nil {
		utils.ErrorResponse(w, req, "fail", http.StatusUnprocessableEntity, err)
		return
	}

	// Return existing link if already present in DB
	link, _ := h.LinkRepository.Find(context.Background(), mongodb.CreateFilter("url", linkReq.Url))
	if link != nil {
		utils.SuccessResponse(w, req, h.newLinkResponse(link), http.StatusOK)
		return
	}

	// Create new link for URL
	linkReq.GenerateShortUrlKey()
	link, err := h.LinkRepository.Create(context.Background(), &linkReq)
	if err != nil {
		utils.ErrorResponse(w, req, "error", http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(w, req, h.newLinkResponse(link), http.StatusCreated)
}

// Find short URL
func (h *LinkHandler) FindUrl(w http.ResponseWriter, req *http.Request) {
	// Parse body request
	var linkReq models.Link
	if err := json.NewDecoder(req.Body).Decode(&linkReq); err != nil {
		utils.ErrorResponse(w, req, "fail", http.StatusUnprocessableEntity, err)
		return
	}

	// Remove redirect URL from link.ShortURL
	linkReq.RemoveUrlPrefix()

	link, err := h.LinkRepository.Find(context.Background(), mongodb.CreateFilter("short_url", linkReq.ShortUrl))
	if err != nil {
		utils.ErrorResponse(w, req, "error", http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(w, req, h.newLinkResponse(link), http.StatusOK)
}

// Delete short URL
func (h *LinkHandler) DeleteShortUrl(w http.ResponseWriter, req *http.Request) {
	// Parse body request
	var linkReq models.Link
	if err := json.NewDecoder(req.Body).Decode(&linkReq); err != nil {
		utils.ErrorResponse(w, req, "fail", http.StatusUnprocessableEntity, err)
		return
	}

	// Remove redirect URL from link.ShortURL
	linkReq.RemoveUrlPrefix()

	link, err := h.LinkRepository.Find(context.Background(), mongodb.CreateFilter("short_url", linkReq.ShortUrl))
	if err != nil {
		utils.ErrorResponse(w, req, "error", http.StatusInternalServerError, err)
		return
	}

	if err := h.LinkRepository.Delete(context.Background(), link); err != nil {
		utils.ErrorResponse(w, req, "error", http.StatusInternalServerError, err)
		return
	}

	data := map[string]interface{}{
		"message": "link deleted",
	}

	utils.SuccessResponse(w, req, data, http.StatusCreated)
}
