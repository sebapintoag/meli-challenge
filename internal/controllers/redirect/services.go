package redirect

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (rh *RedirectHandler) Perform(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Short URL: %v\n", vars["key"])
}
