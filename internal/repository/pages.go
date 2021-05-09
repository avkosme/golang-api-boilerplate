package pages

import (
	"context"

	"github.com/avkosme/golang-api-boilerplate/pkg/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

const collectionName string = "pages"

// All pages in the database
func FindAll() (results []bson.M) {

	dataFind := bson.D{}
	client := mongodb.Client()
	collection := client.Collection(collectionName)

	cursor, err := collection.Find(context.Background(), dataFind)

	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return
}
