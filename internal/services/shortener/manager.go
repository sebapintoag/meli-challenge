package shortener

import (
	"crypto/rand"
	"encoding/hex"
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

func GenerateShortUrl(w http.ResponseWriter, req *http.Request) {
	b := make([]byte, 3) //equals 6 characters
	rand.Read(b)
	key := hex.EncodeToString(b)
	fmt.Fprintf(w, key)
}
