package mongodb_test

import (
	"reflect"
	"testing"

	"github.com/spintoaguero/meli-challenge/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateFilter(t *testing.T) {
	expected := bson.M{"name": "john"}
	result := mongodb.CreateFilter("name", "john")

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("got %v, expected %v", result, expected)
	}
}
