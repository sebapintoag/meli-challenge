package redirect

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spintoaguero/meli-challenge/internal/models"
	"github.com/spintoaguero/meli-challenge/pkg/mongodb"
	"github.com/spintoaguero/meli-challenge/pkg/utils"
)

func (rh *RedirectHandler) Perform(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	shortUrl := vars["key"]
	var link models.Link

	err := link.Find(rh.Database, context.Background(), mongodb.CreateFilter("short_url", fmt.Sprintf("%s/%s", os.Getenv("MELI_REDIRECT_URL"), shortUrl)))
	if err != nil {
		utils.ErrorResponse(w, req, "fail", http.StatusUnprocessableEntity, err)
	}

	http.Redirect(w, req, link.Url, http.StatusSeeOther)
}
