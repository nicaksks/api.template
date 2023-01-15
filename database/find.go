package database

import (
	"api/schemas"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func Find(name string) (*schemas.Data, error) {
	client, ctx := StartConnect()
	defer client.Disconnect(ctx)

	collection := client.Database("Anime").Collection("Catalog")
	filter := bson.M{"name": name}

	var request schemas.Data
	err := collection.FindOne(context.Background(), filter).Decode(&request)
	if err != nil {
		panic(err)
	}

	return &request, nil
}
