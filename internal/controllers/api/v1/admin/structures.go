package admin

import "time"

type LinkResponse struct {
	ShortUrl  string    `json:"short_url"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
