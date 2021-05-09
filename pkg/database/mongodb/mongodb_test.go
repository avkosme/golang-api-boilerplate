package mongodb

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	client := Client()
	client.Drop(context.Background())
}

func TestClientInit(t *testing.T) {
	r := Client()
	actualFormat := fmt.Sprintf("%T", r)
	if "*mongo.Database" != actualFormat {
		t.Error(actualFormat)
	}
}

// go test -v ./pkg/database/mongodb/... -run TestClientInsertOne
func TestClientInsertOne(t *testing.T) {

	var collectionName string = "pages"
	data := bson.D{
		{"name", "testing name"},
		{"title", "testing title"},
	}
	client := Client()
	collection := client.Collection(collectionName)

	res, err := collection.InsertOne(context.Background(), data)

	if err != nil {
		panic(err)
	}
	actualFormat := fmt.Sprintf("%T", res)

	if "*mongo.InsertOneResult" != actualFormat {
		t.Error(actualFormat)
	}
}

func TestFindAll(t *testing.T) {

	var collectionName string = "pages"
	data := bson.D{
		{"name", "testing name second"},
		{"title", "testing title second"},
	}
	dataFind := bson.D{}
	client := Client()
	collection := client.Collection(collectionName)

	_, err := collection.InsertOne(context.Background(), data)

	if err != nil {
		panic(err)
	}
	cursor, err := collection.Find(context.Background(), dataFind)

	if err != nil {
		panic(err)
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	if "testing name" != results[0]["name"] {
		t.Error(results[0])
	}

	if "testing name second" != results[1]["name"] {
		t.Error(results[1])
	}

}
