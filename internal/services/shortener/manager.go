package shortener

import (
	"fmt"
	"net/http"
)

func Headers(w http.ResponseWriter, req *http.Request) {

	for _, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", "aaaa", h)
		}
	}
}
