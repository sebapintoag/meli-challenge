package models

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ILinkRepository interface {
	Find(ctx context.Context, filter interface{}) (*Link, error)
	Create(ctx context.Context, l *Link) (*Link, error)
	Delete(ctx context.Context, l *Link) error
}

type Link struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ShortUrl  string             `bson:"short_url" json:"short_url,omitempty"`
	Url       string             `bson:"url,omitempty" json:"url,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

// Generates a random string that will be used as a short URL key
func (l *Link) GenerateShortUrlKey() {
	keySize := 6
	alphabet := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	alphabetSize := len(alphabet)
	var builder strings.Builder

	for i := 0; i < keySize; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		builder.WriteRune(ch)
	}

	l.ShortUrl = builder.String()
}

// Replaces "http://redirect.link" + "/" in link.ShortURL (if present)
func (l *Link) RemoveUrlPrefix() {
	l.ShortUrl = strings.Replace(l.ShortUrl, fmt.Sprintf("%s/", os.Getenv("MELI_REDIRECT_URL")), "", 1)
}
