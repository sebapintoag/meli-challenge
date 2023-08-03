package repositories

import (
	"context"
	"time"

	"github.com/spintoaguero/meli-challenge/internal/models"
	"github.com/spintoaguero/meli-challenge/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LinkRepository struct {
	Database   *mongodb.Database
	Collection string
}

func NewLinkRepository(database *mongodb.Database) *LinkRepository {
	return &LinkRepository{
		Database:   database,
		Collection: "links",
	}
}

func (r *LinkRepository) Find(ctx context.Context, filter interface{}) (*models.Link, error) {
	var link models.Link

	err := mongodb.FindOne(r.Database.Client, ctx, r.Database.Name, r.Collection, filter, &link)
	if err != nil {
		return nil, err
	}
	return &link, nil
}

func (r *LinkRepository) Create(ctx context.Context, l *models.Link) (*models.Link, error) {
	l.CreatedAt = time.Now()
	l.UpdatedAt = time.Now()

	result, err := mongodb.InsertOne(r.Database.Client, ctx, r.Database.Name, r.Collection, l)
	if err != nil {
		return nil, err
	}

	l.ID = result.InsertedID.(primitive.ObjectID)

	return l, nil
}

func (r *LinkRepository) Delete(ctx context.Context, l *models.Link) error {
	_, err := mongodb.DeleteOne(r.Database.Client, ctx, r.Database.Name, r.Collection, bson.M{"_id": l.ID})
	if err != nil {
		return err
	}

	return nil
}
