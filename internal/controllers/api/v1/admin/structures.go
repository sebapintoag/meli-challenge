package admin

import "time"

type LinkRequest struct {
	ShortUrl string `json:"short_url,omitempty"`
	URL      string `json:"url,omitempty"`
}

type LinkResponse struct {
	ShortUrl  string    `json:"short_url"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
