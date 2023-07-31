package models

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/spintoaguero/meli-challenge/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Link struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ShortUrl  string             `bson:"short_url"`
	Url       string             `bson:"url,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}

func FindByShortUrl(db *mongodb.Database, ctx context.Context, shortUrl string, result interface{}) error {
	err := mongodb.FindOne(db.Client, ctx, "meli-db", "links", bson.M{"short_url": shortUrl}, &result)
	if err != nil {
		return err
	}
	return nil
}

func FindByUrl(db *mongodb.Database, ctx context.Context, url string, result interface{}) error {
	err := mongodb.FindOne(db.Client, ctx, "meli-db", "links", bson.M{"url": url}, &result)
	if err != nil {
		return err
	}
	return nil
}

func (l *Link) Find(db *mongodb.Database, ctx context.Context, filter interface{}) error {
	//bson.M{"_id": l.ID}
	err := mongodb.FindOne(db.Client, ctx, "meli-db", "links", filter, &l)
	if err != nil {
		return err
	}
	return nil
}

func (l *Link) Create(db *mongodb.Database, ctx context.Context) error {
	l.ShortUrl = generateShortUrl()
	l.CreatedAt = time.Now()
	l.UpdatedAt = time.Now()

	result, err := mongodb.InsertOne(db.Client, ctx, "meli-db", "links", &l)
	if err != nil {
		return err
	}

	l.ID = result.InsertedID.(primitive.ObjectID)

	return nil
}

func (l *Link) Delete(db *mongodb.Database, ctx context.Context) error {
	_, err := mongodb.DeleteOne(db.Client, ctx, "meli-db", "links", bson.M{"_id": l.ID})
	if err != nil {
		return err
	}

	return nil
}

func generateShortUrl() string {
	b := make([]byte, 3) //equals 6 characters
	rand.Read(b)

	return hex.EncodeToString(b)
}
