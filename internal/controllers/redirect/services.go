package redirect

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spintoaguero/meli-challenge/internal/models"
	"github.com/spintoaguero/meli-challenge/pkg/mongodb"
	"github.com/spintoaguero/meli-challenge/pkg/utils"
)

// Performs a redirect from short URL to full URL
func (rh *RedirectHandler) Perform(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	shortUrl := vars["key"]
	var link models.Link

	err := link.Find(rh.Database, context.Background(), mongodb.CreateFilter("short_url", shortUrl))
	if err != nil {
		utils.ErrorResponse(w, req, "fail", http.StatusUnprocessableEntity, err)
	}

	http.Redirect(w, req, link.Url, http.StatusSeeOther)
}
